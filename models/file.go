package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type File struct {
	ID		bson.ObjectId	`bson:"_id,omitempty"`
	Name		string
	ParentID	bson.ObjectId	`bson:"parentID,omitempty"`
	FileSize	uint		`bson:"filesize,omitempty"`
	Created		time.Time
	Updated		time.Time
	IsDir		bool
	RealPath	string		`bson:"realpath,omitempty"`
	RealName	string		`bson:"realname,omitempty"`
	UserID		bson.ObjectId
}
