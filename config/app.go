package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type AppConfig struct {
	Debug    bool
	Server   Server
	Context  Context
	Database Database
}

type Server struct {
	Address string
}

type Context struct {
	Timeout int
}

type Database struct {
	Host string
	User string
	Pass string
	Name string
}

var cfg *AppConfig

func InitAppConfig() {
	var (
		appConfig = new(AppConfig)
		err       error
	)

	jsonFile, err := os.Open("./config.json")
	if err != nil {
		panic(fmt.Sprintf("error while open config file, err=%v", err))
	}
	jsonByte, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(jsonByte, appConfig)
	if err != nil {
		panic(fmt.Sprintf("error while unmarshal config file, err=%v", err))
	}
	cfg = appConfig
}

func GetConfig() *AppConfig {
	return cfg
}
