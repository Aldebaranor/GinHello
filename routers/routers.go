package routers

import (
	"GinHello/controller"
	"GinHello/controller/hello"
	"GinHello/controller/userInfo"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	HelloRouter := routers.Group("/hello")
	{
		HelloRouter.GET("/hello", hello.Hello)
	}
	noORMRouter := routers.Group("/noORM")
	{
		noORMRouter.POST("/insert", userInfo.InsertUser)
		noORMRouter.GET("/deleteById", userInfo.DeleteUserById)
		noORMRouter.POST("/updateById", userInfo.UpdateUserInfo)
		noORMRouter.GET("/findAll", userInfo.FindAllUser)

	}
	ORMRouter := routers.Group("/user")
	{
		ORMRouter.POST("/insert", controller.InsertUser)
		ORMRouter.GET("/findAll", controller.GetUserList)
		ORMRouter.GET("/getUser/:id", controller.GetUserById)
		ORMRouter.PUT("/update", controller.UpdateUser)
		ORMRouter.DELETE("/delete/:id", controller.DeleteUserById)

	}
	return routers
}
