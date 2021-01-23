package courses

import (
	model "github.com/0B1t322/distanceLearningWebSite/pkg/models/courses"
	"gorm.io/gorm"
)

type CoursesController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CoursesController {
	return &CoursesController{
		db: db,
	}
}

func (c *CoursesController) AddCourse(course *model.Course) error {
	// check  if course  with this  name exsist
	var m model.Course
	err := c.db.First(&m, "name = ?", course.Name).Error
	if err != gorm.ErrRecordNotFound {
		return ErrCourseExsist
	}

	return c.db.Create(course).Error
}

func (c *CoursesController) GetCourseById(id string) (*model.Course, error) {
	course := &model.Course{}

	err := c.db.Find(course, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrCourseNotFound
	} else if err != nil {
		return nil, err
	}

	return course, nil
}

func (c *CoursesController) GetCourseByName(name string) (*model.Course,  error) {
	course := &model.Course{}

	err := c.db.Find(course, "name = ?", name).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrCourseNotFound
	} else if err != nil {
		return nil, err
	}

	return course, nil
}

func (c *CoursesController) DeleteCourse(course *model.Course) (error) {
	err := c.db.Delete(course, "id = ? or name = ?", course.ID, course.Name).Error
	if err == gorm.ErrRecordNotFound {
		return ErrCourseNotFound
	}

	return err
}

func (c *CoursesController) UpdateCourse(
	course 		*model.Course,
	taskHeaders	[]*model.TaskHeader,
	tasks		[]*model.Task,
) error {

	if err := c.db.Model(course).Updates(course).Error; err != nil {
		return err
	}
	
	if len(taskHeaders) == 0 {
		return nil
	}

	// TODO TaskHeaders  update

	return nil
}

// TaskHeaders

func (c *CoursesController) AddTaskHeader(th *model.TaskHeader) error {
	var taskHeader model.TaskHeader
	
	err := c.db.Where("course_id = ? and name = ?", th.CourseID, th.Name).First(&taskHeader).Error
	if err != gorm.ErrRecordNotFound {
		return ErrTaskHeaderExsist
	}

	return c.db.Create(th).Error
}

func (c *CoursesController) GetTaskHeaderByID(id string) (*model.TaskHeader, error) {
	th := &model.TaskHeader{}

	err := c.db.First(th,"id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrTaskHeaderNotFound 
	}  else if err != nil {
		return nil,  err
	}

	return  th, err
}

func (c *CoursesController) DeleteTaskHeader(th *model.TaskHeader) error {
	err := c.db.Delete(th).Error
	if err == gorm.ErrRecordNotFound {
		return ErrTaskHeaderNotFound
	}

	return err
}

func (c *CoursesController) UpdateTaskHeader(
	th *model.TaskHeader, 
	t []*model.Task,
) (error) {

}