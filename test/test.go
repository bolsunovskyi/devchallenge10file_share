package test

import (
	"file_share/config"
	"fmt"
	"file_share/database"
	"testing"
)

//InitConfig helper function for tests, it opens config file
func InitConfig(configPath string) {
	config.File = "config_test.toml"
	if err := config.Read(configPath); !err {
		fmt.Println("Unable to load config")
	}
}

//TearDown drops test database
func TearDown(t *testing.T) {
	session, db, err := database.GetSession()
	if err != nil {
		t.Error(err.Error())
		return
	}
	session.Fsync(false)
	if err := db.DropDatabase(); err != nil {
		t.Error(err.Error())
	}
	session.Close()
}
