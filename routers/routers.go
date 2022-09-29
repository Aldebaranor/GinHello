package routers

import (
	"GinHello/controller"
	"GinHello/controller/hello"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	HelloRouter := routers.Group("/hello")
	{
		HelloRouter.GET("/hello", hello.Hello)
	}
	UserRouter := routers.Group("/user")
	{
		UserRouter.POST("/insert", controller.InsertUser)
		UserRouter.GET("/findAll", controller.GetUserList)
		UserRouter.GET("/getUser/:id", controller.GetUserById)
		UserRouter.PUT("/update", controller.UpdateUser)
		UserRouter.DELETE("/delete/:id", controller.DeleteUserById)

	}
	return routers
}
