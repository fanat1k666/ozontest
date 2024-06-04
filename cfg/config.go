package cfg

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	Pg       bool   `mapstructure:"pg"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile("cfg/.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &env
}
