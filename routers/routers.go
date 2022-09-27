package routers

import (
	"GinHello/handlers/hello"
	"GinHello/handlers/postgres"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	HelloRouter := routers.Group("/hello")
	{
		HelloRouter.GET("/hello", hello.Hello)
	}
	PostgreRouter := routers.Group("/db")
	{
		PostgreRouter.GET("/connect", postgres.DbOpen)
	}
	return routers
}
