package rest

import (
	"errors"
	"net/http"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(*gin.Context) (*models.LoginResponse, *utils.Error)
	Register(*gin.Context) (*models.RegResponse, *utils.Error)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(as services.AuthService) AuthHandler {
	return &authHandler{authService: as}
}

func (ah *authHandler) Login(ctx *gin.Context) (*models.LoginResponse, *utils.Error) {
	var data models.LoginData

	if err := ctx.ShouldBindJSON(&data); err != nil {
		// TODO: HANDEL BINDING ERRORS
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	response, authenticated := ah.authService.AuthenticateUser(&data)
	if !authenticated {
		return nil, utils.NewError(http.StatusBadRequest, errors.New("authentication failed"), "Invalid Credentials")
	}

	return response, nil
}

func (ah *authHandler) Register(ctx *gin.Context) (*models.RegResponse, *utils.Error) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		// TODO: HANDEL BINDING ERRORS
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	response, err := ah.authService.AddNewUser(&user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
