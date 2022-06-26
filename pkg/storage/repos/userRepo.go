package repos

import (
	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
)

type UserRepo interface {
	GetUserWithID(uint) (*models.User, error)
	GetUserWithEmail(string) (*models.User, error)
	AddNewUser(*models.User) error
}

type userRepo struct {
	db *storage.MainDB
}

func NewUserRepo(db *storage.MainDB) UserRepo {
	return &userRepo{db}
}

func (ur *userRepo) GetUserWithID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.Connection.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) GetUserWithEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Connection.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) AddNewUser(user *models.User) error {
	if err := ur.db.Connection.Create(user).Error; err != nil {
		return err
	}

	return nil
}
