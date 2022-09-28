package global

import (
	"GinHello/pkg/setting"
	"log"
)

// 和yml对应
type ServerSettingS struct {
	RunMode  string
	HttpPort string
}

type PostgresDbSettings struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
	Sslmode    string
}

// 定义全局变量
var (
	ServerSetting     *ServerSettingS
	PostgresDbSetting *PostgresDbSettings
)

// 读取配置到全局便量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	pgdb, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = pgdb.ReadSection("Postgres", &PostgresDbSetting)
	if err != nil {
		return err
	}
	log.Printf("pgSetting:")
	log.Printf("%+v", PostgresDbSetting)
	return nil
}
