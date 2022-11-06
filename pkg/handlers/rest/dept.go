package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/utils"
	"github.com/gin-gonic/gin"
)

type DeptHandler interface {
	GetAllDepts(*gin.Context) (*[]models.Department, error)
	GetDept(*gin.Context) (*models.Department, error)
	AddDept(*gin.Context) (*models.Department, []utils.Error)
	UpdateDept(*gin.Context) (*models.Department, []utils.Error)
	DeleteDept(*gin.Context) error
}

type deptHandler struct {
	dept services.DeptService
}

func NewDeptHandler(dept services.DeptService) DeptHandler {
	return deptHandler{dept}
}

func (d deptHandler) GetAllDepts(ctx *gin.Context) (*[]models.Department, error) {
	return d.dept.GetAllDepts()
}

func (d deptHandler) GetDept(ctx *gin.Context) (*models.Department, error) {
	dID := ctx.Param("id")

	id, err := strconv.Atoi(dID)
	if err != nil {
		return nil, err
	}

	return d.dept.GetDept(id)
}

func (d deptHandler) AddDept(ctx *gin.Context) (*models.Department, []utils.Error) {
	var dept models.Department

	if err := ctx.ShouldBindJSON(&dept); err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	usr := ctx.MustGet("user").(*models.User)

	if err := d.dept.AddDept(usr.RoleID, &dept); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not create a department")
	}

	return &dept, nil
}

func (d deptHandler) UpdateDept(ctx *gin.Context) (*models.Department, []utils.Error) {
	var dept models.Department
	dID := ctx.Param("id")

	id, err := strconv.Atoi(dID)
	if err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Invalid parameter ID")
	}

	if err := ctx.ShouldBindJSON(&dept); err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	usr := ctx.MustGet("user").(*models.User)

	if err := d.dept.UpdateDept(usr.RoleID, id, &dept); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not create a department")
	}

	return &dept, nil
}

func (d deptHandler) DeleteDept(ctx *gin.Context) error {
	dID := ctx.Param("id")

	id, err := strconv.Atoi(dID)
	if err != nil {
		return errors.New("invalid parameter ID")
	}

	usr := ctx.MustGet("user").(*models.User)

	return d.dept.DeleteDept(usr.RoleID, id)
}
