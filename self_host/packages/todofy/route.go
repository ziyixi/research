package main

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/utils"
)

func handleUpdateTodo(c *fiber.Ctx) error {
	jsonRaw, err := io.ReadAll(bytes.NewReader(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	jsonString := string(jsonRaw)
	emailContent := utils.ParseCloudmailin(jsonString)
	if len(emailContent.From) == 0 || len(emailContent.To) == 0 || (len(emailContent.Subject) == 0 && len(emailContent.Content) == 0) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid email content, from/to/subject/content is empty")
	}
	return c.SendStatus(fiber.StatusOK)
}

func handleSummary(c *fiber.Ctx) error {
	return c.SendString("Summary")
}
