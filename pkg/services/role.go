package services

import (
	"errors"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
)

type RoleService interface {
	GetRoles() (*[]models.Role, error)
	GetRole(int) (*models.Role, error)
	CreateNewRole(int, *models.Role) error
	GivePermissionToRole(int, int, int) error
	RevokeRightFromRole(int, int, int) error
}

type roleService struct {
	rightRepo repos.RightsRepo
	roleRepo  repos.RoleRepo
}

func NewRoleService(rights repos.RightsRepo, roles repos.RoleRepo) RoleService {
	return roleService{rights, roles}
}

func (r roleService) GetRoles() (*[]models.Role, error) {
	return r.roleRepo.GetAllRoles()
}

func (r roleService) GetRole(id int) (*models.Role, error) {
	return r.roleRepo.GetRoleByID(id)
}

func (r roleService) CreateNewRole(requestedBy int, role *models.Role) error {
	if r.rightRepo.RoleHasRight(requestedBy, "ROLE/WRITE") {
		if err := r.roleRepo.CreateRole(role); err != nil {
			return err
		}
		return nil
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (r roleService) GivePermissionToRole(requestedBy, roleID, rightID int) error {
	if r.rightRepo.RoleHasRight(requestedBy, "ROLE/GIVE-PERM") {
		if err := r.rightRepo.GiveRightToRole(roleID, rightID); err != nil {
			return err
		}
		return nil
	}
	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (r roleService) RevokeRightFromRole(requestedBy, roleID, rightID int) error {
	if r.rightRepo.RoleHasRight(requestedBy, "ROLE/REVOKE-PERM") {
		if err := r.rightRepo.RemoveRightFromRole(roleID, rightID); err != nil {
			return err
		}
		return nil
	}
	return errors.New("you don't have the appropriate permissions to perform this action")
}
