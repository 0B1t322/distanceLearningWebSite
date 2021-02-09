package server

import (
	"github.com/0B1t322/service.courses/pkg/convertor"
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	cc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/courses"
	cm "github.com/0B1t322/distanceLearningWebSite/pkg/models/courses"
	"gorm.io/gorm"

	pb "github.com/0B1t322/distanceLearningWebSite/protos/coursesservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)
// Server a struct of server
// 
// contains controller
type Server struct {
	pb.UnimplementedCoursesServiceServer

	// controller that work with db
	courseController	*cc.CoursesController
}

// NewServer return a new server
func NewServer(
	courseDB *gorm.DB,
) *Server {
	return &Server {
		courseController: cc.New(courseDB),
	}
}

// Courses -----------------------------

/*
AddCourse add course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) AddCourse(
	ctx context.Context, 
	req *pb.AddCourseReq,
)	(*pb.AddCourseResp, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "Can't get metadata")
	}

	log.Infof("%v", md)

	// if err := getRoleAndCheckThem(ctx); err != nil {
	// 	return nil, err
	// }
	
	model := &cm.Course{Name:  req.Name, ImgURL: req.ImgURL}

	if err := s.courseController.AddCourse(model); err == cc.ErrCourseExsist {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.AddCourseResp{Id: fmt.Sprint(model.ID)}

	if len(req.TaskHeaders) == 0 {
		return resp, status.Error(codes.OK, "All is okey")
	}

	courseID := fmt.Sprint(model.ID)
	for _, th := range req.TaskHeaders {
		_, err := s.AddTaskHeader(
			ctx,
			&pb.AddTaskHeaderReq{
				CourseId: courseID,
				Name: th.Name,
				Tasks: th.Tasks,
			},
		)

		if err != nil {
			return nil, err
		}
	}

	return resp, status.Error(codes.OK, "OK")
}

/*
DeleteCourse delete course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) DeleteCourse(
	ctx context.Context, 
	req *pb.DeleteCourseReq,
)	(*pb.DeleteCourseResp, error) {
	ID, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	if err := s.courseController.DeleteCourse(&cm.Course{ID: ID}); err == cc.ErrCourseNotFound {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteCourseResp{}, status.Error(codes.OK, "Delete course successfully")
}


/*
GetCourse return course
*/
func (s *Server) GetCourse(
	ctx context.Context, 
	req *pb.GetCourseReq,
)	(*pb.GetCourseResp, error) {

	// UID, err := getUIDFromCtx(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	
	course, err := s.courseController.GetCourseById(req.Id)
	if err == cc.ErrCourseNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	courseID := fmt.Sprint(course.ID)

	ths, err := s.getTaskHeaderForCourse(courseID)
	if err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	return &pb.GetCourseResp{
		Course: &pb.Course{
			Id: 			courseID,
			Name: 			course.Name,
			ImgURL: 		course.ImgURL,	
			TaskHeaders:	ths,
		},
	}, status.Error(codes.OK, "Get course!")

	// Check if user in this course or it's teacher or admin
}

/*
UpdateCourse update course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) UpdateCourse(
	ctx context.Context, 
	req *pb.UpdateCourseReq,
)	(*pb.UpdateCourseResp, error) {
	c, ths, ts, err := convertor.PBCourseToModels(req.UpdatedCourse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.courseController.UpdateCourse(c,ths,ts); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateCourseResp{}, status.Error(codes.OK, "Updated")
}
/*
GetAllCourses return all course for user 

UID in metadata
*/
func (s *Server) GetAllCourses(
	ctx context.Context, 
	req *pb.GetAllCoursesReq,
)	(*pb.GetAllCoursesResp, error) {

	UID, err := getUIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	courses, err := s.courseController.GetAllCourseForUser(UID)
	if err == cc.ErrCourseNotFound {
		return nil, status.Error(codes.NotFound, "Not found courses for this user")
	}

	// TODO think if need to load some more
	var cs []*pb.Course
	for _, c := range courses {
		cs = append(
			cs,
			&pb.Course{
				Id: fmt.Sprint(c.ID),
				Name: c.Name,
				ImgURL: c.ImgURL,
				TaskHeaders: nil,
			},
		)
	}

	return &pb.GetAllCoursesResp{Courses: cs}, status.Error(codes.OK, "Ok")
}

// -----------------------------

// TaskHeader -----------------------------

/*
AddTaskHeader add taskHeadewr for course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) AddTaskHeader(
	ctx context.Context,
	req *pb.AddTaskHeaderReq, 
)	(*pb.AddTaskHeaderResp, error) {
	// TODO check for userid and  permission

	model := &cm.TaskHeader{CourseID: req.CourseId, Name: req.Name}
	if err :=  s.courseController.AddTaskHeader(model); err == cc.ErrTaskHeaderExsist {
		return nil, status.Errorf(codes.AlreadyExists, "%v", err)
	} else if err != nil {
		return nil,  status.Errorf(codes.Internal, "%v", err)
	}

	taskHeaderID := fmt.Sprint(model.ID)

	for _, t := range req.Tasks {
		_, err := s.AddTask(
			ctx, 
			&pb.AddTaskReq{
				TaskHeaderId: taskHeaderID, 
				Name: t.Name,
				ImgUrl: t.ImgURL,
				ContentURL: t.ContentURL,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	return &pb.AddTaskHeaderResp{Id: fmt.Sprint(model.ID)}, status.Error(codes.OK, "Okay")
}

/*
UpdateTaskHeader update taskHeadewr for course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) UpdateTaskHeader(
	ctx context.Context,
	req *pb.UpdateTaskHeaderReq,
)	(*pb.UpdateTaskHeaderResp, error) {
	
	th, ts, err := convertor.PBTaskHeaderToModel(req.TaskHeader)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.courseController.UpdateTaskHeader(th, ts); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateTaskHeaderResp{}, nil
}

/*
DeleteTaskHeader delete taskheader from course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) DeleteTaskHeader(
	ctx context.Context,
	req *pb.DeleteTaskHeaderReq,
)	(*pb.DeleteTaskHeaderResp, error) {
	ID, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	if err := s.courseController.DeleteTaskHeader(&cm.TaskHeader{ID: ID}); err == cc.ErrTaskHeaderNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteTaskHeaderResp{}, status.Error(codes.OK, "Deleted")
}
// -----------------------------

// Tasks -----------------------------

/*
AddTask add taask in task header

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) AddTask(
	ctx context.Context,
	req *pb.AddTaskReq,
)	(*pb.AddTaskResp, error) {
	model := &cm.Task {
		Name: req.Name,
		ImgURL: req.ImgUrl,
		TaskHeaderID: req.TaskHeaderId,
	}

	if err := s.courseController.AddTask(model); err == cc.ErrTaskExist {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AddTaskResp{Id: fmt.Sprint(model.ID)}, status.Error(codes.OK, "Task Added")
}

/*
UpdateTask update task header in course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) UpdateTask(
	ctx context.Context,
	req *pb.UpdateTaskReq,
)	(*pb.UpdateTaskResp, error) {
	ID, err := strconv.ParseInt(req.Task.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	model := &cm.Task{
		ID: ID,
		Name: req.Task.Name,
		ImgURL: req.Task.ImgURL, 
	}

	if err := s.courseController.UpdateTask(model); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateTaskResp{},  status.Error(codes.OK, "Task Updated")
}

/*
DeleteTask delete task from task header

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) DeleteTask(
	ctx context.Context,
	req *pb.DeleteTaskReq,
)	(*pb.DeleteTaskResp, error) {
	ID, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	if err := s.courseController.DeleteTask(&cm.Task{ID:ID}); err == cc.ErrTaskNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteTaskResp{}, status.Error(codes.OK, "Task deleted")
}

// -----------------------------

// UserInCourse -----------------------------

/*
AddUserInCourse add user in course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func  (s *Server) AddUserInCourse(
	ctx context.Context,
	req *pb.AddUserInCourseReq,
)	(*pb.AddUserInCourseResp, error) {
	if err := s.courseController.AddUserInCourse(
		&cm.UsersInCourse{
			UserID: req.UserID, 
			CourseID: req.CourseID,
		},
	); err == cc.ErrUserAlredyInCourse {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AddUserInCourseResp{}, status.Error(codes.OK, "User  added in course")
}

/*
DeleteUserInCourse dleete user in course

To call this procedure the role of the user should be - ["teacher", "admin", "moderator"]
*/
func (s *Server) DeleteUserInCourse(
	ctx context.Context,
	req *pb.DeleteUserInCourseReq,
)	(*pb.DeleteUserInCourseResp, error ) {
	if err := s.courseController.DeleteUserInCourse(req.UserID, req.CourseID); err != nil {
		return nil ,status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteUserInCourseResp{}, status.Error(codes.OK, "User deleted")
}

func getRoleFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Internal, "Can't get metadata")
	}

	var role string
	if r := md.Get("role"); len(r) > 0 {
		role = r[0]
	} else {
		return "", status.Error(codes.Internal, "Can't get role from ctx")	
	}

	return role, nil
}

func checkRole(role string) error {
	if role != "teacher" || role != "admin" {
		return status.Error(codes.PermissionDenied, "You can't do this")
	}

	return nil
}

func getRoleAndCheckThem(ctx context.Context) (error) {
	role, err := getRoleFromCtx(ctx)
	if err != nil {
		return err
	}

	err = checkRole(role)
	if err != nil {
		return err
	}

	return nil
}

func getUIDFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return "", status.Error(codes.Internal, "Can't get metadata")
	}

	var UID string
	if i := md.Get("uid"); len(i) > 0 {
		UID = i[0]
	} else {
		return "", status.Error(codes.Internal, "Can't get uid from ctx")	
	}

	return UID, nil
}

func (s *Server) getTaskHeaderForCourse(ID string) ([]*pb.TaskHeader, error) {
	var ths []*pb.TaskHeader

	taskHeaders, err := s.courseController.GetAllTaskHeadearsByCourseID(ID)
	if err == cc.ErrTaskHeaderNotFound {
		// TODO maybe course can be without task Headers
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil,  status.Error(codes.Internal, err.Error())
	}

	for _, th := range taskHeaders {
		thID := fmt.Sprint(th.ID)
		ts, err := s.getTasksForTaskHeader(thID)
		if err != nil {
			return nil, err
		} 

		ths = append(
			ths, 
			&pb.TaskHeader{
				Id: 	thID, 
				Name: 	th.Name,
				Tasks: 	ts,
			},
		)
	}

	return  ths, nil
}

func (s *Server) getTasksForTaskHeader(ID string)  ([]*pb.Task,  error) {
	var ts []*pb.Task
	
	tasks, err := s.courseController.GetAllTasksByTaskHeaderID(ID)
	if err == cc.ErrTaskNotFound {
		log.Warnf("some error on finding tasks: %v\n", err)
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, t := range tasks {
		ts = append(
			ts,
			&pb.Task{
				Id: 		fmt.Sprint(t.ID),
				Name: 		t.Name,
				ImgURL: 	t.ImgURL,
				ContentURL: "",	
			}, 
		)
	}

	return ts, nil
}