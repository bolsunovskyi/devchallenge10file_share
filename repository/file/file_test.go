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
	if appUser, err := user.CreateUser("foo", "bar", "foo1@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		fileName := "tmp_file.txt"
		if appFile, err := os.Create(fileName); err != nil {
			t.Error(err.Error())
		} else {
			appFile.Write([]byte("Test Data"))

			_, err = UploadFile(appFile, fileName, nil, appUser)
			if err != nil {
				t.Error(err.Error())
			}

			appFile.Close();
			os.Remove(fileName)
			os.RemoveAll(config.Config.DataFolder)
		}
	}

	test.TearDown(t)
}

func TestCreateFolder(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo2@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if folder, err := CreateFolder("images1", nil, appUser); err != nil {
			t.Error(err.Error())
		} else {
			folderID := folder.ID.Hex()
			_, err = CreateFolder("summer2016", &folderID, appUser)
			if err != nil {
				t.Error(err.Error())
			}
		}
	}

	test.TearDown(t)
}

func TestListFiles(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo3@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
		return
	} else {

		if _, err = CreateFolder("images2", nil, appUser); err != nil {
			t.Error(err.Error())
		} else {
			if files, err := ListFiles(nil, appUser); err != nil {
				t.Error(err.Error())
			} else {
				if len(files) == 0 {
					t.Error("Empty file list")
				}
			}
		}
	}

	test.TearDown(t)
}

func TestDeleteFile(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo4@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if folder, err := CreateFolder("images3", nil, appUser); err != nil {
			t.Error(err.Error())
		} else {
			folderID := folder.ID.Hex()

			err = DeleteFile(folderID, appUser)
			if err != nil {
				t.Error(err.Error())
			}

			files, err := ListFiles(nil, appUser)
			if err != nil {
				t.Error(err.Error())
			}

			if len(files) > 0 {
				t.Error("File list is not empty")
			}
		}
	}

	test.TearDown(t)
}

func TestRenameFile(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo5@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if folder, err := CreateFolder("images4", nil, appUser); err != nil {
			t.Error(err.Error())
		} else {
			if _, err = RenameFile(folder.ID.Hex(), "photos", appUser); err != nil {
				t.Error(err.Error())
			} else {

				if appFile, err := FindByID(folder.ID); err != nil {
					t.Error(err.Error())
				} else {
					if appFile.Name != "photos" {
						t.Error("File is not updated")
					}
				}
			}
		}
	}

	test.TearDown(t)
}