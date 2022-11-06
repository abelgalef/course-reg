package services

import (
	"errors"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
)

type ConstService interface {
	GetCurrentConstraint(int) (*models.Constraint, error)
	CreateConstraint(*models.Constraint, int) error
	GetDeptConstraintHistory(int) (*[]models.Constraint, error)
	GetAllConstraintHistory() (*[]models.Constraint, error)
}

type constService struct {
	constRepo repos.ConstRepo
	rights    repos.RightsRepo
}

func NewConstService(constRepo repos.ConstRepo, rights repos.RightsRepo) ConstService {
	return constService{constRepo, rights}
}

func (c constService) GetCurrentConstraint(deptID int) (*models.Constraint, error) {
	return c.constRepo.GetLatestConstraintForDept(deptID)
}

func (c constService) CreateConstraint(con *models.Constraint, requestedBy int) error {
	if c.rights.RoleHasDeptRight(requestedBy, con.DeptID, "/CONST-WRITE") {
		return c.constRepo.CreateConstraint(*con)
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (c constService) GetDeptConstraintHistory(deptID int) (*[]models.Constraint, error) {
	return c.constRepo.GetDeptConstraintHistory(deptID)
}

func (c constService) GetAllConstraintHistory() (*[]models.Constraint, error) {
	return c.constRepo.GetAllConstraintHistory()
}
