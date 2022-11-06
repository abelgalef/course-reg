package services

import (
	"errors"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
)

type DeptService interface {
	GetAllDepts() (*[]models.Department, error)
	GetDept(int) (*models.Department, error)
	AddDept(int, *models.Department) error
	UpdateDept(int, int, *models.Department) error
	DeleteDept(int, int) error
}

type deptService struct {
	dept   repos.DeptRepo
	rights repos.RightsRepo
}

func NewDeptService(dept repos.DeptRepo, rights repos.RightsRepo) DeptService {
	return deptService{dept, rights}
}

func (d deptService) GetAllDepts() (*[]models.Department, error) {
	return d.dept.GetAllDepts()
}

func (d deptService) GetDept(id int) (*models.Department, error) {
	return d.dept.GetDept(id)
}

func (d deptService) AddDept(requestedBy int, dept *models.Department) error {
	if d.rights.RoleHasRight(requestedBy, "DEPT/WRITE") {
		if err := d.dept.AddDept(dept); err != nil {
			return err
		}
		return nil
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (d deptService) UpdateDept(requestedBy, deptID int, dept *models.Department) error {
	if d.rights.RoleHasRight(requestedBy, "DEPT/WRITE") {
		if err := d.dept.UpdateDept(deptID, dept); err != nil {
			return err
		}
		return nil
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (d deptService) DeleteDept(requestedBy, deptID int) error {
	if d.rights.RoleHasRight(requestedBy, "DEPT/DELETE") {
		if err := d.dept.DeleteDept(deptID); err != nil {
			return err
		}
		return nil
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}
