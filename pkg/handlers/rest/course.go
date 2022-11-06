package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CourseHandler interface {
	GetAllCourses(*gin.Context) (*[]models.Course, error)
	GetDeptCourses(*gin.Context) (*[]models.Course, error)
	CreateCourse(*gin.Context) (*models.Course, []utils.Error)
	UpdateCourse(*gin.Context) (*models.Course, []utils.Error)
	TakeCourse(*gin.Context) error
	GetAllMyCourseHistory(*gin.Context) (*[]models.Course, error)
	GetMyCurrentCourseHistory(*gin.Context) (*[]models.Course, error)
}

type courseHandler struct {
	course services.CourseService
}

func NewCourseHandler(course services.CourseService) CourseHandler {
	return courseHandler{course}
}

func (c courseHandler) GetAllCourses(_ *gin.Context) (*[]models.Course, error) {
	return c.course.GetAllCourses()
}

func (c courseHandler) GetDeptCourses(ctx *gin.Context) (*[]models.Course, error) {
	dID := ctx.Param("id")

	id, err := strconv.Atoi(dID)
	if err != nil {
		return nil, err
	}

	return c.course.GetDeptCourses(id)
}

func (c courseHandler) CreateCourse(ctx *gin.Context) (*models.Course, []utils.Error) {
	var course models.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	usr := ctx.MustGet("user").(*models.User)

	if err := c.course.CreateCourse(&course, usr.RoleID); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not create Course")
	}

	return &course, nil
}

func (c courseHandler) UpdateCourse(ctx *gin.Context) (*models.Course, []utils.Error) {
	var course models.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Binding Error")
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, utils.NewError(http.StatusBadRequest, err, "Course ID is not in the proper format")
	}

	course.ID = id
	usr := ctx.MustGet("user").(*models.User)
	err = c.course.UpdateCourse(&course, usr.RoleID)
	if err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, err, "Could not update course")
	}

	return &course, nil
}

func (c courseHandler) TakeCourse(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errors.New("course ID is not in the proper format")
	}

	usr := ctx.MustGet("user").(*models.User)
	return c.course.TakeCourse(id, usr.ID, usr.RoleID)
}

func (c courseHandler) GetAllMyCourseHistory(ctx *gin.Context) (*[]models.Course, error) {
	usr := ctx.MustGet("user").(*models.User)
	return c.course.GetAllMyCourseHistory(usr.ID)
}

func (c courseHandler) GetMyCurrentCourseHistory(ctx *gin.Context) (*[]models.Course, error) {
	usr := ctx.MustGet("user").(*models.User)
	return c.course.GetMyCurrentCourseHistory(usr.ID)
}
