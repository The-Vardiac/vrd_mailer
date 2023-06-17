package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/williamluisan/vrd_mailer/config"
	"github.com/williamluisan/vrd_mailer/repository"
	"github.com/williamluisan/vrd_mailer/utils"
)

type Mail repository.Mail

func (mail *Mail) RMQConsumeVrdMailerQueue() {
	var sendData repository.SendData

	log.Println("Start consuming vrdmailerqueue..")
	msgs, err := config.RabbitmqChCons.Consume(
		"vrdmailerqueue", // queue
		"",               // consumer
		false,            // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to consume a message", err)
	}

	go func() {
		for d := range msgs {

			err := json.Unmarshal([]byte(d.Body), &sendData)
			if err != nil {
				log.Printf("%s: %s", "Failed to unmarshal json", err)
			}

			if err = d.Ack(false); err != nil {
				log.Fatal("RabbitMQ: failed to acknowledge message in queue: " + string(d.Body))
			}

			// send email
			log.Println("Sending email to: " + sendData.MailTo + " ..")

			var sendEmailUtil utils.Mail
			sendEmailUtil.Subject = sendData.Subject
			sendEmailUtil.Body = sendData.Body
			sendEmailUtil.To = sendData.MailTo
			sendEmailUtil.From = os.Getenv("SMTP_EMAIL_FROM")

			if err := sendEmailUtil.SendMail(); err != nil {
				// write to log
				log.Println("Failed to send email to: " + sendEmailUtil.To)
			}
			log.Println("Email sent.")
		}
	}()
}