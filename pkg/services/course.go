package services

import (
	"errors"
	"fmt"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
)

type CourseService interface {
	GetAllCourses() (*[]models.Course, error)
	GetDeptCourses(int) (*[]models.Course, error)
	CreateCourse(*models.Course, int) error
	UpdateCourse(*models.Course, int) error
	TakeCourse(int, int, int) error
	GetAllMyCourseHistory(int) (*[]models.Course, error)
	GetMyCurrentCourseHistory(int) (*[]models.Course, error)
}

type courseService struct {
	course    repos.CourseRepo
	rights    repos.RightsRepo
	constRepo repos.ConstRepo
}

func NewCourseService(course repos.CourseRepo, rights repos.RightsRepo, constRepo repos.ConstRepo) CourseService {
	return courseService{course, rights, constRepo}
}

func (c courseService) GetAllCourses() (*[]models.Course, error) {
	return c.course.GetAllCourses()
}

func (c courseService) GetDeptCourses(deptID int) (*[]models.Course, error) {
	return c.course.GetDeptCourses(deptID)
}

func (c courseService) CreateCourse(course *models.Course, requestedBy int) error {
	if c.rights.RoleHasDeptRight(requestedBy, course.DeptID, "/COURSE-WRITE") {
		return c.course.CreateCourse(course)
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (c courseService) UpdateCourse(course *models.Course, requestedBy int) error {
	if c.rights.RoleHasDeptRight(requestedBy, course.DeptID, "/COURSE-WRITE") {
		return c.course.UpdateCourse(course)
	}

	return errors.New("you don't have the appropriate permissions to perform this action")
}

func (c courseService) TakeCourse(courseID, userID, requestedBy int) error {
	course, err := c.course.GetCourse(courseID)
	if err != nil {
		return err
	}

	if !c.rights.RoleHasDeptRight(requestedBy, course.DeptID, "/ADD-COURSE") {
		return errors.New("you don't have the appropriate permissions to perform this action")
	}

	con, err := c.constRepo.GetLatestConstraintForDept(course.DeptID)
	if err != nil {
		return err
	}

	count, totCredits, err := c.course.GetCurrentTakenCreditsAndNumOfCourses(userID, con.CreatedAt)
	if err != nil {
		return err
	}

	if con.AllowedNumOfCourse <= count {
		return errors.New("you have reached the maximum number of courses you can take, " + fmt.Sprint(count))
	}
	if con.AllowedNumOfCredit <= totCredits {
		return errors.New("you have exceded the total amount of credits you can take " + fmt.Sprint(count) + " please choose another course with a lower credit")
	}

	return c.course.TakeCourse(course.ID, userID, con.CreatedAt)
}

func (c courseService) GetAllMyCourseHistory(userID int) (*[]models.Course, error) {
	return c.course.GetAllMyCourseHistory(userID)
}

func (c courseService) GetMyCurrentCourseHistory(userID int) (*[]models.Course, error) {
	con, err := c.constRepo.GetLatestConstraintForDept(1)
	if err != nil {
		return nil, err
	}

	return c.course.GetMyCurrentCourseHistory(userID, con.CreatedAt)
}
