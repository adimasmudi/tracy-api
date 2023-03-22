package services

import (
	"context"
	"errors"
	"time"
	"tracy-api/helper"
	"tracy-api/inputs"
	"tracy-api/models"
	"tracy-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


type UserService interface {
	Signup(ctx context.Context, googleUser helper.GoogleUser) (models.User, string,error)
	Register(ctx context.Context, input inputs.RegisterUserInput) (*mongo.InsertOneResult,error)
	Login(ctx context.Context, input inputs.LoginUserInput) (models.User, string, error)
	GetProfile(ctx context.Context, email string) (models.User, error)
	UpdateProfile(ctx context.Context, email string, input inputs.UpdateUserInput) (*mongo.UpdateResult, error)
}

type userService struct {
	repository repository.UserRepository
	policeRepository repository.PoliceStationRepository
}

func NewUserService(repository repository.UserRepository, policeRepository repository.PoliceStationRepository) *userService {
	return &userService{repository, policeRepository}
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

func (s *userService) Register(ctx context.Context, input inputs.RegisterUserInput) (*mongo.InsertOneResult, error){
	
	userExist, _ := s.repository.IsUserExist(ctx,input.Email )

	if userExist{
		return nil, errors.New("User already exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil{
		return nil, err
	}


	newUser := models.User{
		Username : input.Username,
		NamaLengkap : input.NamaLengkap,
		Email : input.Email,
		Password : string(passwordHash),
		NoHp : input.NoHp,
		DateOfBirth : input.DateOfBirth,
		IsDataValid: true,
		Alamat : input.Alamat,
		CreatedAt: time.Now(),
		UpdatedAt : time.Now(),
	}

	registeredUser, err := s.repository.Save(ctx,newUser)

	if err != nil{
		return nil, err
	}

	return registeredUser, nil
}

func (s *userService) Login(ctx context.Context, input inputs.LoginUserInput) (models.User, string, error){

	user, err := s.repository.FindByEmail(ctx,input.Email)

	if err != nil{
		return user, "", errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil{
		return user, "", errors.New("wrong Password")
	}

	token, err := helper.GenerateToken(user.Email)

	if err != nil{
		return user, "", errors.New("can't generate token")
	}

	return user, token, nil
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
	
	isKodeInstansiExist, _ := s.policeRepository.IsKodeInstansiExist(ctx, input.KodeInstansi)

	if isKodeInstansiExist{
		updateUser["ispolice"] = true
	}else{
		updateUser["ispolice"] = false
	}

	updateUser["kodeinstansi"] = input.KodeInstansi
		
	

	user, err := s.repository.UpdateProfile(ctx,email, updateUser)

	if err != nil{
		return user, err
	}

	return user, nil
}
