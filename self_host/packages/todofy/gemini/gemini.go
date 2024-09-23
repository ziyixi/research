package gemini

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/utils"
	"google.golang.org/api/option"
)

//go:embed prompt.tmpl
var promptTmpl string

// SummaryByGemini generates a summary of the content using Gemini
func SummaryByGemini(ctx *gin.Context, content string) (string, error) {
	// set up client
	envsMap, err := utils.ExtractEnv(ctx)
	geminiApiKey, ok := envsMap["GEMINI_API_KEY"]
	if !ok {
		return "", fmt.Errorf("GEMINI_API_KEY is not in envs")
	}

	c := context.Background()
	client, err := genai.NewClient(c, option.WithAPIKey(geminiApiKey))
	if err != nil {
		return "", fmt.Errorf("genai.NewClient error: %w", err)
	}
	defer client.Close()

	tmpl, err := template.New("prompt").Parse(promptTmpl)
	if err != nil {
		return "", fmt.Errorf("template.ParseFiles error: %w", err)
	}

	// Create a buffer to hold the rendered output
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		EmailContent string
	}{
		EmailContent: content,
	})
	if err != nil {
		return "", fmt.Errorf("tmpl.Execute error: %v", err)
	}
	contentWithPrompt := buf.String()

	// create messages and limit to be smaller than 8192 tokens
	model := client.GenerativeModel("gemini-1.5-pro")
	respToken, err := model.CountTokens(c, genai.Text(contentWithPrompt))
	if err != nil {
		return "", fmt.Errorf("token count error: %w", err)
	}

	for respToken.TotalTokens > 8192 {
		contentWithPrompt = contentWithPrompt[:len(contentWithPrompt)/10*9]
		respToken, err = model.CountTokens(c, genai.Text(contentWithPrompt))
		if err != nil {
			return "", fmt.Errorf("token count error: %w", err)
		}
	}

	// generate summary
	resp, err := model.GenerateContent(c, genai.Text(contentWithPrompt))
	if err != nil {
		return "", fmt.Errorf("generate content from gemini error: %w", err)
	}
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "", fmt.Errorf("no content found in response from gemini")
	}

	resPart := resp.Candidates[0].Content.Parts[0]
	res := fmt.Sprintf("%v", resPart)

	return res, nil
}
