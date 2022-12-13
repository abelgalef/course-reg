package repos

import (
	"fmt"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
	"gorm.io/gorm"
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
	if err := dr.db.Connection.Find(&depts).Error; err != nil {
		return nil, err
	}

	return &depts, nil
}

func (dr deptRepo) GetDept(id int) (*models.Department, error) {
	var dept models.Department
	if err := dr.db.Connection.Preload("Courses").Preload("Constraints").First(&dept, id).Error; err != nil {
		return nil, err
	}

	return &dept, nil
}

func (dr deptRepo) AddDept(dept *models.Department) error {
	err := dr.db.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dept).Error; err != nil {
			return err
		}

		if err := tx.Create(createRightModels(*dept)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
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

func createRightModels(dept models.Department) []models.Right {
	var rights []models.Right

	return append(rights,
		models.Right{Tag: dept.DeptCode + "/CONST-WRITE", Desc: fmt.Sprintf("This permission grants the role to add and update constraints on %s department.", dept.Name)},
		models.Right{Tag: dept.DeptCode + "/ADD-COURSE", Desc: fmt.Sprintf("This permission grants the role to teach a particular course in the %s department.", dept.Name)},
		models.Right{Tag: dept.DeptCode + "/COURSE-WRITE", Desc: fmt.Sprintf("This permission grants the role to add and update courses for %s department.", dept.Name)},
	)

}
