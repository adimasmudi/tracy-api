package services

import (
	"tracy-api/helper"
	"tracy-api/models"
	"tracy-api/repository"
)


type UserService interface {
	Signup(googleUser helper.GoogleUser) (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Signup(googleUser helper.GoogleUser)(models.User, error){

	var user models.User
	return user, nil
}