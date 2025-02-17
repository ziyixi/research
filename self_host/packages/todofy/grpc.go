package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/ziyixi/monorepo/self_host/packages/todofy/proto"
)

type GRPCClients struct {
	llmConn        *grpc.ClientConn
	llmClient      pb.LLMSummaryServiceClient
	todoConn       *grpc.ClientConn
	todoClient     pb.TodoServiceClient
	databaseConn   *grpc.ClientConn
	databaseClient pb.DataBaseServiceClient
}

func NewGRPCClients() (*GRPCClients, error) {
	llmConn, err := grpc.NewClient(*llmAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LLM server: %w", err)
	}

	todoConn, err := grpc.NewClient(*todoAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Todo server: %w", err)
	}

	databaseConn, err := grpc.NewClient(*databaseAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Database server: %w", err)
	}

	return &GRPCClients{
		llmConn:        llmConn,
		llmClient:      pb.NewLLMSummaryServiceClient(llmConn),
		todoConn:       todoConn,
		todoClient:     pb.NewTodoServiceClient(todoConn),
		databaseConn:   databaseConn,
		databaseClient: pb.NewDataBaseServiceClient(databaseConn),
	}, nil
}

func (c *GRPCClients) Close() {
	if c.llmConn != nil {
		c.llmConn.Close()
	}
	if c.todoConn != nil {
		c.todoConn.Close()
	}
	if c.databaseConn != nil {
		c.databaseConn.Close()
	}
}

func (c *GRPCClients) WaitForHealthy(ctx context.Context, timeout time.Duration) error {
	// Validate connections first
	if c.llmConn == nil || c.todoConn == nil || c.databaseConn == nil {
		return fmt.Errorf("one or more gRPC connections are nil")
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Pre-create health clients to validate them
	services := []struct {
		name   string
		conn   *grpc.ClientConn
		client grpc_health_v1.HealthClient
	}{
		{
			name:   "llm",
			conn:   c.llmConn,
			client: grpc_health_v1.NewHealthClient(c.llmConn),
		},
		{
			name:   "todo",
			conn:   c.todoConn,
			client: grpc_health_v1.NewHealthClient(c.todoConn),
		},
		{
			name:   "database",
			conn:   c.databaseConn,
			client: grpc_health_v1.NewHealthClient(c.databaseConn),
		},
	}

	errChan := make(chan error, len(services))

	for _, service := range services {
		go func(name string, client grpc_health_v1.HealthClient) {
			for {
				select {
				case <-ctx.Done():
					errChan <- fmt.Errorf("health check timeout for %s", name)
					return
				default:
					req := &grpc_health_v1.HealthCheckRequest{}
					resp, err := client.Check(ctx, req)

					if err != nil {
						// Log the error for debugging
						log.Warningf("Health check error for %s: %v", name, err)
						time.Sleep(100 * time.Millisecond)
						continue
					}

					if resp.Status == grpc_health_v1.HealthCheckResponse_SERVING {
						errChan <- nil
						return
					}

					time.Sleep(500 * time.Millisecond)
				}
			}
		}(service.name, service.client)
	}

	var errors []error
	for i := 0; i < len(services); i++ {
		if err := <-errChan; err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("health check failed: %v", errors)
	}

	return nil
}
