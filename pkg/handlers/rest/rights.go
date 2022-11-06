package rest

import (
	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/services"
)

type RightHanler interface {
	GetRights() (*[]models.Right, error)
}

type rightHandler struct {
	right services.RightService
}

func NewRightHandler(right services.RightService) RightHanler {
	return rightHandler{right}
}

func (r rightHandler) GetRights() (*[]models.Right, error) {
	return r.right.GetRights()
}
