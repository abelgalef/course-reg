package rest

import (
	"net/http"
	"strconv"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ConstHandler interface {
	GetCurrentConstraint(*gin.Context) (*models.Constraint, error)
	CreateConstraint(*gin.Context) (*models.Constraint, []utils.Error)
	GetDeptConstraintHistory(*gin.Context) (*[]models.Constraint, error)
	GetAllConstraintHistory(*gin.Context) (*[]models.Constraint, error)
}

type constHandler struct {
	constService services.ConstService
}

func NewConstHandler(con services.ConstService) ConstHandler {
	return constHandler{con}
}

func (c constHandler) GetCurrentConstraint(ctx *gin.Context) (*models.Constraint, error) {
	dID := ctx.Param("id")

	id, err := strconv.Atoi(dID)
	if err != nil {
		return nil, err
	}

	return c.constService.GetCurrentConstraint(id)
}

func (c constHandler) GetAllConstraintHistory(ctx *gin.Context) (*[]models.Constraint, error) {
	return c.constService.GetAllConstraintHistory()
}

func (c constHandler) GetDeptConstraintHistory(ctx *gin.Context) (*[]models.Constraint, error) {
	dID := ctx.Param("id")

	id, err := strconv.Atoi(dID)
	if err != nil {
		return nil, err
	}

	return c.constService.GetDeptConstraintHistory(id)
}

func (c constHandler) CreateConstraint(ctx *gin.Context) (*models.Constraint, []utils.Error) {
	var con models.Constraint

	if err := ctx.ShouldBindJSON(&con); err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	usr := ctx.MustGet("user").(*models.User)
	con.CreatedBy = usr.ID

	if err := c.constService.CreateConstraint(&con, usr.RoleID); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not create constraint.")
	}

	return &con, nil
}
