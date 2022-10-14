package service

import (
	"GinHello/entity"
	"GinHello/mapper"
	"errors"
	"log"
	"strconv"
	"time"
)

var CommentSer = &CommentService{}

type CommentService struct {
}

func (t *CommentService) BatchInsert(comments []*entity.Comment) (err error) {
	start := time.Now().Unix()
	datas := make([]*entity.Comment, 1000)
	index := 0
	ch := make(chan int, 5)
	i := 0
	nums := 0
	for commentNum, comment := range comments {
		datas[commentNum%1000] = comment
		index++
		i++
		if i == 1000 {
			go batch(datas, ch)
			log.Printf("开启第[" + strconv.Itoa(nums) + "]个goroutine")
			if len(comments) < (nums+1)*1000 {
				datas = make([]*entity.Comment, len(comments)-(nums+1)*1000)
			} else {
				datas = make([]*entity.Comment, 1000)
			}
			i = 0
			nums++
			continue
		}
		if len(comments) == index {
			go batch(datas, ch)
			log.Printf("开启第[" + strconv.Itoa(nums) + "]个goroutine")
			nums++
		}
	}
	count := 0
	for {
		err := <-ch
		count++
		if err == 1 {
			return errors.New("2222 batch insert error")
		}
		if count == nums {
			break
		}
	}
	end := time.Now().Unix()
	log.Printf("耗时：", end-start)
	return
}

func batch(datas []*entity.Comment, ch chan int) {
	err := mapper.SqlSession.Create(&datas)
	if err != nil {
		log.Printf("1111 batch insert sqlexec create err:", err)
		ch <- 1
		return
	}
	ch <- 0
}
