package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleUpdateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Summary",
	})
}

func handleSummary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Summary",
	})
}
