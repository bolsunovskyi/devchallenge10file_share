package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type File struct {
	ID		bson.ObjectId	`bson:"_id,omitempty"`
	Name		string
	ParentID	bson.ObjectId	`bson:"parentID,omitempty"`
	FileSize	uint		`bson:"fileSize,omitempty"`
	Created		time.Time
	Updated		time.Time
	IsDir		bool		`bson:"isDir"`
	RealPath	string		`bson:"realPath,omitempty"`
	RealName	string		`bson:"realName,omitempty"`
	UserID		bson.ObjectId	`bson:"userID"`
}
