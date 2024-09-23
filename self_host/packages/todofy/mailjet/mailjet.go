package mailjet

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/cloudmailin"
	"github.com/ziyixi/monorepo/self_host/packages/todofy/utils"
)

const (
	sender       = "ziyixi@mailjet.ziyixi.science"
	senderName   = "Todofy"
	receiverName = "dida365"
)

//go:embed todoDescription.tmpl
var descriptionTmpl string

// SendEmail sends an email using mailjet
func SendEmail(c *gin.Context, mailContent cloudmailin.MailInfo) (interface{}, error) {
	envs, ok := c.Get("envs")
	if !ok {
		return nil, fmt.Errorf("envs is not in context")
	}
	envsMap, ok := envs.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("envs is not a map")
	}

	for _, env := range []string{"MJ_APIKEY_PUBLIC", "MJ_APIKEY_PRIVATE", "DIDA_EMAIL"} {
		if _, ok := envsMap[env]; !ok {
			return nil, fmt.Errorf("%s is not in envs", env)
		}
	}
	publicKey := envsMap["MJ_APIKEY_PUBLIC"]
	privateKey := envsMap["MJ_APIKEY_PRIVATE"]
	didaEmail := envsMap["DIDA_EMAIL"]

	mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)

	// prepare task description, load template
	tmpl, err := template.New("todoistDescription").Parse(descriptionTmpl)
	if err != nil {
		return nil, fmt.Errorf("parse template todoDescription.tmpl error: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, mailContent)
	if err != nil {
		return nil, fmt.Errorf("execute template todoDescription.tmpl error: %w", err)
	}
	taskDescription := buf.String()

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: sender,
				Name:  senderName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: didaEmail,
					Name:  receiverName,
				},
			},
			Subject:  fmt.Sprintf("%v [%v]", mailContent.Subject, mailContent.From),
			TextPart: taskDescription,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return nil, fmt.Errorf("mailjet send email error: %w", err)
	}

	if len(res.ResultsV31) == 0 || len(res.ResultsV31[0].To) == 0 {
		return nil, fmt.Errorf("mailjet send email API response error: %v", res)
	}
	mailjetHref := res.ResultsV31[0].To[0].MessageHref

	// send request to mailjet API to get email send status
	result, err := utils.FetchWithBasicAuth(mailjetHref, publicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("fetch mailjet email status error: %w", err)
	}

	return result, nil
}
