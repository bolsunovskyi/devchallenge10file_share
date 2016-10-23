package file

import (
	"file_share/test"
	"testing"
	"os"
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
	os.RemoveAll("_data")

	test.TearDown(t)
}


