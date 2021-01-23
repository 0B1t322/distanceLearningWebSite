package server

import (
	"context"

	pb "github.com/0B1t322/service.courses/coursesservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCoursesServiceServer
}

// Courses

func (s *Server) AddCourse(
	ctx context.Context, 
	req *pb.AddCourseReq,
)	(*pb.AddCourseResp, error) {
	role, err := getRoleFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = checkRole(role)
	if err != nil {
		return nil, err
	}

	

}

func (s *Server) DeleteCourse(
	ctx context.Context, 
	req *pb.DeleteCourseReq,
)	(*pb.DeleteCourseResp, error) {
	
}

func (s *Server) GetCourse(
	ctx context.Context, 
	req *pb.GetCourseReq,
)	(*pb.GetCourseResp, error) {
	
}

func (s *Server) UpdateCourse(
	ctx context.Context, 
	req *pb.UpdateCourseReq,
)	(*pb.UpdateCourseResp, error) {
	
}

func (s *Server) GetAllCourses(
	ctx context.Context, 
	req *pb.GetAllCoursesReq,
)	(*pb.GetAllCoursesResp, error) {
	
}

func getRoleFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Internal, "Can't get metada")
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