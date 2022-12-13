package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/utils"
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	GetRoles(*gin.Context) (*[]models.Role, error)
	GetRole(*gin.Context) (*models.Role, error)
	AddNewRole(*gin.Context) (*models.Role, []utils.Error)
	GivePerm(*gin.Context) error
	RevokPerm(*gin.Context) error
	GetUsers(*gin.Context) (*[]models.User, error)
	AddUserToRole(*gin.Context) error
	DeleteRole(*gin.Context) error
}

type roleHandler struct {
	roleSer services.RoleService
}

func NewRoleHandler(role services.RoleService) RoleHandler {
	return roleHandler{roleSer: role}
}

func (rh roleHandler) GetRoles(ctx *gin.Context) (*[]models.Role, error) {
	return rh.roleSer.GetRoles()
}

func (rh roleHandler) GetRole(ctx *gin.Context) (*models.Role, error) {
	qID := ctx.Param("id")

	id, err := strconv.Atoi(qID)
	if err != nil {
		return nil, err
	}

	return rh.roleSer.GetRole(id)
}

func (rh roleHandler) DeleteRole(ctx *gin.Context) error {
	qID := ctx.Param("id")

	id, err := strconv.Atoi(qID)
	if err != nil {
		return err
	}

	return rh.roleSer.DeleteRole(id)
}

func (rh roleHandler) AddNewRole(ctx *gin.Context) (*models.Role, []utils.Error) {
	var role models.Role

	if err := ctx.ShouldBindJSON(&role); err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	usr := ctx.MustGet("user").(*models.User)

	if err := rh.roleSer.CreateNewRole(usr.ID, &role); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not create a new role")
	}

	return &role, nil
}

func (rh roleHandler) GivePerm(ctx *gin.Context) error {
	roID := ctx.Param("role_id")
	riID := ctx.Param("right_id")

	fmt.Println(roID, riID)

	rightID, err := strconv.Atoi(riID)
	if err != nil {
		return err
	}

	roleID, err := strconv.Atoi(roID)
	if err != nil {
		return err
	}

	usr := ctx.MustGet("user").(*models.User)

	return rh.roleSer.GivePermissionToRole(usr.ID, roleID, rightID)
}

func (rh roleHandler) RevokPerm(ctx *gin.Context) error {
	roID := ctx.Param("role_id")
	riID := ctx.Param("right_id")

	rightID, err := strconv.Atoi(riID)
	if err != nil {
		return err
	}

	roleID, err := strconv.Atoi(roID)
	if err != nil {
		return err
	}

	usr := ctx.MustGet("user").(*models.User)

	return rh.roleSer.RevokeRightFromRole(usr.ID, roleID, rightID)
}

func (rh roleHandler) GetUsers(ctx *gin.Context) (*[]models.User, error) {
	roID := ctx.Param("role_id")
	roleID, err := strconv.Atoi(roID)
	if err != nil {
		return nil, err
	}

	return rh.roleSer.GetUsers(roleID)
}

func (rh roleHandler) AddUserToRole(ctx *gin.Context) error {
	roID := ctx.Param("role_id")
	roleID, err := strconv.Atoi(roID)
	if err != nil {
		return err
	}

	uID := ctx.Param("user_id")
	usrID, err := strconv.Atoi(uID)
	if err != nil {
		return err
	}

	usr := ctx.MustGet("user").(*models.User)

	return rh.roleSer.AddUserToRole(usr.ID, roleID, usrID)
}
