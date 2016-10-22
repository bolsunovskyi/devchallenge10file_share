package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type File struct {
	ID		bson.ObjectId
	Name		string
	ParentID	bson.ObjectId
	FileSize	uint
	Created		time.Time
	Updated		time.Time
	IsDir		bool
	RealPath	string
}
