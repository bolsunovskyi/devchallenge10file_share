package middleware

import (
	"testing"
	"net/http"
	"file_share/models"
	"file_share/repository/user"
	"file_share/jwt"
	"net/http/httptest"
	"file_share/test"
)

func init() {
	test.InitConfig("../")
}

func createUserAndToken() (*models.User, *string, error) {
	appUser, err := user.CreateUser("foo", "bar", "foo8@gmail.com", "123456");
	if err != nil {
		return nil, nil, err
	}

	token, err := jwt.CreateToken(*appUser);
	if  err != nil {
		return nil, nil, err
	}

	return appUser, token, nil
}

func TestAuth(t *testing.T) {
	defer test.TearDown(t)

	req, _ := http.NewRequest(
		"GET",  "/",
		nil)

	recorder := httptest.NewRecorder()

	f := func(http.ResponseWriter, *http.Request, *models.User) {}

	Auth(f)(recorder, req)

	if status := recorder.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}

	req, _ = http.NewRequest(
		"GET",  "/?access-token=sdsdsd",
		nil)

	recorder = httptest.NewRecorder()

	Auth(f)(recorder, req)

	if status := recorder.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err)
		return
	}

	req, _ = http.NewRequest(
		"GET",  "/",
		nil)
	req.Header.Add("Access-Token", *token)

	recorder = httptest.NewRecorder()

	Auth(f)(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
