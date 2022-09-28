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
	PostgreRouter := routers.Group("/user")
	{
		PostgreRouter.POST("/insert", userInfo.InsertUser)
		PostgreRouter.GET("/deleteById", userInfo.DeleteUserById)
		PostgreRouter.POST("/updateById", userInfo.UpdateUserInfo)
		PostgreRouter.GET("/findAll", userInfo.FindAllUser)

	}
	return routers
}
