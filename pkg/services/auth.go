package services

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
	"github.com/abelgalef/course-reg/pkg/utils"
)

type AuthService interface {
	AuthenticateUser(*models.LoginData) (*models.LoginResponse, bool)
	AddNewUser(*models.User) (*models.RegResponse, []utils.Error)
}

type authService struct {
	repo repos.UserRepo
	ts   JWTTokenService
}

func NewAuthService(userRepo repos.UserRepo, tokenSerivce JWTTokenService) AuthService {
	return &authService{repo: userRepo, ts: tokenSerivce}
}

func (as *authService) AuthenticateUser(data *models.LoginData) (*models.LoginResponse, bool) {
	user, err := as.repo.GetUserWithEmail(strings.TrimSpace(data.Email))
	if err != nil {
		return nil, false
	}

	hash := sha256.Sum256([]byte(strings.TrimSpace(data.Pass)))

	if user.Password == base64.URLEncoding.EncodeToString(hash[:]) {
		token, err := as.ts.GenerateToken(*user)
		if err != nil {
			return nil, false
		}

		return &models.LoginResponse{
			User:  *user,
			Token: token,
		}, true
	}
	return nil, false
}

func (as *authService) AddNewUser(user *models.User) (*models.RegResponse, []utils.Error) {
	hash := sha256.Sum256([]byte(strings.TrimSpace(user.Password)))
	user.Password = base64.URLEncoding.EncodeToString(hash[:])

	if err := as.repo.AddNewUser(user); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not create your account, Please try again.")
	}

	token, err := as.ts.GenerateToken(*user)
	if err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Your account has been created, Please try loging in.")
	}

	return &models.RegResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		RoleID:    user.RoleID,
		Token:     token,
	}, nil
}
