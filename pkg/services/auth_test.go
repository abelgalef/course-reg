package services

import (
	"testing"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
)

type mockUser struct{}

func (m mockUser) GetUserWithID(id uint) (*models.User, error) {
	return &models.User{ID: int(id)}, nil
}

func (m mockUser) GetUserWithEmail(email string) (*models.User, error) {
	return &models.User{Email: email, Password: "5F2QlX7sc4dybGobF02ntWaiT_TLBg3Lzf67kxqT_-M="}, nil
}

func (m mockUser) AddNewUser(user *models.User) error {
	return nil
}

func newMockRepo() repos.UserRepo {
	return mockUser{}
}

type jwtS struct{}

func (jwtS) GenerateToken(user models.User) (string, error) {
	return "my token", nil
}

func (jwtS) ValidateToken(t string) (interface{}, bool) {
	if t == "my token" {
		return map[string]interface{}{"id": 1}, true
	}

	return nil, false
}

func newJwtMockService() JWTTokenService {
	return jwtS{}
}

func TestAuthenticateUser(t *testing.T) {
	mockU := newMockRepo()
	mockJWT := newJwtMockService()
	aService := NewAuthService(mockU, mockJWT)

	data := models.LoginData{Email: "my@email.com", Pass: "passw"}

	resp, authnticate := aService.AuthenticateUser(&data)
	if !authnticate {
		t.Fatalf("auth_test.go: could not authnticate user")
	}

	if resp.Token != "my token" {
		t.Fatalf("auth_test.go: did not generate the correct token: genetrated-token = %s", resp.Token)
	}
}

func TestAddNewUser(t *testing.T) {
	mockU := newMockRepo()
	mockJWT := newJwtMockService()
	aService := NewAuthService(mockU, mockJWT)

	u := models.User{}
	resp, err := aService.AddNewUser(&u)
	if err != nil {
		t.Fatalf("auth_test.go: could not add new user: %v", err)
	}

	if resp.Token != "my token" {
		t.Fatalf("auth_test.go: did not generate the correct token: genetrated-token = %s", resp.Token)
	}
}
