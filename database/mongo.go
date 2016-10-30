package database

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"file_share/config"
	"time"
)

func GetSession() (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.DialWithTimeout(
		fmt.Sprintf("%s:%d", config.Config.Mongo.Host, config.Config.Mongo.Port),
		time.Second * config.Config.Mongo.Timeout)
	if err != nil {
		return nil, nil, err
	}

	return session, session.DB(config.Config.Mongo.DB), nil
}
