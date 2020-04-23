package client

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
Subscribe subscribes to broker
*/
//func Subscribe(cfg config.ClientConfig) {
func Subscribe(user, pass, host, port string) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	brokerServer := fmt.Sprintf("tcp://%s:%s", host, port)
	fmt.Println("user: ", user)
	fmt.Println("pass: ", pass)
	fmt.Println("server: ", brokerServer)

	connOpts := MQTT.NewClientOptions()
	connOpts.AddBroker(brokerServer)
	connOpts.SetClientID("cfg.ClientID")
	connOpts.SetCleanSession(true)

	if user != "" {
		connOpts.SetUsername(user)
		if pass != "" {
			connOpts.SetPassword(pass)
		}
	}

	connOpts.OnConnect = func(c MQTT.Client) {
		handler := func(c MQTT.Client, msg MQTT.Message) {
			fmt.Println(fmt.Sprintf("---\n%s:\n%s\n---", msg.Topic(), string(msg.Payload())))
			msg.Ack()
		}
		if token := c.Subscribe("#", 2, handler); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", host)
	}
	<-ch
}
