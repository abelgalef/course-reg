package services

import (
	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
)

type RightService interface {
	GetRights() (*[]models.Right, error)
}

type rightService struct {
	rights repos.RightsRepo
}

func NewRightService(rights repos.RightsRepo) RightService {
	return &rightService{rights}
}

func (r rightService) GetRights() (*[]models.Right, error) {
	return r.rights.GetRights()
}
