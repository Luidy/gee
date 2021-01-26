package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// DefaultConfig ...
type DefaultConfig struct {
	ConfENV        string
	ConfServerPort int
	ConfDBHOST     string
	ConfDBPORT     int
	ConfDBNAME     string
	ConfDBUSER     string
	ConfDBPASS     string
}

var defaultConfig = DefaultConfig{
	ConfServerPort: 18000,
	ConfDBHOST:     "gee",
	ConfDBPORT:     3306,
	ConfDBNAME:     "geedb",
	ConfDBUSER:     "geeuser",
	ConfDBPASS:     "geepass",
}

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// Gee ...
var Gee *ViperConfig

func init() {
	pflag.IntP("port", "p", defaultConfig.ConfServerPort, "server port")
	pflag.String("db_host", defaultConfig.ConfDBHOST, "db host")
	pflag.Int("db_port", defaultConfig.ConfDBPORT, "db port")
	pflag.String("db_name", defaultConfig.ConfDBNAME, "db name")
	pflag.String("db_pass", defaultConfig.ConfDBPASS, "db pass")

	pflag.Parse()

}

func setConfig(defaults map[string]interface{}) *viper.Viper {
	// default ->
	config := viper.New()
	for k, v := range defaults {
		config.SetDefault(k, v)
	}

	return config
}

// Configs ...
var Configs = map[string]interface{}{
	"serverPort":     18000,
	"serverTimeout":  30,
	"serverLogLevel": "debug",

	"dbHost": "gee",
	"dbPort": 3306,
	"dbUser": "geeuser",
	"dbPass": "geepass",
	"dbName": "geedb",
}

func init() {
	v := viper.New()
	for key, value := range Configs {
		v.SetDefault(key, value)
	}
	Gee = &ViperConfig{
		Viper: v,
	}
}
