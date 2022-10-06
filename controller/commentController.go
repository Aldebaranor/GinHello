package controller

import (
	"GinHello/entity"
	"GinHello/service"
	"log"
	"time"
)

var commentContr = &commentController{}

type commentController struct {
}

func (t *commentController) BatchInsert() {
	comments := []*entity.Comment{}
	for i := 0; i < 9999; i++ {
		comment := entity.Comment{}
		comment.Content = string(time.Now().Unix())
		comments = append(comments, &comment)
	}
	err := service.CommentSer.BatchInsert(comments)
	if err != nil {
		log.Println("controllerError", err)
	}
}
