package server

import (
	"context"

	"github.com/0B1t322/distanceLearningWebSite/pkg/marshall"
	"google.golang.org/grpc/metadata"

	log "github.com/sirupsen/logrus"

	"github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/0B1t322/service.auth/authservice"
	"github.com/0B1t322/service.auth/pkg/auth"
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

/*
SignIn create a user if they not exsist
	require: a role in ctx with key "role"
*/
func (s *Server) SignIn(
	ctx context.Context, 
	req *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	// check user
	u, err := s.checkUserInDBAndGet(req.Username, req.Password)
	if err == ErrIncorrectUserNamePass {
		return &pb.AuthResponse{
			Token: "", 
		}, status.Error(codes.Unauthenticated, err.Error())
	} else if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
		}, status.Error(
			codes.Internal,
			err.Error(),
		)
	}
	token, err := s.authManager.CreateToken(u)
	if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
		}, status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	return &pb.AuthResponse{
		Token: token,
	}, status.Error(codes.OK, "You signup")
}


// SignUp check if user not exsist
// if not - give error: ErrIncorrectUserNamePass
func (s *Server) SignUp(
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
		}, status.Error(
			codes.Internal,
			"Error of get role from ctx",
		)
	}

	u := user.NewUser(req.Username, req.Password, role)
	err := u.AddUser()
	if err == user.ErrUserExsist {
		return &pb.AuthResponse{
			Token: "",
		}, status.Error(codes.AlreadyExists, err.Error())
	} else if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
		},
		status.Error(codes.Internal, err.Error())
	}
	
	token, err := s.authManager.CreateToken(u)
	if err != nil {
		log.Error(err)
		return &pb.AuthResponse{
			Token: "",
		}, status.Error(
			codes.Internal,
			err.Error(),
		)
	}
	return &pb.AuthResponse{
		Token: token,
	}, status.Error(codes.OK, "You sucsessfully sign in")
}

/*
Check - check token 
*/
func (s *Server) Check(
	ctx context.Context, 
	req *pb.Token,
) (*pb.TokenInfo , error) {
	tokenInfo, err := s.authManager.ParseToken(req.Token)
	if err == auth.ErrInvalidToken {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	} else if err == auth.ErrTokenExpire {
		return nil, status.Error(
			codes.Unauthenticated,
			err.Error(),
		)
	} else if err != nil {
		return nil, status.Error(
			codes.Internal,
			err.Error(),
		)
	}
	// TODO сделать через преобразование информазии в json структру и обратно в новую
	return s.unmarshallTokenInfo(tokenInfo)
}

func (s *Server) unmarshallTokenInfo(tokenInfo *auth.TokenInfo) (*pb.TokenInfo, error) {
	res := &pb.TokenInfo{}
	err := marshall.Marshall(tokenInfo, res)
	if err != nil {
		log.Errorf("Error on marshalling: %v" ,err)
		return nil, status.Errorf(codes.Internal, "Internal server err")
	}

	return res, err
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


