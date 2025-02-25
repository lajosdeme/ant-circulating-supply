package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ArbRpcUrl string `mapstructure:"ARB_RPC_URL"`
}

var C Config

func Load() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read env file: ", err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatal("failed to decode env file: ", err)
	}
}
