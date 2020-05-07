package listener

import (
	"net/http"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*
MqttMessageHandler handles the message and send it to websockets...
*/
func MqttMessageHandler(c MQTT.Client, msg MQTT.Message) {

}
