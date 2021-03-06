package mqttclient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/config"
	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/logger"
	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/network"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
SendPingResponse sends the ping to MQTT broker
*/
func SendPingResponse(p network.PingResponse, cfg config.MQTTConfig) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.Host, cfg.Port))

	opts.SetClientID(cfg.ClientName)

	if cfg.User != "" {
		opts.SetUsername(cfg.User)
	}

	if cfg.Pass != "" {
		opts.SetPassword(cfg.Pass)
	}
	//opts.SetCleanSession(*cleansess)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Sample Publisher Started")
	log.Println("---- doing publish ----")
	topic := fmt.Sprintf("%s/%s", cfg.Topic, p.Host)
	token := client.Publish(topic, cfg.Qos, false, serialize(p.Convert()))
	token.Wait()

	client.Disconnect(250)
	log.Println("Sample Publisher Disconnected")
}

/*
SendPingResponses sends a list of ping results
*/
func SendPingResponses(pings []network.PingResponse, cfg config.MQTTConfig) {
	for _, p := range pings {
		SendPingResponse(p, cfg)
	}
}

func serialize(obj interface{}) []byte {
	if data, err := json.Marshal(obj); err != nil {
		panic(err.Error())
	} else {
		logger.Debug(fmt.Sprintf("serialized:\n%s", string(data)))
		return data
	}
}
