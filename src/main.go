package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/williamluisan/vrd_mailer/routes"
)

func init() {
	// initialize godotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// initialize gin
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":4748"))
}