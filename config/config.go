package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbSchema   string
}

func GetConfig() Configuration {
	config := Configuration{}

	err := gonfig.GetConf("config/config.json", &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
