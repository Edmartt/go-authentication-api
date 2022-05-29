package transport

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Edmartt/go-authentication-api/internal/users/data"
	"github.com/Edmartt/go-authentication-api/internal/users/models"
)

var mockedCreate func(user models.User) string
var mockedFind func(id string) *models.User

type mockedRepo struct{}

func(m mockedRepo) Create(user models.User) string{
	return mockedCreate(user)
}

func(m mockedRepo) Find(id string) *models.User{
	return mockedFind(id)
}

func TestSignup(t *testing.T) {

	h := Handlers{}

	data.RepoAccessInterface = mockedRepo{}

	Newuser := models.User{
		Id: "1",
		Username: "edmartt",
		Password: "12345678",
	}

	mockedCreate = func(user models.User) string {
		user = Newuser
		return user.Id
	}

	body := strings.NewReader(`{"username":"edmartt","password":"12345678"}`)

	req, err := http.NewRequest("POST", "/api/v1/public/signup", body)

	if err != nil{
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(h.Signup)

	//req.Header.Set("Content-Type", "application/json")

	handler.ServeHTTP(resp, req)
	expected := `{"status":"User Created"}`

	if status := resp.Code; status != http.StatusCreated{
		t.Errorf("Expected 201 and got %d", resp.Code)
	}


	//We need to cut the linebreak that response body has with the strings.TrimRight method, because if not, the strings will be different
	if strings.TrimRight(resp.Body.String(), "\n") != expected{
		t.Errorf("unexpected body: got %v and want %v", resp.Body, expected)
	}
}

func TestLogin(t *testing.T){

	dbUser := models.User{
		Id: "1",
		Username: "edmartt",
	}

	mockedFind = func(id string) models.User {
		return dbUser
	}

	h := Handlers{}
	h.userRepo = mockedRepo{}
	body := strings.NewReader(`{"username":"edmartt", "password":"12345678"}`)

	req, err := http.NewRequest("POST", "api/v1/public/login", body)

	if err != nil{
		t.Errorf("Request error: %v", err.Error())
	}

	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(h.Login)

	handler.ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK{
		t.Errorf("Bad Status. Got %d and want %d", status, http.StatusOK)
	}
}
