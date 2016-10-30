package file

import (
	"file_share/models"
	"file_share/database"
	"gopkg.in/mgo.v2/bson"
)

//FindByNameAndDir look for file by name and parent folder
func FindByNameAndDir(name string, parent *models.File) (*models.File, error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	findFile := models.File{}
	if parent == nil {
		err = db.C(Collection).Find(bson.M{"name": name}).One(&findFile)
		if err != nil {
			return nil, err
		}
	} else {
		err = db.C(Collection).Find(bson.M{"name": name, "parentID": parent.ID}).One(&findFile)
		if err != nil {
			return nil, err
		}
	}

	return &findFile, nil
}

//FindByID look for file by it's ID
func FindByID(fileID bson.ObjectId) (*models.File, error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	findFile := models.File{}
	err = db.C(Collection).Find(bson.M{"_id": fileID}).One(&findFile)
	if err != nil {
		return nil, err
	}

	return &findFile, nil
}

//FindByIDUser look for file by it's id and owner
func FindByIDUser(fileID bson.ObjectId, userID bson.ObjectId) (*models.File, error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	findFile := models.File{}
	err = db.C(Collection).Find(bson.M{"_id": fileID, "userID": userID}).One(&findFile)
	if err != nil {
		return nil, err
	}

	return &findFile, nil
}
