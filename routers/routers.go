package routers

import (
	"GinHello/handlers/hello"
	"GinHello/handlers/userInfo"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	HelloRouter := routers.Group("/hello")
	{
		HelloRouter.GET("/hello", hello.Hello)
	}
	PostgresRouter := routers.Group("/user")
	{
		PostgresRouter.POST("/insert", userInfo.InsertUser)
		PostgresRouter.GET("/deleteById", userInfo.DeleteUserById)
		PostgresRouter.POST("/updateById", userInfo.UpdateUserInfo)
		PostgresRouter.GET("/findAll", userInfo.FindAllUser)

	}
	return routers
}
