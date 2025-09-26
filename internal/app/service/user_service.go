package service

import (
	"errors"

	"github.com/Brian-XL/blog_test/internal/app/repository"
	"github.com/Brian-XL/blog_test/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repository: repo}
}

func (s *UserService) RegisterUser(user model.User) error {
	if _, err := s.Repository.FindUserByName(user.Username); err == nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := s.Repository.CreateNewUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(email string, password string) (*model.User, error) {
	user, err := s.Repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.Repository.FindUserByID(id)
}
