package main

import (
	"GinHello/global"
	"GinHello/mapper"
	"GinHello/message/mqtt/config"
	"GinHello/message/udp"
	"GinHello/routers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	config.InitMqtt()

}

func main() {

	var err error
	if strings.Compare(global.DataSourceSetting.Source, "mysql") == 0 {
		err = mapper.ConnectMysqlDB()
	} else {
		err = mapper.ConnectPGDB()
	}
	if err != nil {
		log.Printf("%+v", err)
	}
	if strings.Compare(global.DataSourceSetting.Source, "mysql") == 0 {
		defer mapper.CloseMysql()
	} else {
		defer mapper.ClosePG()
	}

	udp.UdpServer()
	udp.UdpClient()

	f, _ := os.Create("./logs/logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
