package file

import (
	"file_share/models"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/satori/go.uuid"
	"fmt"
	"file_share/config"
	"errors"
	"os"
)

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
