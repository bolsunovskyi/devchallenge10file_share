package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type File struct {
	ID		bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	Name		string		`json:"name"`
	ParentID	bson.ObjectId	`bson:"parentID,omitempty" json:"parent_id,omitempty"`
	FileSize	uint		`bson:"fileSize,omitempty" json:"file_size"`
	Created		time.Time	`json:"created_at"`
	Updated		time.Time	`json:"updated_at"`
	IsDir		bool		`bson:"isDir" json:"is_folder"`
	RealPath	string		`bson:"realPath,omitempty" json:"-"`
	RealName	string		`bson:"realName,omitempty" json:"-"`
	UserID		bson.ObjectId	`bson:"userID" json:"-"`
}
