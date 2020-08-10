package config

import "github.com/spf13/viper"

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// Gee ...
var Gee *ViperConfig

// Configs ...
var Configs = map[string]interface{}{
	"serverPort":     10620,
	"serverTimeout":  30,
	"serverLogLevel": "debug",

	"dbHost": "",
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
