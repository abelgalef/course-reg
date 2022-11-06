package repos

import (
	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
)

type RoleRepo interface {
	CreateRole(*models.Role) error
	UpdateRole(int, *models.Role) error
	DeleteRole(int) error
	GetRoleByID(int) (*models.Role, error)
	GetRoleByTag(string) (*models.Role, error)
	GetAllRoles() (*[]models.Role, error)
}

type roleRepo struct {
	db *storage.MainDB
}

func NewRoleRepository(db *storage.MainDB) RoleRepo {
	return roleRepo{db}
}

func (r roleRepo) CreateRole(role *models.Role) error {
	return r.db.Connection.Create(role).Error
}

func (r roleRepo) UpdateRole(id int, role *models.Role) error {
	return r.db.Connection.Model(&models.Role{ID: id}).Updates(role).Error
}

func (r roleRepo) DeleteRole(id int) error {
	return r.db.Connection.Delete(&models.Role{ID: id}).Error
}

func (r roleRepo) GetRoleByID(id int) (*models.Role, error) {
	var role models.Role
	if err := r.db.Connection.Preload("Rights").Preload("Users").First(&role, id).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r roleRepo) GetRoleByTag(tag string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Connection.Preload("Rights").Preload("Users").Where("tag = ?", tag).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r roleRepo) GetAllRoles() (*[]models.Role, error) {
	var roles []models.Role
	if err := r.db.Connection.Preload("Rights").Preload("Users").Find(&roles).Error; err != nil {
		return nil, err
	}

	return &roles, nil
}
