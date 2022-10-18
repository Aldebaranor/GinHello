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
		UserRouter.POST("/insert", controller.UserControl.InsertUser)
		UserRouter.GET("/findAll", controller.UserControl.GetUserList)
		UserRouter.GET("/getUser/:id", controller.UserControl.GetUserById)
		UserRouter.PUT("/update", controller.UserControl.UpdateUser)
		UserRouter.DELETE("/delete/:id", controller.UserControl.DeleteUserById)

	}
	CommentRouter := routers.Group("/comment")
	{
		CommentRouter.GET("/goRoutineBatchInsert", controller.CommentContr.GoBatchInsert)
		CommentRouter.GET("/batchInsert", controller.CommentContr.BatchInsert)
	}
	MathRouter := routers.Group("/q1")
	{
		MathRouter.GET("/q1")
	}
	MqttRouter := routers.Group("/mqtt")
	{
		MqttRouter.GET("/sub", controller.MqttControl.Subscribe)
		MqttRouter.POST("/pub", controller.MqttControl.Produce)
	}

	return routers
}
