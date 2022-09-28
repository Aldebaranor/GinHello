package controller

import (
	"GinHello/entity"
	"GinHello/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InsertUser(ctx *gin.Context) {
	user := entity.User_info{}
	ctx.ShouldBind(&user)
	user.CreateTime = time.Now()
	err := service.InsertUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "新增用户成功",
			"data": user,
		})
	}
}
func GetUserList(ctx *gin.Context) {
	userList, err := service.FindAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": userList,
		})
	}
}
func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := service.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": user,
		})
	}
}
func UpdateUser(ctx *gin.Context) {
	user := entity.User_info{}
	ctx.ShouldBind(&user)
	err := service.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改用户成功",
			"data": user,
		})
	}
}
func DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := service.DeleteUserById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除用户成功",
		})
	}
}
