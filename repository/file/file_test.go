package file

import (
	"file_share/test"
	"testing"
	"os"
	"file_share/config"
	"file_share/repository/user"
	"strings"
)

func init() {
	test.InitConfig("../../")
}

func TestCheckParent(t *testing.T) {
	defer  test.TearDown(t)
	defer os.RemoveAll(config.Config.DataFolder)

	parent := "sfsadsa"
	if _, err := checkParent(&parent, nil); err == nil {
		t.Error("No error on wrong parent")
	}

	appUser, err := user.CreateUser("foo", "bar", "foo5@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	uFile, err := UploadFile(strings.NewReader("hello"), "sadsad", nil, appUser)
	if err != nil {
		t.Error(err)
		return
	}

	parent = uFile.ID.Hex()
	if _, err := checkParent(&parent, appUser); err == nil {
		t.Error("No error on wrong parent")
	}
}

