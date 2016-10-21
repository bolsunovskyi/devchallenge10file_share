package database

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"file_share/config"
)

func GetSession() (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.Dial(fmt.Sprintf("%s:%d", config.Config.Mongo.Host, config.Config.Mongo.Port))
	if err != nil {
		return nil, nil, err
	}

	return session, session.DB(config.Config.Mongo.DB), nil
}
