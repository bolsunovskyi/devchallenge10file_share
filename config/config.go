package config

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"time"
)

//Mongo type for mongo config settings
type Mongo struct {
	Host	string
	Port	uint
	DB	string
	Timeout	time.Duration
}

//AppConfig type for app configuration
type AppConfig struct {
	Secret		string
	Port		uint
	Mongo		Mongo
	DataFolder	string
}

//File name of the config file
var File = "config.toml"
//Config var for app config
var Config AppConfig

//Read func for reading configuration file
func Read(path string) bool {
	if _, err := toml.DecodeFile(fmt.Sprintf("%s%s", path, File), &Config); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}