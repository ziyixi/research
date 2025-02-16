package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"slices"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/ziyixi/monorepo/self_host/packages/todofy/proto"
)

var (
	port         = flag.Int("port", 50051, "The server port of the LLM service")
	geminiApiKey = flag.String("gemini-api-key", "", "The API key for Gemini")
)

type LLMServer struct {
	pb.UnimplementedLLMSummaryServiceServer
}

func (s *LLMServer) Summarize(ctx context.Context, req *pb.LLMSummaryRequest) (*pb.LLMSummaryResponse, error) {
	if !slices.Contains(supportedModelFamily, req.ModelFamily) {
		return nil, status.Errorf(codes.InvalidArgument, "Unsupported model family: %s", req.ModelFamily)
	}
	maxTokens := tokenLimit
	if req.MaxTokens != 0 {
		maxTokens = req.MaxTokens
	}
	prompt := DefaultpromptToSummaryEmail
	if req.Prompt != "" {
		prompt = req.Prompt
	}
	selectedModels := llmModelPriority
	if req.Model != pb.Model_MODEL_UNSPECIFIED {
		selectedModels = []pb.Model{req.Model}
	}

	for _, model := range selectedModels {
		if _, ok := llmModelNames[model]; !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Unsupported model: %s", model)
		}
		var (
			summary string
			err     error
		)
		switch req.ModelFamily {
		case pb.ModelFamily_MODEL_FAMILY_GEMINI:
			summary, err = SummaryEmailByGemini(prompt, req.Text, model, maxTokens)
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Unsupported model family: %s", req.ModelFamily)
		}
		if err != nil || len(summary) == 0 {
			log.Printf("Error generating summary with model %s: %v", model, err)
			continue
		}
		return &pb.LLMSummaryResponse{
			Summary: summary,
		}, nil
	}
	return nil, status.Errorf(codes.Internal, "Failed to generate summary with all models: %v", selectedModels)
}

func SummaryEmailByGemini(prompt, content string, llmModel pb.Model, maxTokens int32) (string, error) {
	// set up client
	ctx := context.Background()
	if *geminiApiKey == "" {
		return "", status.Error(codes.InvalidArgument, "gemini-api-key is empty")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(*geminiApiKey))
	if err != nil {
		return "", fmt.Errorf("genai.NewClient error: %v", err)
	}
	defer client.Close()

	contentWithPrompt := fmt.Sprintf("%s\n%s", prompt, content)

	// create messages and limit to be smaller than 8192 tokens
	llmModelName, ok := llmModelNames[llmModel]
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Unsupported model: %s", llmModel)
	}
	model := client.GenerativeModel(llmModelName)
	respToken, err := model.CountTokens(ctx, genai.Text(contentWithPrompt))
	if err != nil {
		return "", fmt.Errorf("EncodingForModel: %v", err)
	}

	for respToken.TotalTokens > maxTokens {
		contentWithPrompt = contentWithPrompt[:len(contentWithPrompt)/10*9]
		respToken, err = model.CountTokens(ctx, genai.Text(contentWithPrompt))
		if err != nil {
			return "", fmt.Errorf("EncodingForModel: %v", err)
		}
	}

	// generate summary
	resp, err := model.GenerateContent(ctx, genai.Text(contentWithPrompt))
	if err != nil {
		return "", fmt.Errorf("client.CreateChatCompletion error: %v", err)
	}
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "", fmt.Errorf("client.CreateChatCompletion: no candidates returned")
	}

	resPart := resp.Candidates[0].Content.Parts[0]
	res := fmt.Sprintf("%v", resPart)

	return res, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterLLMSummaryServiceServer(srv, &LLMServer{})
	reflection.Register(srv)

	log.Printf("LLM server is running on port %d", *port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
