package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func setup(t *testing.T) {
	currPath, err := os.Getwd()
	if err != nil {
		t.Errorf("Error trying to find current dir")
	}
	fmt.Println("current path", currPath)
	configFile := filepath.Join(currPath, "samples/config.yml")
	fmt.Println("config file", configFile)

	if _, err := os.Stat(configFile); err != nil {
		t.Errorf("Config file doesn't exists:\n'%s'", configFile)
	}

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		t.Errorf("Failed using config file: %v\n%v", viper.ConfigFileUsed(), err.Error())
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestAppConfigLoad(t *testing.T) {
	setup(t)
	cfg := AppConfig()

	if len(cfg.PingHosts) != 5 {
		t.Errorf("Must have 5 hosts to ping (%v)", cfg.PingHosts)
	}

	if cfg.QtdPackets != 4 {
		t.Errorf("Must send 4 packets to ping (%v)", cfg.QtdPackets)
	}

	if cfg.MQTT.Host != "localhost" {
		t.Errorf("Host must be localhost (%v)", cfg.MQTT.Host)
	}
	if cfg.MQTT.User != "user" {
		t.Errorf("Host must be user (%v)", cfg.MQTT.User)
	}
	if cfg.MQTT.Pass != "pass" {
		t.Errorf("Host must be pass (%v)", cfg.MQTT.Pass)
	}
}

func TestAppConfigSample(t *testing.T) {
	setup(t)
	pingHosts := PingHosts{
		"8.8.8.8",
		"1.1.1.1",
		"8.8.4.4",
		"192.168.0.1",
		"192.168.100.1",
	}

	mqttConfig := MQTTConfig{
		Host: "localhost",
		User: "user",
		Pass: "pass",
	}

	cfg := NetworkMonitorConfig{
		PingHosts: pingHosts,
		MQTT:      mqttConfig,
	}

	WriteConfig(cfg)

}
