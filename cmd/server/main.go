package main

import (
	"github.com/Brian-XL/blog_test/internal/app"
	"github.com/Brian-XL/blog_test/internal/config"
	"github.com/Brian-XL/blog_test/internal/model"
	"github.com/Brian-XL/blog_test/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDB(cfg)

	db.AutoMigrate(model.User{}, model.Post{}, model.Comment{})

	r := gin.Default()

	app.RegisterAllRoutes(r, db, cfg.JWT.Secret)

	r.Run(":8080")
}
