package file

import (
	"io"
	"time"
	"file_share/models"
	"gopkg.in/mgo.v2/bson"
	"os"
	"file_share/database"
	"errors"
)

var Collection string = "file"

func UploadFile(reader io.Reader, fileName string, parentID *string, appUser *models.User) (uploadedFile *models.File, err error) {

	parent, err := checkParent(parentID)
	if err != nil {
		return
	}

	if sameName, err := FindByName(fileName); err == nil && !sameName.IsDir {
		return nil, errors.New("File already exists")
	}

	realName, realPath, fullPath := createRealPath()

	file, err := os.Create(fullPath)
	if err != nil {
		return
	}

	size, err := io.Copy(file, reader)
	if err != nil {
		return
	}

	uploadedFile = &models.File{
		ID:		bson.NewObjectId(),
		FileSize:	uint(size),
		RealName:	realName,
		RealPath:	realPath,
		IsDir:		false,
		Name:		fileName,
		Created:	time.Now(),
		Updated:	time.Now(),
		UserID:		appUser.ID,
	}

	if parent != nil {
		uploadedFile.ParentID = parent.ID
	}

	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	err = db.C(Collection).Insert(uploadedFile)

	if err != nil {
		return
	}

	return uploadedFile, nil
}

func CreateFolder(fileName string, parentID *string, appUser *models.User) (uploadedFile *models.File, err error) {
	parent, err := checkParent(parentID)
	if err != nil {
		return
	}

	if sameName, err := FindByName(fileName); err == nil && sameName.IsDir {
		return nil, errors.New("File already exists")
	}


	uploadedFile = &models.File{
		ID:		bson.NewObjectId(),
		IsDir:		true,
		Name:		fileName,
		Created:	time.Now(),
		Updated:	time.Now(),
		UserID:		appUser.ID,
	}

	if parent != nil {
		uploadedFile.ParentID = parent.ID
	}

	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	err = db.C(Collection).Insert(uploadedFile)

	if err != nil {
		return
	}

	return uploadedFile, nil
}

//ListFiles returns all files by folder
func ListFiles(folderID *string) (files []models.File, err error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var parent *models.File

	if parent, err = checkParent(folderID); err == nil && parent != nil {
		folderOBJID := bson.ObjectIdHex(*folderID)
		err = db.C(Collection).Find(bson.M{"parentID": folderOBJID}).All(&files)
	} else if err == nil && parent == nil {
		err = db.C(Collection).Find(bson.M{"parentID": bson.M{"$exists": false}}).All(&files)
	}

	return
}
