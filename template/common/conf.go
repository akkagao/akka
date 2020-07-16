package common

import (
	"log"

	"github.com/akkagao/goutil"
	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

/**
加载配置文件
*/
func InitConfig(configFile string) {
	Config = viper.New()
	Config.SetConfigFile(configFile)
	err := Config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		goutil.ChkErr(err)
	}
	log.Println("config env value :", Config.Get("env"))
}
