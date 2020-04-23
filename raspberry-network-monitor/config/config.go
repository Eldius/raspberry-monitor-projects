package config

import (
	"fmt"

	"github.com/spf13/viper"
	yml "gopkg.in/yaml.v2"
)

/*
MQTTConfig is a config abstraction for the MQTT client
*/
type MQTTConfig struct {
	Host       string
	Port       string
	User       string
	Pass       string
	ClientName string
	Topic      string
	Qos        byte
}

/*
PingHosts is a list of hosts
*/
type PingHosts []string

/*
NetworkMonitorConfig is an abstraction for the app config
*/
type NetworkMonitorConfig struct {
	PingHosts  PingHosts
	QtdPackets int
	MQTT       MQTTConfig
}

/*
AppConfig loads the config from file
*/
func AppConfig() NetworkMonitorConfig {
	var config NetworkMonitorConfig
	if err := viper.UnmarshalKey("cfg", &config); err != nil {
		panic(err.Error())
	}

	return config
}

/*
WriteConfig writes config to file
*/
func WriteConfig(cfg NetworkMonitorConfig) string {
	if cfgBytes, err := yml.Marshal(cfg); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(string(cfgBytes))
		return string(cfgBytes)
	}
}
