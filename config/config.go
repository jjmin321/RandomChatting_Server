package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

// Config - 설정 파일
var Config config

type configMethod interface {
	InitConfig()
}

type config struct {
	App app
	DB  database `toml:"database"`
}

type app struct {
	Name string `toml:"name"`
}

type database struct {
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

// InitConfig Config 데이터 초기화
func InitConfig() {
	configBytes, err := ioutil.ReadFile("config/config.toml")
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(configBytes), &Config)
	if err != nil {
		panic(err)
	}

	log.Print("[CONFIG] 환경설정 초기화")
}
