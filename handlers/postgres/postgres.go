package postgres

import (
	"GinHello/global"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func DbOpen(context *gin.Context) {
	db, _ = sql.Open(global.PostgresDbSetting.Name,
		global.PostgresDbSetting.Host+" "+
			global.PostgresDbSetting.Port+" "+
			global.PostgresDbSetting.User+" "+
			global.PostgresDbSetting.Password+" "+
			global.PostgresDbSetting.Dbname+" "+
			global.PostgresDbSetting.Sslmode)
	err := db.Ping()
	if err != nil {
		log.Printf("%+v", err)
	}
	context.JSON(200, gin.H{
		"msg": err,
	})
}
