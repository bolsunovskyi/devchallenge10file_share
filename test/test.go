package test

import (
	"file_share/config"
	"fmt"
	"file_share/database"
	"testing"
)

func InitConfig(configPath string) {
	config.File = "config_test.toml"
	if err := config.Read(configPath); !err {
		fmt.Println("Unable to load config")
	}
}

func TearDown(t *testing.T) {
	session, db, err := database.GetSession()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer  session.Close()

	if err := db.DropDatabase(); err != nil {
		t.Error(err.Error())
	}
}
