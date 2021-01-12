package server

import (
	"google.golang.org/grpc/metadata"
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/0B1t322/auth-service/models/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/0B1t322/auth-service/authservice"
	"github.com/0B1t322/auth-service/pkg/auth"
)

// Server is auth microservice struct
type Server struct {
	
	authManager *auth.AuthManager

	// To provide a standart service interfave
	pb.UnimplementedAuthServiceServer
}

// NewServer create a server
// 	params:
//		signingKey - key for hashing JWT
// 		hashSalt - salt for unhash password (now not use)
// 		expireDuration - duration of JWT token
func NewServer(authManager *auth.AuthManager) *Server {
	return &Server{
		authManager: authManager,
	}

}
// SignUp check if user not exsist
// if not - create them and give token
func (s *Server) SignUp(
	ctx context.Context, 
	req *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	// check user
	u, err := s.checkUserInDBAndGet(req.Username, req.Password)
	if err == ErrIncorrectUserNamePass {
		return &pb.AuthResponse{
			Token: "", 
			Error: ErrIncorrectUserNamePass.Error(),
		}, nil
	} else if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
			Error: status.Error(
				codes.Internal,
				"Internal",
			).Error(),
		}, err
	}
	token, err := s.authManager.CreateToken(u)
	if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
			Error: status.Error(
				codes.Internal,
				"Internal",
			).Error(),
		}, err
	}

	return &pb.AuthResponse{
		Token: token,
		Error: "",
	}, nil
}

func (s *Server) SignIn(
	ctx context.Context,
	req *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	var role string
	if len(md["role"]) > 0 {
		role = md["role"][0]
	} else {
		ok = false
	}
	
	if !ok {
		log.Warn("Error of get role from ctx")
		log.Infoln("role is: ",role)
		log.Infoln(ctx.Value("role"))
		return &pb.AuthResponse{
			Token: "",
			Error: status.Error(
				codes.Internal,
				"Internal",
			).Error(),
		}, errors.New("Error of get role from ctx")
	}

	u := user.NewUser(req.Username, req.Password, role)
	err := u.AddUser()
	if err == user.ErrUserExsist {
		return &pb.AuthResponse{
			Token: "",
			Error: err.Error(),
		}, nil
	} else if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
			Error: status.Error(
				codes.Internal,
				"Internal",
			).Error(),
		}, err
	}
	
	token, err := s.authManager.CreateToken(u)
	if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
			Error: status.Error(
				codes.Internal,
				"Internal",
			).Error(),
		}, err
	}
	return &pb.AuthResponse{
		Token: token,
		Error: "",
	}, nil
}

func (s *Server) checkUserInDBAndGet(
	username string,
	password string,
) (*user.User, error) {
	u, err := user.GetUserByUserName(username)
	if err == user.ErrUserNotFound {
		return nil, ErrIncorrectUserNamePass
	} else if err != nil {
		return nil, err
	}

	// check password
	if u.Password != password {
		return nil, ErrIncorrectUserNamePass
	}

	return u, nil
}

