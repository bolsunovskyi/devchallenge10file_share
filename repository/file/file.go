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
	"fmt"
	"archive/zip"
	"file_share/config"
)

//Collection file collection name
var Collection = "file"

//UploadFile uploads file and saves it to fs and db
func UploadFile(reader io.Reader, fileName string, parentID *string, appUser *models.User) (uploadedFile *models.File, err error) {

	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	parent, err := checkParent(parentID, appUser)
	if err != nil {
		return
	}

	if sameName, err := FindByNameAndDir(fileName, parent); err == nil && !sameName.IsDir {
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

	err = db.C(Collection).Insert(uploadedFile)
	if err != nil {
		return
	}

	return uploadedFile, nil
}

//CreateFolder creates folder and save it to db
func CreateFolder(fileName string, parentID *string, appUser *models.User) (uploadedFile *models.File, err error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	parent, err := checkParent(parentID, appUser)
	if err != nil {
		return
	}

	if sameName, err := FindByNameAndDir(fileName, parent); err == nil && sameName.IsDir {
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

//DeleteFile deleted file from db and fs
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

	if !deleteFile.IsDir {
		os.Remove(fmt.Sprintf("%s/%s", deleteFile.RealPath, deleteFile.RealName))
	} else {
		var children []models.File
		err := db.C(Collection).Find(bson.M{"parentID": deleteFile.ID}).All(&children)
		if err == nil {
			for _, v := range children {
				DeleteFile(v.ID.Hex(), appUser)
			}
		}
	}

	err = db.C(Collection).Remove(bson.M{"_id": deleteFile.ID, "userID": appUser.ID})
	if err != nil {
		return err
	}

	return nil
}

//RenameFile changes file name
func RenameFile(fileID string, fileName string, appUser *models.User) (*models.File, error) {
	if !bson.IsObjectIdHex(fileID) {
		return nil, errors.New("Wrong file ID")
	}

	match, err := regexp.MatchString("^[0-9a-zA-Z._]+$", fileName)
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

//MoveFile moves file to another folder
func MoveFile(fileID string, parentID *string, appUser *models.User) (*models.File, error) {
	if !bson.IsObjectIdHex(fileID) || (parentID != nil && !bson.IsObjectIdHex(*parentID)) {
		return nil, errors.New("Wrong ID")
	}

	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	parent, err := checkParent(parentID, appUser)
	if err != nil {
		return nil, err
	}

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

//SearchFiles search for files
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

//CreateZipArchive creates zip archive from folder
func CreateZipArchive(folder models.File, appUser models.User) (*string, *int64, error) {
	if !folder.IsDir {
		return nil, nil, errors.New("Only folder may be zipped")
	}

	filePath := fmt.Sprintf("%s/zip_%d", config.Config.DataFolder, time.Now().Unix())

	zipFile, err := os.Create(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = archiveFolder(folder, archive, appUser)
	if err != nil {
		return nil, nil, err
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return nil, nil, err
	}

	size := info.Size()
	return &filePath, &size, nil
}

func archiveFolder(folder models.File, writer *zip.Writer, appUser models.User) error {
	parent := folder.ID.Hex()
	files, err := ListFiles(&parent, &appUser)
	if err != nil {
		return err
	}

	for _, userFile := range files {
		if userFile.IsDir {
			archiveFolder(userFile, writer, appUser)
		} else {
			f, err := writer.Create(userFile.Name)
			if err != nil {
				return err
			}
			uFile, err := os.Open(fmt.Sprintf("%s/%s", userFile.RealPath, userFile.RealName))
			if err != nil {
				return err
			}
			io.Copy(f, uFile)
			uFile.Close()
		}
	}

	return nil
}