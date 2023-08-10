package routers

import (
	"webserver/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("*.html")

	router.GET("/", controllers.LoginPage)
	router.GET("/profile", controllers.UserPage)
	router.GET("/user/:email", controllers.GetUserByEmail)

	return router
}
