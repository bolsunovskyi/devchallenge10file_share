package database

import (
	"testing"
	"file_share/config"
)

func TestGetSession(t *testing.T) {
	config.File = "config_test.toml"
	if err := config.Read("../"); !err {
		t.Error(err)
		return
	}

	session, _, err := GetSession()
	if err != nil {
		t.Error(err.Error())
		return
	}
	session.Close()

	port := config.Config.Mongo.Port
	config.Config.Mongo.Port = 64011

	_, _, err = GetSession()
	if err == nil {
		t.Error("No error on wrong config")
	}

	config.Config.Mongo.Port = port
}
