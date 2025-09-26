package repository

import (
	"errors"
	"fmt"

	"github.com/Brian-XL/blog_test/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func GetNewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateNewUser(user model.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", name).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindUserByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error

	return &user, err
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found with email.", email)
		} else {
			fmt.Println(err)
		}
	}
	return &user, err
}
