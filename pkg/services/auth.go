package services

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
	
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
	var userDetails models.User

	if result := s.H.DB.Where(&models.User{EmailAddress: req.EmailAddress}).First(&userDetails); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "User already exists",
		}, nil
	}

	user := &models.User{
		// ID:   uuid.New(),
		// ID:   uuid.New(),
		Name: req.Name, EmailAddress: req.EmailAddress, PhoneNumber: req.PhoneNumber, Password: utils.HashPassword(req.Password),
		DateOfBirth: req.DateOfBirth, CreditScore: req.CreditScore, IsFinanceWorthy: req.IsFinanceWorthy, WorkPlace: req.WorkPlace, NIN: req.Nin, EmployerName: req.EmployerName,
		SalaryScale: req.SalaryScale, IsKYCVerified: req.KycVerified,
	}

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{EmailAddress: req.EmailAddress}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User acount not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Wrong Password, Try again",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
		Error:  "dada",
		UserId: "dara",
		Name:   user.Name,
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
