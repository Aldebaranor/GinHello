package userInfo

import (
	"GinHello/controller/postgres"
	"GinHello/entity"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// CRUD
func InsertUser(context *gin.Context) {
	db := postgres.ConnectDB()
	cJson := make(map[string]interface{})
	context.Bind(&cJson)
	userName := cJson["user_name"]
	log.Printf("%+v", userName)
	stmt, err := db.Prepare("INSERT INTO user_info(user_name,create_time) VALUES ($1,$2)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(userName, time.Now())
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	db.Close()
	log.Printf("%+v", res)
	context.JSON(200, gin.H{
		"msg":         "新增成功",
		"user_name":   userName,
		"create_time": time.Now(),
	})
}

func DeleteUserById(context *gin.Context) {
	db := postgres.ConnectDB()
	uId := context.Query("id")
	log.Printf("%+v", uId)
	stmt, err := db.Prepare("DELETE  FROM user_info WHERE u_id=$1")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(uId)
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	db.Close()
	log.Printf("%+v", res)
	context.JSON(200, gin.H{
		"msg": "删除成功",
	})
}

func UpdateUserInfo(context *gin.Context) {
	db := postgres.ConnectDB()
	cJson := make(map[string]interface{})
	context.Bind(&cJson)
	uId := cJson["u_id"]
	userName := cJson["user_name"]
	log.Printf("%+v", uId)
	log.Printf("%+v", userName)
	stmt, err := db.Prepare("UPDATE user_info SET user_name = $1 WHERE u_id = $2")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(userName, uId)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	db.Close()
	log.Printf("%+v", res)
	context.JSON(200, gin.H{
		"msg":       "更新成功",
		"u_id":      uId,
		"user_name": userName,
	})

}

func FindAllUser(context *gin.Context) {
	db := postgres.ConnectDB()
	stmt, err := db.Query("SELECT * FROM user_info")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var userInfoList []entity.User_info
	for stmt.Next() {
		user := entity.User_info{}
		err := stmt.Scan(&user.UId, &user.UserName, &user.CreateTime)
		if err != nil {
			panic(err)
		}
		userInfoList = append(userInfoList, user)
		log.Printf("%+v", user)
	}
	defer stmt.Close()
	db.Close()
	context.JSON(200, gin.H{
		"userInfoList": userInfoList,
	})
}
