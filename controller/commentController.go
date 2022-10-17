package controller

import (
	"GinHello/entity"
	"GinHello/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

var (
	CommentContr = &CommentController{}
)

type CommentController struct {
}

func (t *CommentController) GoBatchInsert(ctx *gin.Context) {
	comments := []*entity.Comment{}
	for i := 0; i < 10000; i++ {
		comment := entity.Comment{}
		comment.CId = strconv.Itoa(i)
		comment.Content = time.Now().String()
		comments = append(comments, &comment)
	}
	err := service.CommentSer.GoBatchInsert(comments)
	if err != nil {
		log.Println("3333 controllerError", err)
	}
}

func (t *CommentController) BatchInsert(ctx *gin.Context) {
	comments := []*entity.Comment{}
	for i := 0; i < 10000; i++ {
		comment := entity.Comment{}
		comment.CId = strconv.Itoa(i)
		comment.Content = time.Now().String()
		comments = append(comments, &comment)
	}
	err := service.CommentSer.BatchInsert(comments)
	if err != nil {
		log.Println("3333 controllerError", err)
	}
}
