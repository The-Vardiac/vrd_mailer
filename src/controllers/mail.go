package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/williamluisan/vrd_mailer/utils"
)

type sendData struct {
	Subject	string	`json:"subject"`
	Body	string	`json:"body"`
	MailTo	string	`json:"mailto"`
}

func Send(c *gin.Context) {
	var data sendData

	// pls add validation
	// ...

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// send email
	var sendEmailUtil utils.Mail
	sendEmailUtil.Subject = data.Subject
	sendEmailUtil.Body = data.Body
	sendEmailUtil.To = data.MailTo
	sendEmailUtil.From = os.Getenv("SMTP_EMAIL_FROM")

	if err := sendEmailUtil.SendMail(); err != nil {
		// write to log
		log.Println("Failed to send email to: " + sendEmailUtil.To)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "done",
	})
}