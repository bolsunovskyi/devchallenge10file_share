package file

import (
	"testing"
	"file_share/config"
	"os"
	"file_share/test"
	"file_share/repository/user"
	"strings"
)

func TestCreateZipArchive(t *testing.T) {
	defer test.TearDown(t)
	defer os.RemoveAll(config.Config.DataFolder)

	appUser, err := user.CreateUser("foo", "bar", "foo1@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := CreateFolder("folder1", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	parent := folder1.ID.Hex()
	_, err = UploadFile(strings.NewReader("dqwdqwdw"), "dqwdqw.txt", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder2, err := CreateFolder("folder1", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
	parent2 := folder2.ID.Hex()

	_, err = UploadFile(strings.NewReader("dqwdqwdw"), "dqwdqw.txt", &parent2, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, _, err = CreateZipArchive(*folder1, *appUser)
	if err != nil {
		t.Error(err.Error())
	}

}
