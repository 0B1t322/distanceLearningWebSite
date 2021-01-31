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

	pb "github.com/0B1t322/service.courses/coursesservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCoursesServiceServer

	courseController	*cc.CoursesController
}

func NewServer(
	courseDB *gorm.DB,
) *Server {
	return &Server {
		courseController: cc.New(courseDB),
	}
}

// Courses

func (s *Server) AddCourse(
	ctx context.Context, 
	req *pb.AddCourseReq,
)	(*pb.AddCourseResp, error) {

	if err := getRoleAndCheckThem(ctx); err != nil {
		return nil, err
	}
	
	model := &cm.Course{Name:  req.Name, ImgURL: req.ImgURL}

	if err := s.courseController.AddCourse(model); err == cc.ErrCourseExsist {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(req.TaskHeaders) == 0 {
		return nil, status.Error(codes.OK, "All is okey")
	}

	// TODO AddTaskHeader
	return &pb.AddCourseResp{}, status.Error(codes.OK, "OK")
}

func (s *Server) DeleteCourse(
	ctx context.Context, 
	req *pb.DeleteCourseReq,
)	(*pb.DeleteCourseResp, error) {

	if err := getRoleAndCheckThem(ctx); err != nil {
		return nil, err
	}

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
	if err != nil {
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