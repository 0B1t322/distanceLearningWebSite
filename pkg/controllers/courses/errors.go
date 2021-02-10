package courses

import "errors"

var (
	ErrCourseExsist 		= errors.New("Course exsist")
	ErrCourseNotFound		= errors.New("Course not found")
	ErrTaskHeaderExsist		= errors.New("Task Header in  this course  with  this  name is  exist")
	ErrTaskHeaderNotFound	= errors.New("Task Header not found")
	ErrTaskExist			= errors.New("Task exist")
	ErrTaskNotFound			= errors.New("Task not found")
	ErrUserAlredyInCourse	= errors.New("User already in this course")
	ErrUserNotFoundInCourse = errors.New("User not found in course")
)