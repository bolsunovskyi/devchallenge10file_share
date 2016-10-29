package file

import (
	"file_share/test"
	"testing"
	"os"
	"file_share/config"
	"file_share/repository/user"
)

func init() {
	test.InitConfig("../../")
}

func TestUploadFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo1@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	fileName := "tmp_file.txt"
	appFile, err := os.Create(fileName)
	if  err != nil {
		t.Error(err.Error())
		return
	}
	defer appFile.Close();

	appFile.Write([]byte("Test Data"))

	defer os.Remove(fileName)
	defer os.RemoveAll(config.Config.DataFolder)

	_, err = UploadFile(appFile, fileName, nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestCreateFolder(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo2@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images1", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folderID := folder.ID.Hex()
	_, err = CreateFolder("summer2016", &folderID, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestListFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo3@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = CreateFolder("images2", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err := ListFiles(nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) == 0 {
		t.Error("Empty file list")
	}
}

func TestDeleteFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images3", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folderID := folder.ID.Hex()
	err = DeleteFile(folderID, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err := ListFiles(nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) > 0 {
		t.Error("File list is not empty")
	}
}

func TestRenameFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo5@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images4", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = RenameFile(folder.ID.Hex(), "photos", appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	appFile, err := FindByID(folder.ID)
	if  err != nil {
		t.Error(err.Error())
		return
	}
	if appFile.Name != "photos" {
		t.Error("File is not updated")
	}
}

func TestMoveFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo1255@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := CreateFolder("images223s4", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	folder2, err := CreateFolder("image23rs223s4", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	parentID := folder2.ID.Hex()
	_, err = MoveFile(folder1.ID.Hex(), &parentID, appUser)
	if  err != nil {
		t.Error(err.Error())
		return
	}

	_, err = MoveFile(folder1.ID.Hex(), nil, appUser)
	if  err != nil {
		t.Error(err.Error())
		return
	}
}