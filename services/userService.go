package services

import (
	"context"
	"time"
	"tracy-api/helper"
	"tracy-api/inputs"
	"tracy-api/models"
	"tracy-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserService interface {
	Signup(ctx context.Context, googleUser helper.GoogleUser) (models.User, string,error)
	GetProfile(ctx context.Context, email string) (models.User, error)
	UpdateProfile(ctx context.Context, email string, input inputs.UpdateUserInput) (*mongo.UpdateResult, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Signup(ctx context.Context, googleUser helper.GoogleUser)(models.User, string,error){

	userExist, _ := s.repository.IsUserExist(ctx,googleUser.Email)

	if !userExist{
		newUser := models.User{
			Username : "",
			NamaLengkap : "",
			Email : googleUser.Email,
			Password : "",
			NoHp : "",
			DateOfBirth : "",
			Picture : googleUser.Picture,
			IsDataValid: false,
			Alamat : "",
			CreatedAt: time.Now(),
			UpdatedAt : time.Now(),
		}

		_, err := s.repository.Save(ctx,newUser)

		if err != nil{
			return newUser,"", err
		}

		token,err := helper.GenerateToken(googleUser.Email)

		if err != nil{
			return newUser, "",err
		}

		return newUser,token, nil
	}

	userFound, err := s.repository.FindByEmail(ctx,googleUser.Email)

	
	if err != nil{
		return userFound,"", err
	}

	token,err := helper.GenerateToken(googleUser.Email)

	if err != nil{
		return userFound, "",err
	}

	return userFound, token, nil
}

func (s *userService) GetProfile(ctx context.Context, email string) (models.User, error){
	var user models.User
	user, err := s.repository.FindByEmail(ctx,email)

	if err != nil{
		return user, err
	}

	return user, nil
}

func (s *userService) UpdateProfile(ctx context.Context, email string, input inputs.UpdateUserInput) (*mongo.UpdateResult, error){

	updateUser := bson.M{
		"username" :input.UserName, 
		"namalengkap" : input.NamaLengkap, 
		"dateofbirth" : input.DateOfBirth,
		"nohp" : input.NoHp,
		"alamat" : input.Alamat, 
		"isdatavalid" : true,
		"updatedat" : time.Now(),
	}
	user, err := s.repository.UpdateProfile(ctx,email, updateUser)

	if err != nil{
		return user, err
	}

	return user, nil
}
