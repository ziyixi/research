package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/utils"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

var (
	allowedUsers       = flag.String("allowed-users", "", "Comma-separated list of allowed users in the format 'username:password'")
	port               = flag.Int("port", 8080, "Port to run the server on")
	healthCheckTimeout = flag.Int("health-check-timeout", 10, "Timeout for health check in seconds")
	// GRPC addresses for the services
	llmAddr      = flag.String("llm-addr", ":50051", "Address of the LLM server")
	todoAddr     = flag.String("todo-addr", ":50052", "Address of the Todo server")
	databaseAddr = flag.String("database-addr", ":50053", "Address of the database server")
)

var GitCommit string // Will be set by Bazel at build time

func main() {
	log.Infof("Server Starting time: %s", time.Now().Format(time.RFC3339))
	flag.Parse()
	if *allowedUsers == "" {
		log.Fatal("No allowed users provided. Use --allowed-users flag to specify them.")
	}

	grpcClients, err := NewGRPCClients()
	if err != nil {
		log.Fatalf("Failed to create gRPC clients: %v", err)
	}
	defer grpcClients.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*healthCheckTimeout)*time.Second)
	defer cancel()
	if err := grpcClients.WaitForHealthy(ctx, time.Duration(*healthCheckTimeout)*time.Second); err != nil {
		log.Fatalf("Failed to connect to gRPC services: %v", err)
	}
	log.Infof("Connected to gRPC services: LLM: %s, Todo: %s, Database: %s", *llmAddr, *todoAddr, *databaseAddr)

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(utils.RateLimitMiddleware())

	allowedUsers, allowedUsersStrings := utils.ParseAllowedUsers(*allowedUsers)
	if len(allowedUsers) == 0 {
		log.Fatal("No valid users found in the allowed users list.")
	}
	log.Infof("Allowed users (hidden passwords): %s", allowedUsersStrings)

	api := app.Group("/api", gin.BasicAuth(gin.Accounts(allowedUsers)))
	v1 := api.Group("/v1")

	v1.POST("/update_todo", func(c *gin.Context) {
		handleUpdateTodo(c)
	})
	v1.GET("/summary", func(c *gin.Context) {
		handleSummary(c)
	})

	// start the server
	listenAddr := fmt.Sprintf(":%d", *port)
	log.Infof("Git commit: %s\n", GitCommit)
	log.Infof("Gin has started in %s mode on %s", gin.Mode(), listenAddr)
	if err := app.Run(listenAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
