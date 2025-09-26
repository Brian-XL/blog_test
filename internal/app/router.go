package app

import (
	"github.com/Brian-XL/blog_test/internal/app/handler"
	"github.com/Brian-XL/blog_test/internal/app/middleware"
	"github.com/Brian-XL/blog_test/internal/app/repository"
	"github.com/Brian-XL/blog_test/internal/app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAllRoutes(r *gin.Engine, db *gorm.DB, jwtSecret string) {
	usr_repo := repository.GetNewUserRepository(db)
	usr_service := service.NewUserService(usr_repo)
	usr_handler := handler.GetNewUserHandler(usr_service)

	r.POST("/register", usr_handler.RegisterUser)
	r.POST("/login", usr_handler.Login)

	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware([]byte(jwtSecret)))
	{
		auth.GET("/info")
	}
}
