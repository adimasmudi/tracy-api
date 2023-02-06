package services

import "tracy-api/models"

type UserService interface {
	Login() (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Login()(models.User, error){
	return interface{}, nil
}