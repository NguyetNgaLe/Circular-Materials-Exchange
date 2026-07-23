package handler

import (
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pb"
	"context"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	user, token, err := h.svc.Register(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPhone())
	if err != nil {
		return nil, err
	}
	return &pb.AuthResponse{
		Token: token,
		User:  userToProto(user),
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user, token, err := h.svc.Login(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.AuthResponse{
		Token: token,
		User:  userToProto(user),
	}, nil
}

func (h *AuthHandler) VerifyToken(ctx context.Context, req *pb.TokenRequest) (*pb.UserResponse, error) {
	user, err := h.svc.VerifyToken(req.GetToken())
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{User: userToProto(user)}, nil
}

func (h *AuthHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := h.svc.GetUser(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{User: userToProto(user)}, nil
}

func (h *AuthHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	user, err := h.svc.GetUserByEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{User: userToProto(user)}, nil
}

func (h *AuthHandler) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserResponse, error) {
	user, err := h.svc.UpdateProfile(req.GetId(), req.GetName(), req.GetPhone(), req.GetAvatar())
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{User: userToProto(user)}, nil
}

func userToProto(u *repository.User) *pb.User {
	companyID := ""
	if u.CompanyID != nil {
		companyID = *u.CompanyID
	}
	return &pb.User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Role:      u.Role,
		Avatar:    u.Avatar,
		CompanyId: companyID,
	}
}
