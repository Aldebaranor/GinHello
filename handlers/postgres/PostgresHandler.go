package postgres

import (
	"GinHello/global"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// 连接数据库
func ConnectDB() *sql.DB {
	DriverName := global.PostgresDbSetting.DriverName
	Host := global.PostgresDbSetting.Host
	Port := global.PostgresDbSetting.Port
	User := global.PostgresDbSetting.User
	Password := global.PostgresDbSetting.Password
	Dbname := global.PostgresDbSetting.Dbname
	Sslmode := global.PostgresDbSetting.Sslmode
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", Host, Port, User, Password, Dbname, Sslmode)
	db, err := sql.Open(DriverName, psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
