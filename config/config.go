package config

import (
	"fmt"
	"os"

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
	pflag.String("db_user", defaultConfig.ConfDBUSER, "db user")
	pflag.String("db_pass", defaultConfig.ConfDBPASS, "db pass")

	pflag.Parse()
	var err error
	Gee, err = setConfig(map[string]interface{}{})
	if err != nil {
		fmt.Printf("Config setting Error: %v\n", err)
		os.Exit(1)
	}
	Gee.BindPFlags(pflag.CommandLine)
}

func setConfig(defaults map[string]interface{}) (*ViperConfig, error) {
	// setting config process: default -> env file -> command line
	v := viper.New()
	for k, d := range defaults {
		v.SetDefault(k, d)
	}

	v.AddConfigPath("./")
	v.AddConfigPath("./config")
	v.SetConfigName(".env")
	if err := v.ReadInConfig(); err != nil {
		return &ViperConfig{}, err
	}

	return &ViperConfig{
		Viper: v,
	}, nil
}
