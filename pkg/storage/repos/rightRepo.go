package repos

import (
	"fmt"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
)

type RightsRepo interface {
	GetRights() (*[]models.Right, error)
	RoleHasRight(int, string) bool
	RoleHasDeptRight(int, int, string) bool
	GiveRightToRole(int, int) error
	RemoveRightFromRole(int, int) error
}

type rightsRepo struct {
	db *storage.MainDB
}

func NewRightRepository(db *storage.MainDB) RightsRepo {
	return &rightsRepo{db}
}

func (r rightsRepo) GetRights() (*[]models.Right, error) {
	var rights []models.Right
	if err := r.db.Connection.Find(&rights).Error; err != nil {
		return nil, err
	}

	return &rights, nil
}

func (r rightsRepo) RoleHasRight(roleID int, rightIdentifier string) bool {
	var roleHasRight bool

	rawSQL := "SELECT EXISTS(SELECT 1 FROM role_permission WHERE role_id = %d AND right_id = (SELECT id FROM rights WHERE tag = '%s' LIMIT 1)) AS roleHasRight"

	r.db.Connection.Raw(fmt.Sprintf(rawSQL, roleID, rightIdentifier)).Scan(&roleHasRight)

	return roleHasRight
}

func (r rightsRepo) RoleHasDeptRight(roleID, deptID int, rightIdentifier string) bool {
	var roleHasRight bool

	rawSQL := "SELECT EXISTS(SELECT 1 FROM role_permission WHERE role_id = %d AND right_id = (SELECT id FROM rights WHERE tag = (SELECT CONCAT(dept_code, '%s') FROM departments WHERE id = %d)LIMIT 1) LIMIT 1) AS roleHasRight LIMIT 1"

	r.db.Connection.Raw(fmt.Sprintf(rawSQL, roleID, rightIdentifier, deptID)).Scan(&roleHasRight)

	return roleHasRight
}

func (r rightsRepo) GiveRightToRole(roleID int, rightID int) error {
	return r.db.Connection.Model(&models.Role{ID: roleID}).Omit("Rights.*").Association("Rights").Append(&models.Right{ID: uint(rightID)})
}

func (r rightsRepo) RemoveRightFromRole(roleID int, rightID int) error {
	return r.db.Connection.Model(&models.Role{ID: roleID}).Association("Rights").Delete(&models.Right{ID: uint(rightID)})
}
