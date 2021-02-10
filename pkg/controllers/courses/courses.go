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
/* 
AddCourse Add Course if course with this name  does'nt exist
	errors:
		ErrCourseExsist
		some gorm internal errors
*/
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

/*
GetCourseById return a course with this id
	errors:
		ErrCourseNotFound
		some gorm errors
*/
func (c *CoursesController) GetCourseById(id string) (*model.Course, error) {
	course := &model.Course{}

	err := c.db.First(course, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrCourseNotFound
	} else if err != nil {
		return nil, err
	}

	return course, nil
}

/*
GetAllCourseForUser  return a slice of pointers  of course model by a user id
	errors:
		CourseNotFound
		some gorm  errors
*/
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


/*
GetCourseByName  return  a  course by his name
	errors:
		ErrCourseNotFound
		some  gorm  errors
*/
func (c *CoursesController) GetCourseByName(name string) (*model.Course,  error) {
	course := &model.Course{}

	err := c.db.First(course, "name = ?", name).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrCourseNotFound
	} else if err != nil {
		return nil, err
	}

	return course, nil
}


/*
DeleteCourse delete course by a model
	errors:
		ErrCourseNotFound
		some  gorm  errors
*/
func (c *CoursesController) DeleteCourse(course *model.Course) (error) {
	err := c.db.Delete(course, "id = ? or name = ?", course.ID, course.Name).Error
	if err == gorm.ErrRecordNotFound {
		return ErrCourseNotFound
	} else if err != nil {
		return err
	}

	ths, err := c.GetAllTaskHeadearsByCourseID(fmt.Sprint( course.ID) )
	if err == ErrTaskHeaderNotFound {
		return nil
	} else if err != nil {
		return err
	}

	for _, th := range ths {
		if err := c.DeleteTaskHeader(th); err != nil {
			return  err
		}
	}

	return nil
}


/*
UpdateCourse update exsisted course and taskHeader and tasks in them.
if you want to update only  courses send nil to others args
	errors:
		// TODO check all errors
*/
func (c *CoursesController) UpdateCourse(
	course 		*model.Course,
	taskHeaders	[]*model.TaskHeader,
	tasks		[]*model.Task,
) error {
	m := model.TasksToMap(tasks)

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

/*
AddTaskHeader add task header; if in this course have the task header with same  name will return error
	errors:
		ErrTaskHeaderExsist
		some gorm errors
*/
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


/*
GetTaskHeaderByID retuen a  task header  by id
	errors:
		ErrTaskHeaderNotFound
		some gorm errors
*/
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

/*
GetAllTaskHeadearsByCourseID return  a  slice  of  task  headers  by course id
	errors:
		ErrTaskHeaderNotFound
		some  gorm errors
*/
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

/*
DeleteTaskHeader delete  task header  by a model;
	errors:
		ErrTaskHeaderNotFound
*/
func (c *CoursesController) DeleteTaskHeader(th *model.TaskHeader) error {
	err := c.db.Delete(th).Error
	// If record not found return nil
	// idk what i can do for this
	if err == gorm.ErrRecordNotFound {
		return ErrTaskHeaderNotFound
	} else if err != nil {
		return err
	}

	ts, err := c.GetAllTasksByTaskHeaderID(fmt.Sprint( th.ID))
	if err == ErrTaskNotFound {
		return nil
	} else if err != nil {
		return err
	}

	for _, t := range ts {
		if err := c.DeleteTask(t); err != nil {
			return err
		}
	}

	return nil
}

/*
UpdateTaskHeader update task header
	errors:
		some gorm errors
		// TODO check other errors
*/
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

// Tasks


/*
AddTask - add task; in one taskheader can't exist two tasks with same names
	errors:
		ErrTaskExist
		some  gorm errors
*/
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

/*
GetTaskByID - get task by id
	errors:
		ErrTaskNotFound
		some  gorm errors
*/
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


/*
GetAllTasksByTaskHeaderID return a slice of tasks by  taskheader id
	errors:
		ErrTaskNotFound
		some gorm errors
*/
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


/*
DeleteTask delete task
	errors:
		TaskNotFound
		some gorm errors
*/
func (c *CoursesController) DeleteTask(t *model.Task) (error) {
	err := c.db.Delete(t).Error
	// If record not found return nil
	// idk what i can do for this
	if err == gorm.ErrRecordNotFound {
		return ErrTaskHeaderNotFound
	}

	return err
}
/*
UpdateTask update task that exist
	errors:
	some  gorm errors
*/
func (c *CoursesController) UpdateTask(t *model.Task) (error) {
	return c.db.Model(t).Updates(t).Error
}


// UsersInCourses

/*
AddUserInCourse add user to course by id
	errors:
		ErrUserAlreadyInCourse
		some  gorm errors
*/
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

/*
DeleteUserInCourse delete  user  from course by user id and course id

*/
func (c* CoursesController) DeleteUserInCourse(UID, CID string) (error) {
	return c.db.Table("users_in_courses").
				Where("user_id = ? AND course_id = ?", UID, CID).
				Delete(
						&model.UsersInCourse{},
				).Error
}


// TODO DeleteUserInCourse by model
// TODO  user not found