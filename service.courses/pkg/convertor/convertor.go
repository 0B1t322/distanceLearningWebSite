package convertor

import (
	"fmt"
	"strconv"

	cm "github.com/0B1t322/distanceLearningWebSite/pkg/models/courses"
	pb "github.com/0B1t322/distanceLearningWebSite/protos/coursesservice"
)

// PBCourseToModels  convert pb models to curses model
func PBCourseToModels(course *pb.Course) (*cm.Course, []*cm.TaskHeader, []*cm.Task, error) {
	var (
		c 	*cm.Course
		ths	[]*cm.TaskHeader
		ts	[]*cm.Task
	)

	c, err := PBCourseToCourseModel(course)
	if err != nil {
		return nil, nil, nil, err
	}

	courseID := fmt.Sprint(c.ID)

	for _, th := range course.TaskHeaders {
		ID, err := strconv.ParseInt(th.Id, 10, 64)
		if err != nil {
			return nil, nil, nil, err
		}

		ths = append(
			ths,
			&cm.TaskHeader{
				ID: ID,
				Name: th.Name,
				CourseID: courseID,
			},
		)

		for _, t := range th.Tasks {
			ID, err := strconv.ParseInt(t.Id, 10, 64)
			if err != nil {
				return nil, nil, nil, err
			}

			ts = append(
				ts,
				&cm.Task{
					ID: ID,
					Name: t.Name,
					ImgURL: t.ImgURL,
					TaskHeaderID: th.Id,
				}, 
			)
		}
	}

	return c, ths, ts, nil
}

func PBCourseToCourseModel(course *pb.Course) (*cm.Course, error) {
	ID, err := strconv.ParseInt(course.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	
	c := &cm.Course{
		ID: ID,
		Name: course.Name,
		ImgURL: course.ImgURL,
	}

	return c, nil
}

func PBTaskHeaderToModel(taskHeader *pb.TaskHeader) (*cm.TaskHeader, []*cm.Task, error) {

	var (
		th 	*cm.TaskHeader
		ts	[]*cm.Task
	)

	ID, err := strconv.ParseInt(taskHeader.Id, 10, 64)
	if err != nil {
		return nil, nil, err
	}

	th = &cm.TaskHeader{
		ID: ID,
		Name: taskHeader.Name,
	}

	for _, t := range taskHeader.Tasks {
		ID, err := strconv.ParseInt(t.Id, 10, 64)
		if err != nil {
			return nil, nil, err
		}

		ts = append(
			ts,
			&cm.Task{
				ID: ID,
				Name: t.Name,
				ImgURL: t.ImgURL,
				TaskHeaderID: taskHeader.Id,
			}, 
		)
	}

	return th, ts, nil
}