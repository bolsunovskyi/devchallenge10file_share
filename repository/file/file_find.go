package file

import (
	"file_share/models"
	"file_share/database"
	"gopkg.in/mgo.v2/bson"
)

func FindByName(name string) (*models.File, error) {
	session, db, err := database.GetSession()
	defer session.Close()

	if err != nil {
		return nil, err
	}

	findFile := models.File{}
	err = db.C(Collection).Find(bson.M{"name": name}).One(&findFile)
	if err != nil {
		return nil, err
	}

	return &findFile, nil
}

func FindByID(fileID bson.ObjectId) (*models.File, error) {
	session, db, err := database.GetSession()
	defer session.Close()

	if err != nil {
		return nil, err
	}

	findFile := models.File{}
	err = db.C(Collection).Find(bson.M{"_id": fileID}).One(&findFile)
	if err != nil {
		return nil, err
	}

	return &findFile, nil
}
