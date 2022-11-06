package repos

import (
	"fmt"
	"time"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/abelgalef/course-reg/pkg/storage"
)

type CourseRepo interface {
	GetCourse(int) (*models.Course, error)
	GetAllCourses() (*[]models.Course, error)
	GetDeptCourses(int) (*[]models.Course, error)
	CreateCourse(*models.Course) error
	UpdateCourse(*models.Course) error
	TakeCourse(int, int, time.Time) error
	GetCurrentTakenCreditsAndNumOfCourses(int, time.Time) (int, int, error)
	GetAllMyCourseHistory(int) (*[]models.Course, error)
	GetMyCurrentCourseHistory(int, time.Time) (*[]models.Course, error)
}

type courseRepo struct {
	db *storage.MainDB
}

func NewCourseRepo(db *storage.MainDB) CourseRepo {
	return courseRepo{db}
}

func (c courseRepo) GetCourse(courseID int) (*models.Course, error) {
	var course models.Course
	if err := c.db.Connection.First(&course, courseID).Error; err != nil {
		return nil, err
	}

	return &course, nil
}
func (c courseRepo) GetAllCourses() (*[]models.Course, error) {
	var courses []models.Course
	if err := c.db.Connection.Find(&courses).Error; err != nil {
		return nil, err
	}

	return &courses, nil
}

func (c courseRepo) GetDeptCourses(deptID int) (*[]models.Course, error) {
	var courses []models.Course
	if err := c.db.Connection.Where("dept_id = ?", deptID).Find(&courses).Error; err != nil {
		return nil, err
	}

	return &courses, nil
}

func (c courseRepo) CreateCourse(course *models.Course) error {
	return c.db.Connection.Model(&models.Department{ID: course.DeptID}).Association("Courses").Append(course)
}

func (c courseRepo) UpdateCourse(course *models.Course) error {
	return c.db.Connection.Model(&models.Course{ID: course.ID}).Updates(course).Error
}

func (c courseRepo) TakeCourse(courseID, userID int, timeAfter time.Time) error {
	return c.db.Connection.Where("course_id = ? AND user_id = ? AND created_at > ?", courseID, userID, timeAfter).FirstOrCreate(&models.CourseUser{CourseID: courseID, UserID: userID}).Error
}

func (c courseRepo) GetCurrentTakenCreditsAndNumOfCourses(userID int, timeAfter time.Time) (int, int, error) {
	type result struct{ count, totCredits int }
	var r result

	if err := c.db.Connection.Raw("SELECT count(*) count, sum(credits) num FROM courses WHERE id in (select course_id from course_users where user_id = ? and created_at > ?)", userID, timeAfter).Scan(&r).Error; err != nil {
		return 0, 0, err
	}

	return r.count, r.totCredits, nil
}

func (c courseRepo) GetAllMyCourseHistory(userID int) (*[]models.Course, error) {
	var courses []models.Course

	if err := c.db.Connection.Raw(fmt.Sprintf("SELECT courses.id, courses.course_id, courses.credits, courses.name, courses.active, courses.dept_id, course_users.created_at FROM courses INNER JOIN course_users ON courses.id = course_users.course_id WHERE user_id = %d ORDER BY course_users.created_at DESC", userID)).Scan(&courses).Error; err != nil {
		return nil, err
	}

	return &courses, nil
}

func (c courseRepo) GetMyCurrentCourseHistory(userID int, timeAfter time.Time) (*[]models.Course, error) {
	var courses []models.Course

	if err := c.db.Connection.Raw(fmt.Sprintf("SELECT courses.id, courses.course_id, courses.credits, courses.name, courses.active, courses.dept_id, course_users.created_at FROM courses INNER JOIN course_users ON courses.id = course_users.course_id WHERE user_id = %d AND course__users.created_at > %v ORDER BY course_users.created_at DESC", userID, timeAfter)).Scan(&courses).Error; err != nil {
		return nil, err
	}

	return &courses, nil
}
