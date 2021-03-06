package listener

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
MqttClient is the configuration object
*/
type MqttClient struct {
	user string
	pass string
	host string
	port string
}

/*
Subscribe connects to the server
*/
func (c *MqttClient) Subscribe(mqttMsgHandler func(c MQTT.Client, msg MQTT.Message)) {
	brokerServer := fmt.Sprintf("tcp://%s:%s", c.host, c.port)

	connOpts := MQTT.NewClientOptions()
	connOpts.AddBroker(brokerServer)
	connOpts.SetClientID("cfg.ClientID")
	connOpts.SetCleanSession(true)

	if c.user != "" {
		connOpts.SetUsername(c.user)
		if c.pass != "" {
			connOpts.SetPassword(c.pass)
		}
	}

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe("#", 2, mqttMsgHandler); token.Wait() && token.Error() != nil {
			log.Panic(token.Error())
		}

	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Printf("Connected to %s\n", c.host)
	}
}

/*
NewMqttClient subscribes to broker
*/
func NewMqttClient(user, pass, host, port string) MqttClient {
	return MqttClient{
		user: user,
		pass: pass,
		host: host,
		port: port,
	}
}
