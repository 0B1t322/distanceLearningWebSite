package courses_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/courses"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	cm "github.com/0B1t322/distanceLearningWebSite/pkg/models/courses"
)

var (
	DBManger = db.NewManager(
		"root",
		"root",
		"127.0.0.1:3306",
		15*time.Second,
	)
)

func TestFunc_AddCourse(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)
	c := &cm.Course{
		Name:   "course_1",
		ImgURL: "dasdawasfasf",
	}

	err = controller.AddCourse(c)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer controller.DeleteCourse(c)
}

func TestFunc_GetCourseByID(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	course, err := controller.GetCourseById("3")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(course)

}

func TestFunc_GetCourseByName(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	course, err := controller.GetCourseByName("lol")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(course)
}

func TestFunc_DeleteCourseByName(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	err = controller.AddCourse(&cm.Course{Name: "new_course"})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = controller.DeleteCourse(&cm.Course{Name: "new_course"})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_DeleteCourseById(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	err = controller.AddCourse(&cm.Course{Name: "new_course"})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := controller.GetCourseByName("new_course")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c.Name = ""

	err = controller.DeleteCourse(c)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTaskHeader_And_Delete(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)
	th := &cm.TaskHeader{CourseID: "2", Name: "teahcers  tasks"}
	err = controller.AddTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = controller.DeleteTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTaskHeader_ErrExist(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)
	th := &cm.TaskHeader{CourseID: "2", Name: "teahcers  tasks"}
	err = controller.AddTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	same_th := &cm.TaskHeader{CourseID: "2", Name: "teahcers  tasks"}
	err = controller.AddTaskHeader(same_th)
	if err != cc.ErrTaskHeaderExsist {
		t.Log(err)
		t.FailNow()
	}

	_th := &cm.TaskHeader{CourseID: "1", Name: "teahcers  tasks"}
	err = controller.AddTaskHeader(_th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = controller.DeleteTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = controller.DeleteTaskHeader(_th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_GetTaskHeaderByID(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)
	th := &cm.TaskHeader{CourseID: "2", Name: "teahcers  tasks"}
	err = controller.AddTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	getTh, err := controller.GetTaskHeaderByID(fmt.Sprint(th.ID))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	assert.EqualValues(t, th, getTh, "Should bee equal")

	err = controller.DeleteTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_DeleteTaskHeader_ErrTaskHeaderNotFound(t *testing.T) {
	// Skip while refactor func because if taskheader not exsist retirn nil
	t.SkipNow()
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	th := &cm.TaskHeader{ID: 10, CourseID: "10000", Name: "dasd"}

	err = controller.DeleteTaskHeader(th)
	if err != cc.ErrTaskHeaderNotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTask(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	task := &cm.Task{TaskHeaderID: "1", Name: "some_task"}

	err = controller.AddTask(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = controller.DeleteTask(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTask_ErrTaskExist(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	task := &cm.Task{TaskHeaderID: "1", Name: "some_task"}

	err = controller.AddTask(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteTask(task)

	sameTask := &cm.Task{TaskHeaderID: "1", Name: "some_task"}

	if err := controller.AddTask(sameTask); err != cc.ErrTaskExist {
		t.Log(err)
		t.FailNow()
	}

	moreTask := &cm.Task{TaskHeaderID: "2", Name: "some_task"}

	err = controller.AddTask(moreTask)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteTask(moreTask)
}

func TestFunc_GetTaskByID(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	task := &cm.Task{TaskHeaderID: "2", Name: "task"}

	err = controller.AddTask(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteTask(task)

	getTask, err := controller.GetTaskByID(fmt.Sprint(task.ID))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	assert.Equal(t, task, getTask, "should be equal")
}

func TestFunc_UpdateTask(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	task := &cm.Task{TaskHeaderID: "1", Name: "task"}
	err = controller.AddTask(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteTask(task)

	task.Name = "updated name"

	err = controller.UpdateTask(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_UpdateTaskHeader(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)
	th := &cm.TaskHeader{CourseID: "2", Name: "task_header_1"}
	err = controller.AddTaskHeader(th)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteTaskHeader(th)
	th.Name = "updated_task_header_1"

	ts := []*cm.Task{
		&cm.Task{Name: "task_1", TaskHeaderID: "2"},
		&cm.Task{Name: "task_2", TaskHeaderID: "2"},
		&cm.Task{Name: "task_3", TaskHeaderID: "2"},
		&cm.Task{Name: "task_4", TaskHeaderID: "2"},
	}

	for i, task := range ts {
		err := controller.AddTask(task)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		task.Name = fmt.Sprintf("updated_task_%v", i+1)
		defer func(task *cm.Task) {
			err := controller.DeleteTask(task)
			if err != nil {
				t.Log(err)
			}
		}(task)
	}

	err = controller.UpdateTaskHeader(th, ts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_UpdateCourse(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	c := &cm.Course{Name: "course_1"}

	if err := controller.AddCourse(c); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteCourse(c)

	c.Name = "updated_course_1"
	if err := controller.UpdateCourse(c, nil, nil); err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestFunc_GetAllTasksByTaskHeaderID(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	tasks := []*cm.Task{
		&cm.Task{TaskHeaderID: "2", Name: "tasks_1"},
		&cm.Task{TaskHeaderID: "2", Name: "tasks_2"},
		&cm.Task{TaskHeaderID: "2", Name: "tasks_3"},
		&cm.Task{TaskHeaderID: "1", Name: "tasks_4"},
		&cm.Task{TaskHeaderID: "1", Name: "tasks_5"},
		&cm.Task{TaskHeaderID: "3", Name: "tasks_6"},
	}

	for _, task := range tasks {
		err := controller.AddTask(task)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		defer controller.DeleteTask(task)
	}

	allTasks, err := controller.GetAllTasksByTaskHeaderID("2")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%v len: %v", allTasks, len(allTasks))

	for _, e := range allTasks {
		t.Logf("%v", e)
	}

}

func TestFunc_GetAllCoursesForUser(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	cFirst 	:= &cm.Course{Name: "course_1"}
	cSecond	:= &cm.Course{Name: "course_2"}

	if err := controller.AddCourse(cFirst); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteCourse(cFirst)

	if err := controller.AddCourse(cSecond); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteCourse(cSecond)

	ucF := &cm.UsersInCourse{UserID: "4", CourseID: fmt.Sprint(cFirst.ID) } 
	ucS := &cm.UsersInCourse{UserID: "4", CourseID: fmt.Sprint(cSecond.ID) } 

	if err := controller.AddUserInCourse(ucF); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteUserInCourse(ucF.UserID, ucF.CourseID)

	if err := controller.AddUserInCourse(ucS); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer controller.DeleteUserInCourse(ucS.UserID, ucS.CourseID)

	cs, err := controller.GetAllCourseForUser("4")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%v", cs)
}

func TestFunc_AddUserInCourse(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	controller := cc.New(db)

	uc := &cm.UsersInCourse{UserID: "1", CourseID: "2"}

	if err := controller.AddUserInCourse(uc); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if err := controller.DeleteUserInCourse(uc.UserID, uc.CourseID); err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestFunc_SliceToMap(t *testing.T) {
	slice := []*cm.Task{
		&cm.Task{TaskHeaderID: "2", Name: "tasks_1"},
		&cm.Task{TaskHeaderID: "2", Name: "tasks_2"},
		&cm.Task{TaskHeaderID: "2", Name: "tasks_3"},
		&cm.Task{TaskHeaderID: "1", Name: "tasks_4"},
		&cm.Task{TaskHeaderID: "1", Name: "tasks_5"},
		&cm.Task{TaskHeaderID: "3", Name: "tasks_6"},
	}

	m := make(map[string][]*cm.Task)

	for _, e := range slice {
		if len(m[e.TaskHeaderID]) == 0 {
			m[e.TaskHeaderID] = []*cm.Task{}
		}

		m[e.TaskHeaderID] = append(m[e.TaskHeaderID], e)
	}

	for k, v := range m {
		t.Logf("key: %s, val: %v", k, v)
	}
}

// TODO test for UserUnCourse
