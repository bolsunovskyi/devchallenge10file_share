package jwt

import (
	"testing"
	"file_share/repository/user"
	"file_share/test"
)

func init() {
	test.InitConfig("../")
}

func TestCreateToken(t *testing.T) {
	defer test.TearDown(t)

	tokenUser, err := user.CreateUser("foo", "bar", "foo22@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	token, err := CreateToken(*tokenUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	validUser, err := CheckToken(*token)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if validUser.ID != tokenUser.ID {
		t.Error("User not equals")
		t.Error(validUser.ID)
		t.Error(tokenUser.ID)

		t.Error(validUser.Email)
		t.Error(tokenUser.Email)
	}
}

func TestCheckToken(t *testing.T) {
	_, err := CheckToken("fff")
	if err == nil {
		t.Error("No error on bad token")
		return
	}
}

func TestCheckTokenWrongSign(t *testing.T) {
	defer test.TearDown(t)

	badToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6Ijk0NTNjNGZhMjg4ODM3MzQxYTBhYWEwYWRlZDE4YjllZGFkZDc1MGQxYmFiYjAwYjgzMWM5OTY5YWM4ZWFlNGViMzIwMWVkZWVkM2M2YmY2In0.eyJhdWQiOiI1MSIsImp0aSI6Ijk0NTNjNGZhMjg4ODM3MzQxYTBhYWEwYWRlZDE4YjllZGFkZDc1MGQxYmFiYjAwYjgzMWM5OTY5YWM4ZWFlNGViMzIwMWVkZWVkM2M2YmY2IiwiaWF0IjoxNDc0MTQ4NDY0LCJuYmYiOjE0NzQxNDg1MjQsImV4cCI6MTQ3NDE1NTY2NCwic3ViIjoiNTciLCJzY29wZXMiOlsiXCJcXFwiYXV0aFxcXCJcIiIsIlwiXFxcImF1dGhcXFwiXCIiXX0.hbVww1lc6lxlyj1HnSLUJUwbpYKUxnHmMdQig-DPyCBF6rG0uiiMHm08Qouc-KHmvZYKOsPgfiXqQMTfRI1jd2R78j2c60YG3voncBu2PxkYvxWLbIeCA6przrmTXWW5kDchIVf1uQuGgfrAis7rqsd0p51CMrxfvheEiAv2HU8"

	_, err := CheckToken(badToken)
	if err == nil {
		t.Error("No error on bad token")
		return
	}
}