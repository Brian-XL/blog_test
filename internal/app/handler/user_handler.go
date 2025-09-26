package handler

import (
	"net/http"
	"time"

	"github.com/Brian-XL/blog_test/internal/app/service"
	"github.com/Brian-XL/blog_test/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	Service      *service.UserService
	JWTSecretKey []byte
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func GetNewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	if err := u.Service.RegisterUser(user); err != nil {
		c.JSON(400, gin.H{"Failed to register": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (u *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"binding error": err.Error()})
		return
	}

	user, error := u.Service.Login(req.Email, req.Password)
	if error != nil {
		c.JSON(400, gin.H{"Failed to Login": error.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(u.JWTSecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
		"msg":   "login successfully",
	})
}
