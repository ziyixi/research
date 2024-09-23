package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/cloudmailin"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/gemini"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/mailjet"
)

func updateTodoPost(c *gin.Context) {
	// get the post data
	jsonRaw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in reading json body": err.Error()})
		return
	}
	jsonString := string(jsonRaw)
	emailContent := cloudmailin.ParseCloudmailin(jsonString)
	if len(emailContent.From) == 0 || len(emailContent.To) == 0 || (len(emailContent.Subject) == 0 && len(emailContent.Content) == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error in parsing json body": "from/to/subject/content is empty"})
		return
	}

	// summary the email content
	geminiContent, err := gemini.SummaryByGemini(c, emailContent.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in generating gemini summary": err.Error()})
		return
	}
	emailContent.Content = geminiContent

	// send the email
	res, err := mailjet.SendEmail(c, emailContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in sending mailjet email": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": res})
}
