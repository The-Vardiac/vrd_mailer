package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/williamluisan/vrd_mailer/controllers"
)

func Routes(router *gin.Engine) {
	router.POST("/send", controllers.Send)
}