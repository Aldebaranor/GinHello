package controller

import (
	"GinHello/message/mqtt/producer"
	"GinHello/message/mqtt/subscriber"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	MqttControl = &MqttController{}
)

type MqttController struct {
}

func (p *MqttController) Subscribe(ctx *gin.Context) {
	subscriber.Subscribe()
	return
}

func (p *MqttController) Produce(ctx *gin.Context) {
	cJson := make(map[string]interface{})
	ctx.ShouldBind(&cJson)
	msg := fmt.Sprintf("%v", cJson["msg"])
	producer.Produce(msg)
	return
}
