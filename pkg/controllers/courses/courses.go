package courses

import (
	"fmt"

	model "github.com/0B1t322/distanceLearningWebSite/pkg/models/courses"
	"gorm.io/gorm"
)

type CoursesController struct {
	db 		*gorm.DB
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
	if err == nil {
		return ErrCourseExsist
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return err
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

func (c *CoursesController) GetAllCourseForUser(UID string) ([]*model.Course, error) {
	var (
		courses 	[]*model.Course
		coursesID	[]int64
	)

	err := c.db.Table("users_in_courses").
				Select("course_id").
				Where("user_id = ?", UID).
				Find(&coursesID).Error

	if err != nil {
		return nil, err
	}

	if len(coursesID) == 0 {
		return nil, ErrCourseNotFound
	}

	err = c.db.Table("courses").
				Find(&courses, coursesID).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
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
	m := tasksToMap(tasks)

	if err := c.db.Model(course).Updates(course).Error; err != nil {
		return err
	}
	
	if len(taskHeaders) == 0 {
		return nil
	}

	for _, th := range taskHeaders {
		c.UpdateTaskHeader(th, m[fmt.Sprint(th.ID)])
	}

	return nil
}

// TaskHeaders

func (c *CoursesController) AddTaskHeader(th *model.TaskHeader) error {
	var taskHeader model.TaskHeader
	
	err := c.db.Where("course_id = ? and name = ?", th.CourseID, th.Name).First(&taskHeader).Error
	if err == nil {
		return ErrTaskHeaderExsist
	} else if err == nil && err != gorm.ErrRecordNotFound {
		return err
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

func (c *CoursesController) GetAllTaskHeadearsByCourseID(
	id string,
) ([]*model.TaskHeader, error) {
	var ths []*model.TaskHeader

	err := c.db.Table("task_headers").Where("course_id = ?", id).Find(&ths).Error
	if err != nil {
		return nil, err
	}

	if len(ths) == 0 {
		return nil, ErrTaskHeaderNotFound
	}

	return ths, nil
}

func (c *CoursesController) DeleteTaskHeader(th *model.TaskHeader) error {
	err := c.db.Delete(th).Error
	// If record not found return nil
	// idk what i can do for this
	if err == gorm.ErrRecordNotFound {
		return ErrTaskHeaderNotFound
	}

	return err
}

func (c *CoursesController) UpdateTaskHeader(
	th *model.TaskHeader, 
	tasks []*model.Task,
) (error) {
	if err := c.db.Model(th).Updates(th).Error; err != nil {
		return err
	}
	
	if len(tasks) == 0 {
		return nil
	}

	for _, t := range tasks {
		err := c.UpdateTask(t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CoursesController) AddTask(t *model.Task) (error) {
	var task model.Task
	
	err := c.db.Where("task_header_id = ? and name = ?", t.TaskHeaderID, t.Name).First(&task).Error
	if err == nil {
		return ErrTaskExist
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return c.db.Create(t).Error
}

func (c *CoursesController) GetTaskByID(id string) (*model.Task, error) {
	t := &model.Task{}

	err := c.db.First(t, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrTaskNotFound
	} else if err != nil {
		return nil, err
	}

	return t, nil
}

func (c* CoursesController) GetAllTasksByTaskHeaderID(
	id string,
) ([]*model.Task, error) {
	var tasks []*model.Task

	err := c.db.Table("tasks").Where("task_header_id = ?", id).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, ErrTaskNotFound
	}

	return tasks, nil
}

func (c *CoursesController) DeleteTask(t *model.Task) (error) {
	err := c.db.Delete(t).Error
	// If record not found return nil
	// idk what i can do for this
	if err == gorm.ErrRecordNotFound {
		return ErrTaskHeaderNotFound
	}

	return err
}

func (c *CoursesController) UpdateTask(t *model.Task) (error) {
	return c.db.Model(t).Updates(t).Error
}

func tasksToMap(tasks []*model.Task) map[string][]*model.Task {
	m := make(map[string][]*model.Task)

	for _, t := range tasks {
		if len(m[t.TaskHeaderID]) == 0 {
			m[t.TaskHeaderID] = []*model.Task{}
		}

		m[t.TaskHeaderID] = append(m[t.TaskHeaderID], t)
	}

	return m
}

func (c *CoursesController) AddUserInCourse(uc *model.UsersInCourse) (error) {
	var userInCourse model.UsersInCourse
	
	err := c.db.Table("users_in_courses").
				Where(
					"user_id = ? AND course_id = ?", 
					uc.UserID, uc.CourseID,
				).First(&userInCourse).Error
	if err == nil { // значит смог найти запись
		return ErrUserAlredyInCourse
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return c.db.Create(uc).Error
}

func (c* CoursesController) DeleteUserInCourse(UID, CID string) (error) {
	return c.db.Table("users_in_courses").
				Where("user_id = ? AND course_id = ?", UID, CID).
				Delete(
						&model.UsersInCourse{},
				).Error
}

// TODO DeleteUserInCourse by model