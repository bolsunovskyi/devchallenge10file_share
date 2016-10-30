package file

import (
	"testing"
	"file_share/test"
	"file_share/repository/user"
)

func TestSearchFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo2221@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = CreateFolder("images", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err := SearchFiles("imag", appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) == 0 {
		t.Error("No files found")
		return
	}
}
