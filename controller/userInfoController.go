package controller

import (
	"GinHello/entity"
	"GinHello/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	UserControl = &UserController{}
)

type UserController struct {
}

func (t *UserController) InsertUser(ctx *gin.Context) {
	user := entity.User_info{}
	ctx.ShouldBind(&user)
	user.CreateTime = time.Now()
	err := service.UserServ.InsertUser(&user)
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
func (t *UserController) GetUserList(ctx *gin.Context) {
	userList, err := service.UserServ.FindAllUser()
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
func (t *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := service.UserServ.GetUserById(id)
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
func (t *UserController) UpdateUser(ctx *gin.Context) {
	user := entity.User_info{}
	ctx.ShouldBind(&user)
	err := service.UserServ.UpdateUser(&user)
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
func (t *UserController) DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := service.UserServ.DeleteUserById(id)
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
