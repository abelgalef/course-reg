package repos

import (
	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
)

type DeptRepo interface {
	GetAllDepts() (*[]models.Department, error)
	GetDept(int) (*models.Department, error)
	AddDept(*models.Department) error
	UpdateDept(int, *models.Department) error
	DeleteDept(int) error
}

type deptRepo struct {
	db *storage.MainDB
}

func NewDeptRepo(db *storage.MainDB) DeptRepo {
	return &deptRepo{db}
}

func (dr deptRepo) GetAllDepts() (*[]models.Department, error) {
	var depts []models.Department
	if err := dr.db.Connection.First(&depts).Error; err != nil {
		return nil, err
	}

	return &depts, nil
}

func (dr deptRepo) GetDept(id int) (*models.Department, error) {
	var dept models.Department
	if err := dr.db.Connection.First(&dept, id).Error; err != nil {
		return nil, err
	}

	return &dept, nil
}

func (dr deptRepo) AddDept(dept *models.Department) error {
	if err := dr.db.Connection.Create(dept).Error; err != nil {
		return err
	}
	return nil
}

func (dr deptRepo) UpdateDept(id int, dept *models.Department) error {
	if err := dr.db.Connection.Model(&models.Department{ID: id}).Updates(dept).Error; err != nil {
		return err
	}

	return nil
}

func (dr deptRepo) DeleteDept(id int) error {
	if err := dr.db.Connection.Delete(&models.Department{ID: id}).Error; err != nil {
		return err
	}

	return nil
}
