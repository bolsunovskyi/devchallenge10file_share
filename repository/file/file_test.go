package file

import (
	"file_share/test"
	"testing"
	"os"
	"file_share/config"
)

func init() {
	test.InitConfig("../../")
}

func TestUploadFile(t *testing.T) {
	fileName := "tmp_file.txt"
	appFile, err := os.Create(fileName)
	if err != nil {
		t.Error(err.Error())
	}

	appFile.Write([]byte("Test Data"))

	_, err = UploadFile(appFile, fileName, nil)
	if err != nil {
		t.Error(err.Error())
	}

	appFile.Close()
	os.Remove(fileName)
	os.RemoveAll(config.Config.DataFolder)

	test.TearDown(t)
}

func TestCreateFolder(t *testing.T) {
	folder, err := CreateFolder("images", nil)
	if err != nil {
		t.Error(err.Error())
	}

	folderID := folder.ID.Hex()
	_, err = CreateFolder("summer2016", &folderID)
	if err != nil {
		t.Error(err.Error())
	}

	test.TearDown(t)
}

func TestListFiles(t *testing.T) {
	TestCreateFolder(t)

	_, err := ListFiles(nil)
	if err != nil {
		t.Error(err.Error())
	}

	test.TearDown(t)
}

