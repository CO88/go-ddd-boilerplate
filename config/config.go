package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type DBInfo struct {
	Host         string `mapstructure:"hostname"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"dbname"`
	UserName     string `mapstructure:"username"`
	UserPassword string `mapstructure:"password"`
}

type Context struct {
	Timeout int
}

type Server struct {
	Address string
}

type Configuration struct {
	Db      DBInfo `mapstructure:"DB"`
	Context Context
	Server  Server
}

// singleton
var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {

	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Config read error")
			panic(err)
		}

		config = &Configuration{}

		err = viper.Unmarshal(&config)
		if err != nil {
			log.Fatal("config unmarshal error")
			panic(err)
		}
	})
	return config
}
