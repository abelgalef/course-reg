package repos

import (
	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
)

type ConstRepo interface {
	CreateConstraint(models.Constraint) error
	GetLatestConstraintForDept(int) (*models.Constraint, error)
	GetDeptConstraintHistory(int) (*[]models.Constraint, error)
	GetAllConstraintHistory() (*[]models.Constraint, error)
}

type constRepo struct {
	db *storage.MainDB
}

func NewConstRepository(db *storage.MainDB) ConstRepo {
	return constRepo{db}
}

func (c constRepo) CreateConstraint(con models.Constraint) error {
	return c.db.Connection.Create(&con).Error
}

func (c constRepo) GetLatestConstraintForDept(deptID int) (*models.Constraint, error) {
	var con models.Constraint

	if err := c.db.Connection.Where("dept_id = ?", deptID).Last(&con).Error; err != nil {
		return nil, err
	}

	return &con, nil
}

func (c constRepo) GetDeptConstraintHistory(deptID int) (*[]models.Constraint, error) {
	var consts []models.Constraint
	if err := c.db.Connection.Where("dept_id = ?", deptID).Find(&consts).Error; err != nil {
		return nil, err
	}

	return &consts, nil
}

func (c constRepo) GetAllConstraintHistory() (*[]models.Constraint, error) {
	var consts []models.Constraint
	if err := c.db.Connection.Find(&consts).Error; err != nil {
		return nil, err
	}

	return &consts, nil
}
