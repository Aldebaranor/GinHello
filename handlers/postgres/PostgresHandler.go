package postgres

import (
	"GinHello/global"
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// 连接数据库
func ConnectDB() *sql.DB {
	db, err := sql.Open(global.PostgresDbSetting.Name,
		global.PostgresDbSetting.Host+" "+
			global.PostgresDbSetting.Port+" "+
			global.PostgresDbSetting.User+" "+
			global.PostgresDbSetting.Password+" "+
			global.PostgresDbSetting.Dbname+" "+
			global.PostgresDbSetting.Sslmode)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
