package courses_test

import (
	cm "github.com/0B1t322/distanceLearningWebSite/pkg/models/courses"
	cc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/courses"
	"time"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"testing"
)


var (
	DBManger = db.NewManager(
		"root",
		"root",
		"127.0.0.1:3306",
		15 * time.Second,
	)
)

func TestFunc_AddCourse(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err  != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	err = controller.AddCourse(&cm.Course{
		Name: "course_1",
		ImgURL: "dasdawasfasf",
	})

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_GetByID(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err  != nil {
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

func TestFunc_GetByName(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err  != nil {
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

func TestFunc_Delete_ByName(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err  != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	err = controller.AddCourse(&cm.Course{Name: "new_course"})
	if err !=  nil {
		t.Log(err)
		t.FailNow()
	}

	err = controller.DeleteCourse(&cm.Course{Name: "new_course"})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_Delete_ById(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err  != nil {
		t.Log(err)
		t.FailNow()
	}
	controller := cc.New(db)

	err = controller.AddCourse(&cm.Course{Name: "new_course"})
	if err !=  nil {
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
	if err  != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_GetAllTaskHeaders(t *testing.T) {
	db, err := DBManger.OpenDataBase("courses")
	if err  != nil {
		t.Log(err)
		t.FailNow()
	}
	// controller := cc.New(db)
	taskHeaders :=  []*cm.TaskHeader{}
	if  err := db.Where("course_id = ?", 1).Order("name").Find(&taskHeaders).Error;  err !=  nil {
		t.Log(err)
		t.FailNow()
	}

	for  _, th := range taskHeaders {
		t.Log(th)
	}
}