package handler

import (
	authpb "api-gateway/internal/pb/auth"
	companypb "api-gateway/internal/pb/company"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth    authpb.AuthServiceClient
	company companypb.CompanyServiceClient
}

func NewAuthHandler(auth authpb.AuthServiceClient, company companypb.CompanyServiceClient) *AuthHandler {
	return &AuthHandler{auth: auth, company: company}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.auth.Login(ctx, &authpb.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		writeRPCError(c, err, "Email hoac mat khau khong dung")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token": response.GetToken(),
			"user":  h.userJSON(ctx, response.GetUser()),
		},
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Phone    string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.auth.Register(ctx, &authpb.RegisterRequest{
		Name: req.Name, Email: req.Email, Password: req.Password, Phone: req.Phone,
	})
	if err != nil {
		writeRPCError(c, err, "Loi tao tai khoan")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token": response.GetToken(),
			"user":  h.userJSON(ctx, response.GetUser()),
		},
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.auth.GetUser(ctx, &authpb.GetUserRequest{Id: stringValue(userID)})
	if err != nil {
		writeRPCError(c, err, "Khong tim thay nguoi dung")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": h.userJSON(ctx, response.GetUser())})
}

func (h *AuthHandler) userJSON(ctx context.Context, user *authpb.User) gin.H {
	if user == nil {
		return gin.H{}
	}
	companyID := user.GetCompanyId()
	if companyID == "" {
		if company, err := h.company.GetCompanyByOwner(ctx, &companypb.GetCompanyByOwnerRequest{OwnerId: user.GetId()}); err == nil {
			companyID = company.GetId()
		}
	}
	return gin.H{
		"id":        user.GetId(),
		"name":      user.GetName(),
		"email":     user.GetEmail(),
		"phone":     user.GetPhone(),
		"role":      user.GetRole(),
		"avatar":    user.GetAvatar(),
		"companyId": companyID,
	}
}
