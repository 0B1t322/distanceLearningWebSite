package courses

import "errors"

var (
	ErrCourseExsist 		= errors.New("Course exsist")
	ErrCourseNotFound		= errors.New("Corese not found")
	ErrTaskHeaderExsist		= errors.New("Task Header in  this course  with  this  name is  exist")
	ErrTaskHeaderNotFound	=  errors.New("Task Header not found")
)