package client_test

import (
	"context"
	"fmt"
	"testing"

	cc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/courses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/0B1t322/distanceLearningWebSite/protos/authservice"
	"github.com/0B1t322/distanceLearningWebSite/protos/coursesservice"

	grpc_middleware "github.com/0B1t322/distanceLearningWebSite/pkg/middleware"

	auth_client "github.com/0B1t322/distanceLearningWebSite/service.auth/client"
	"github.com/0B1t322/service.courses/client"
	"google.golang.org/grpc"
)

// func init() {
// 	c, err := auth_client.NewClient("127.0.0.1","5050", []grpc.DialOption{grpc.WithInsecure()})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer c.Close()

// 	ctx :=  metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", "admin"))

// 	resp, err := c.SignUp(ctx, &authservice.AuthRequest{Username: "dan", Password: "123"})
// 	if err != nil {
// 		panic(err)
// 	}

// 	opts = append(opts, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_middleware.TokenUnaryClientInterceptor(resp.Token)))
// }

func init() {
	opts = append(opts, grpc.WithInsecure())
}

func getToken() (string, error) {
	c, err := auth_client.NewClient("127.0.0.1","5050", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return "", err
	}
	defer c.Close()

	resp, err := c.SignIn(context.Background(), &authservice.AuthRequest{Username: "dan", Password:"123"})
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func preapareOpts() error {
	token, err := getToken()
	if err != nil {
		return err
	}

	opts = append(
		opts, 
		grpc.WithUnaryInterceptor(grpc_middleware.TokenUnaryClientInterceptor(token)),
	)

	return nil
}

var opts []grpc.DialOption


func TestFunc_NewClient(t *testing.T) {
	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()
	
}

func TestFunc_AddCourse(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()
	
	resp, err := c.AddCourse(
		context.Background(), 
		&coursesservice.AddCourseReq{
			Name: "course",
			ImgURL: "img",

		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("course add with id: %s\n", resp.Id)

	_, err = c.DeleteCourse(context.Background(), &coursesservice.DeleteCourseReq{Id: resp.Id})
	if  err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddCourse_AlreadyExsist(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()
	
	resp, err := c.AddCourse(
		context.Background(), 
		&coursesservice.AddCourseReq{
			Name: "course",
			ImgURL: "img",

		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Logf("course add with id: %s\n", resp.Id)

	_, err = c.AddCourse(
		context.Background(), 
		&coursesservice.AddCourseReq{
			Name: "course",
			ImgURL: "img",

		},
	)

	if status.Convert(err).Message() != cc.ErrCourseExsist.Error() {
		t.Log(err)
		t.FailNow()
	}

	_, err = c.DeleteCourse(context.Background(), &coursesservice.DeleteCourseReq{Id: resp.Id})
	if  err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_GetCourse(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	resp, err := c.AddCourse(
		context.Background(), 
		&coursesservice.AddCourseReq{
			Name: "course",
			ImgURL: "img",

		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		_, err = c.DeleteCourse(context.Background(), &coursesservice.DeleteCourseReq{Id: resp.Id})
		if  err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	t.Logf("course add with id: %s\n", resp.Id)

	course, err := c.GetCourse(
		context.Background(),
		&coursesservice.GetCourseReq{Id: resp.Id},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(course.String())
}

func TestFunc_GetCourse_NotFound(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	_, err = c.GetCourse(context.Background(), &coursesservice.GetCourseReq{Id: "213"})
	if status.Code(err) != codes.NotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_UpdateCourse(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	resp, err := c.AddCourse(
		context.Background(), 
		&coursesservice.AddCourseReq{
			Name: "course",
			ImgURL: "img",

		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		_, err = c.DeleteCourse(context.Background(), &coursesservice.DeleteCourseReq{Id: resp.Id})
		if  err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
	t.Logf("course add with id: %s\n", resp.Id)

	if _, err := c.UpdateCourse(
		context.Background(),
		&coursesservice.UpdateCourseReq{
			UpdatedCourse: &coursesservice.Course{
				Id: resp.Id,
				Name: "new_course",
			},
		},

	); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if course, err := c.GetCourse(
		context.Background(),
		&coursesservice.GetCourseReq{Id: resp.Id},	
	); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		t.Log(course.String())
	}
}

func TestFunc_GetAllCourses(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	for i := 0; i < 5; i++ {
		resp, err := c.AddCourse(
			context.Background(),
			&coursesservice.AddCourseReq{
				Name: fmt.Sprintf("course_%v", i),
				ImgURL: "img",
			},
		)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		defer func(id string) {
			_, err = c.DeleteCourse(context.Background(), &coursesservice.DeleteCourseReq{Id: id})
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
		}(resp.Id)
		if _, err := c.AddUserInCourse(
			context.Background(), 
			&coursesservice.AddUserInCourseReq{CourseID: resp.Id, UserID: "1"},
		); err != nil {
			t.Log(err)
			t.FailNow()
		}

		defer func(CID string) {
			_, err := c.DeleteUserInCourse(
				context.Background(), 
				&coursesservice.DeleteUserInCourseReq{UserID: "1", CourseID: CID})
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
		}(resp.Id)
		
	}

	cs, err := c.GetAllCourses(context.Background(), &coursesservice.GetAllCoursesReq{})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(cs.String())
}

func TestFunc_GetAllCourses_NotFound(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	if _, err := c.GetAllCourses(context.Background(), &coursesservice.GetAllCoursesReq{}); status.Code(err) != codes.NotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTaskHeader(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	var courseID string
	if course, err := c.AddCourse(
		context.Background(),
		&coursesservice.AddCourseReq{
			Name: "course",
			ImgURL: "img",
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		courseID = course.Id
	}
	defer func() {
		if _, err := c.DeleteCourse(
			context.Background(),
			&coursesservice.DeleteCourseReq{
				Id: courseID,
			},
		); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	var ID string
	if resp, err := c.AddTaskHeader(
		context.Background(),
		&coursesservice.AddTaskHeaderReq{
			CourseId: courseID,
			Name: "task_header_1",
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		t.Logf("Task_header with id: %v was added\n", resp.Id)
		ID = resp.Id	
	}

	if _, err := c.DeleteTaskHeader(
		context.Background(),
		&coursesservice.DeleteTaskHeaderReq{
			Id: ID,
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTaskHeader_AlreadyExsist(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	var taskHeaderID string
	if resp, err := c.AddTaskHeader(
		context.Background(),
		&coursesservice.AddTaskHeaderReq{
			Name: "task_header_1",
			CourseId: "2",
		},
	);err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		taskHeaderID = resp.Id
		t.Logf("Task_header with id: %v was added\n", resp.Id)
	}
	defer func() {
		if _, err := c.DeleteTaskHeader(
			context.Background(),
			&coursesservice.DeleteTaskHeaderReq{
				Id: taskHeaderID,
			},
		); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if _, err := c.AddTaskHeader(
		context.Background(),
		&coursesservice.AddTaskHeaderReq{
			Name: "task_header_1",
			CourseId: "2",
		},
	); status.Code(err) != codes.AlreadyExists {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_UpdateTaskHeader(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	var taskHeaderID string
	if resp, err := c.AddTaskHeader(
		context.Background(),
		&coursesservice.AddTaskHeaderReq{
			CourseId: "0",
			Name: "task_header_1",
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		taskHeaderID = resp.Id
	}
	defer func() {
		if _,  err := c.DeleteTaskHeader(
			context.Background(),
			&coursesservice.DeleteTaskHeaderReq{
				Id: taskHeaderID,
			},
		); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if _, err := c.UpdateTaskHeader(
		context.Background(),
		&coursesservice.UpdateTaskHeaderReq{
			TaskHeader: &coursesservice.TaskHeader{
				Id: taskHeaderID,
				Name: "updated_task_header_1",
			},
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddTask(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	var taskID string
	if resp, err := c.AddTask(
		context.Background(),
		&coursesservice.AddTaskReq{
			TaskHeaderId: "2",
			Name: "task_1",
			ImgUrl: "img",
			ContentURL: "content",
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		t.Logf("Task with id: %v added", resp.Id)
		taskID = resp.Id
	}
	defer func() {
		if _, err := c.DeleteTask(
			context.Background(),
			&coursesservice.DeleteTaskReq{
				Id: taskID,
			},
		); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
}

func TestFunc_AddTask_AlreadyExsist(t *testing.T) {
	err := preapareOpts()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	c, err := client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	var taskID string
	if resp, err := c.AddTask(
		context.Background(),
		&coursesservice.AddTaskReq{
			TaskHeaderId: "2",
			Name: "task_1",
			ImgUrl: "img",
			ContentURL: "content",
		},
	); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		t.Logf("Task with id: %v added", resp.Id)
		taskID = resp.Id
	}
	defer func() {
		if _, err := c.DeleteTask(
			context.Background(),
			&coursesservice.DeleteTaskReq{
				Id: taskID,
			},
		); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if _, err := c.AddTask(
		context.Background(),
		&coursesservice.AddTaskReq{
			TaskHeaderId: "2",
			Name: "task_1",
			ImgUrl: "img",
			ContentURL: "content",
		},
	); status.Code(err) != codes.AlreadyExists {
		t.Log(err)
		t.FailNow()
	}
}