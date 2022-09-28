package userInfo

import (
	"GinHello/handlers/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// CRUD
func InsertUser(context *gin.Context) {
	cJson := make(map[string]interface{})
	context.Bind(&cJson)
	user_name := cJson["user_name"]
	log.Printf("%+v", user_name)
	stmt, err := postgres.DB.Prepare("INSERT INTO user_info(user_name,create_time) VALUES ($1,$2)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(user_name, time.Now())
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", res)
}
