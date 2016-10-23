package config

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

type Mongo struct {
	Host	string
	Port	uint
	DB	string
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