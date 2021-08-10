package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var AppConfig *viper.Viper

func init()  {
	if AppConfig == nil {
		AppConfig = viper.New()
		AppConfig.SetConfigName("config")
		AppConfig.SetConfigType("yaml")
		AppConfig.AddConfigPath("./conf/")

		if err := AppConfig.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}
