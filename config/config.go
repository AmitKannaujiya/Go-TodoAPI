package config

import (
	_ "fmt"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var once sync.Once

type App struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	App    App    `yaml:"app"`
	Tables Tables `yaml:"tables"`
	DB     Db     `yaml:"db"`
}

type Tables struct {
	Todos string `yaml:"todos"`
	Users string `yaml:"users"`
}

type Db struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUserName string `yaml:"db_username"`
	DbPassword string `yaml:"db_password"`
}

func GetConfig() (*Config, error) {
	var config Config
	var e error
	once.Do(func() {
		configFile, err := os.ReadFile("./config/config.yml")
		if err != nil {
			log.Fatalf("Config file not found %v : ", err)
			e = err
			return
		}
		er := yaml.Unmarshal(configFile, &config)

		if er != nil {
			log.Fatalf("Config file not able to unmarshall %v", er)
			e = er
			return
		}
	})

	if e != nil {
		return nil, e
	}
	return &config, nil
}
