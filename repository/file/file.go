package file

import (
	"io"
	"time"
	"fmt"
	"file_share/config"
	"file_share/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
	"os"
	"file_share/database"
	"errors"
)

var Collection string = "file"

func checkParent(parentID *string) (*models.File, error) {
	var parent *models.File
	if parentID != nil {
		if !bson.IsObjectIdHex(*parentID) {
			return nil, errors.New("Wrong parent ID")
		}

		parentIDObj := bson.ObjectIdHex(*parentID)
		var err error
		parent, err = FindByID(parentIDObj)
		if err != nil {
			return nil, err
		}
	}

	return parent, nil
}

func createRealPath() (realName string, realPath string, fullPath string) {
	now := time.Now()

	realName = uuid.NewV4().String()
	realPath = fmt.Sprintf(
		"%s/%d/%d/%d/%d",
		config.Config.DataFolder,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour())

	os.MkdirAll(realPath, 0777)
	fullPath = fmt.Sprintf("%s/%s", realPath, realName)
	return
}

func UploadFile(reader io.Reader, fileName string, parentID *string) (uploadedFile *models.File, err error) {

	parent, err := checkParent(parentID)
	if err != nil {
		return
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
	}

	if parent != nil {
		uploadedFile.ParentID = parent.ID
	}

	session, db, err := database.GetSession()
	defer session.Close()

	err = db.C(Collection).Insert(uploadedFile)

	if err != nil {
		return
	}

	return uploadedFile, nil
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
