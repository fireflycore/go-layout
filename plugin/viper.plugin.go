package plugin

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"microservice-go/config"
	"os"
)

func SetupViper(file *[]byte) *config.EntranceEntity {
	logPrefix := "setup viper"

	fmt.Printf("%s %s\n", logPrefix, "start ->")

	v := viper.New()
	v.SetConfigType("yaml")

	var _config config.EntranceEntity

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

	fmt.Printf("%s %s\n", logPrefix, "success ->")

	return &_config
}
