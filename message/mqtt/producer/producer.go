package producer

import (
	"GinHello/global"
	"GinHello/message/mqtt/config"
)

func Produce(msg string) {
	config.MqttClient.Publish(global.MqttSetting.Topic, global.MqttSetting.QoS, true, msg)
}
