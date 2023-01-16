package services

import (
	"context"
	"log"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
	"net/http"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.AuthServiceServer
}

// mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer
func (*Server) mustEmbedUnimplementedAuthServiceServer() {
	panic("unimplemented")
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// var userDetails models.User

	// if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&userDetails); result.Error == nil {
	// 	return &pb.RegisterResponse{
	// 		Status: http.StatusConflict,
	// 		Error:  "User already exists",
	// 	}, nil
	// }
	// user := &models.User{
	// 	Name:        req.Name,
	// 	Email:       req.Email,
	// 	PhoneNumber: req.PhoneNumber,
	// 	Password:    utils.HashPassword(req.Password),
	// }
	log.Printf("Received register request: %s", req)

	// user := &models.User{Name: "Ronnie Zad", EmailAddress: "zadcorna@gmail.com", PhoneNumber: "+256702703612", Password: utils.HashPassword(req.Password)}
	// s.H.DB.Create(&models.User{
	// 	Name:        "req.Name",
	// 	Email:       req.Email,
	// 	PhoneNumber: req.PhoneNumber,
	// 	Password:    utils.HashPassword(req.Password),
	// })

	// s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{EmailAddress: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User account not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User account not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
		Error:  "dada",
		UserId: "dara",
		Name:   "daea",
		Email:  user.EmailAddress,
		Phone:  user.PhoneNumber,
		RoleId: "3",
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{EmailAddress: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: 1,
	}, nil

}
