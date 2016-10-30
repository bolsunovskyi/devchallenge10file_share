package config

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"time"
)

type Mongo struct {
	Host	string
	Port	uint
	DB	string
	Timeout	time.Duration
}

type AppConfig struct {
	Secret		string
	Port		uint
	Mongo		Mongo
	DataFolder	string
}

var File string = "config.toml"
var Config AppConfig

func Read(path string) bool {
	if _, err := toml.DecodeFile(fmt.Sprintf("%s%s", path, File), &Config); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}