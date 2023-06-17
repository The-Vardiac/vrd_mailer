package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	config "github.com/williamluisan/vrd_mailer/config"
	"github.com/williamluisan/vrd_mailer/routes"
	"github.com/williamluisan/vrd_mailer/services"
)

func init() {
	var config config.Config

	// initialize godotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize rabbitmq
	config.InitRabbitmq()
}

func main() {
	var mailServices services.Mail

	defer config.RabbitmqChCons.Close()

	// consume from queue
	mailServices.RMQConsumeVrdMailerQueue()

	// initialize gin
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":4748"))
}