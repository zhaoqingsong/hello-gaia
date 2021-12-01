package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	viper.SetConfigName("config-dev")
	viper.AddConfigPath("./zhaoqingsong/studyweb/config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		//log.Log.Error(err.Error())
		log.Println(err.Error())
	}

	// 启动时解析出错，程序退出
	conf, err := parseConfig()
	if err != nil {
		os.Exit(1)
	}
	Config = *conf
}

// parseConfig 解析并解密配置
func parseConfig() (*config, error) {
	var conf config
	//  解析
	err := viper.Unmarshal(&conf)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
