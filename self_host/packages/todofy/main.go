package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/utils"

	"github.com/gofiber/fiber/v2"
)

var (
	allowedUsers = flag.String("allowed-users", "", "Comma-separated list of allowed users in the format 'username:password'")
	port         = flag.Int("port", 8080, "Port to run the server on")
)

var GitCommit string // Will be set by Bazel at build time

func main() {
	log.Infof("Server Starting time: %s", time.Now().Format(time.RFC3339))
	flag.Parse()
	if *allowedUsers == "" {
		log.Fatal("No allowed users provided. Use --allowed-users flag to specify them.")
	}

	app := fiber.New(
		fiber.Config{
			AppName: fmt.Sprintf("todofy (git commit: %s)", GitCommit[:7]),
		},
	)

	allowedUsers, allowedUsersStrings := utils.ParseAllowedUsers(*allowedUsers)
	if len(allowedUsers) == 0 {
		log.Fatal("No valid users found in the allowed users list.")
	}
	log.Infof("Allowed users (hidden passwords): %s", allowedUsersStrings)
	app.Use(basicauth.New(basicauth.Config{
		Users: allowedUsers,
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/update_todo", func(c *fiber.Ctx) error {
		return handleUpdateTodo(c)
	})
	v1.Get("/summary", func(c *fiber.Ctx) error {
		return handleSummary(c)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
