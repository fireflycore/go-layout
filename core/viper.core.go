package core

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"microservice-go/config"
	"microservice-go/utils"
	"os"
)

func (_Entrance) SetupViper(file *[]byte) *config.Entrance {
	logPrefix := "setup viper"

	utils.Use.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	v := viper.New()
	v.SetConfigType("yaml")

	var _config config.Entrance

	if _, e := os.Stat("config.yaml"); e != nil {
		if err := v.ReadConfig(bytes.NewReader(*file)); err != nil {
			panic("config file not fount")
		}
	} else {
		v.SetConfigFile("config.yaml")
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Viper ReadInConfig error: %s \n", err))
		}
	}

	if err := v.Unmarshal(&_config); err != nil {
		panic(fmt.Errorf("Viper unmarshal error: %s \n", err))
	}

	utils.Use.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	return &_config
}
