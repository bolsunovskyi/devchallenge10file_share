package file

import (
	"io"
	"time"
	"file_share/models"
	"gopkg.in/mgo.v2/bson"
	"os"
	"file_share/database"
	"errors"
	"regexp"
)

var Collection string = "file"

func UploadFile(reader io.Reader, fileName string, parentID *string, appUser *models.User) (uploadedFile *models.File, err error) {

	parent, err := checkParent(parentID, appUser)
	if err != nil {
		return
	}

	if sameName, err := FindByName(fileName); err == nil && !sameName.IsDir {
		//TODO: check parent
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
	parent, err := checkParent(parentID, appUser)
	if err != nil {
		return
	}

	if sameName, err := FindByName(fileName); err == nil && sameName.IsDir {
		//TODO: check parent
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
func ListFiles(folderID *string, appUser *models.User) (files []models.File, err error) {
	files = make([]models.File, 0)
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var parent *models.File

	if parent, err = checkParent(folderID, appUser); err == nil && parent != nil {
		folderOBJID := bson.ObjectIdHex(*folderID)
		err = db.C(Collection).Find(bson.M{"parentID": folderOBJID, "userID": appUser.ID}).All(&files)
	} else if err == nil && parent == nil {
		err = db.C(Collection).Find(bson.M{"parentID": bson.M{"$exists": false}, "userID": appUser.ID}).All(&files)
	}

	return
}

func DeleteFile(fileID string, appUser *models.User) error {
	if !bson.IsObjectIdHex(fileID) {
		return errors.New("Wrong file ID")
	}

	session, db, err := database.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()

	deleteFile, err := FindByID(bson.ObjectIdHex(fileID))
	if err != nil {
		return err
	}

	err = db.C(Collection).Remove(bson.M{"_id": deleteFile.ID, "userID": appUser.ID})
	if err != nil {
		return err
	}

	//TODO: remove all children if folder
	//TODO: remove real file

	return nil
}

func RenameFile(fileID string, fileName string, appUser *models.User) (*models.File, error) {
	if !bson.IsObjectIdHex(fileID) {
		return nil, errors.New("Wrong file ID")
	}

	match, err := regexp.MatchString("[0-9a-zA-Z._]+", fileName)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, errors.New("Wrong file name")
	}

	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	updateFile, err := FindByIDUser(bson.ObjectIdHex(fileID), appUser.ID)
	if err != nil {
		return nil, err
	}

	updateFile.Name = fileName
	updateFile.Updated = time.Now()

	err = db.C(Collection).Update(bson.M{"_id": updateFile.ID, "userID": appUser.ID}, updateFile)
	if err != nil {
		return nil, err
	}

	return updateFile, nil
}

func MoveFile(fileID string, parentID *string, appUser *models.User) (*models.File, error) {
	if !bson.IsObjectIdHex(fileID) || (parentID != nil && !bson.IsObjectIdHex(*parentID)) {
		return nil, errors.New("Wrong ID")
	}

	parent, err := checkParent(parentID, appUser)
	if err != nil {
		return nil, err
	}

	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	updateFile, err := FindByIDUser(bson.ObjectIdHex(fileID), appUser.ID)
	if err != nil {
		return nil, err
	}

	if parent != nil {
		updateFile.ParentID = parent.ID
	} else {
		var parent bson.ObjectId
		updateFile.ParentID = parent
	}

	updateFile.Updated = time.Now()

	err = db.C(Collection).Update(bson.M{"_id": updateFile.ID, "userID": appUser.ID}, updateFile)
	if err != nil {
		return nil, err
	}

	return updateFile, nil
}

func SearchFiles(keyword string, appUser *models.User) (files []models.File, err error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	if len(keyword) < 3 {
		return nil, errors.New("Too short keyword")
	}

	err = db.C(Collection).Find(bson.M{"name": bson.RegEx{
		Pattern: keyword,
		Options: "",
	}, "userID": appUser.ID}).All(&files)

	return
}