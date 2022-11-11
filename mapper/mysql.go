package mapper

import (
	"GinHello/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var MySqlSession *gorm.DB

func ConnectMysqlDB() (err error) {
	DriverName := global.MysqlDbSetting.DriverName
	Host := global.MysqlDbSetting.Host
	Port := global.MysqlDbSetting.Port
	User := global.MysqlDbSetting.User
	Password := global.MysqlDbSetting.Password
	Dbname := global.MysqlDbSetting.Dbname
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", User, Password, Host, Port, Dbname)
	MySqlSession, err = gorm.Open(DriverName, psqlInfo)
	if err != nil {
		log.Println("connectDBError", err)
	}
	return MySqlSession.DB().Ping()
}

func CloseMysql() {
	MySqlSession.Close()
}
