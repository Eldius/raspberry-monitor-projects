package config

import "github.com/spf13/viper"

/*
ClientConfig is the configuration object
*/
type ClientConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	ClientID string
}

/*
AppConfig loads the config from file
*/
func AppConfig() ClientConfig {
	var config ClientConfig
	if err := viper.Unmarshal(&config); err != nil {
		panic(err.Error())
	}

	return config
}
