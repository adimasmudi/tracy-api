package services

import (
	"context"
	"errors"
	"os"
	"time"
	"tracy-api/helper"
	"tracy-api/models"
	"tracy-api/repository"

	"github.com/golang-jwt/jwt/v4"
)


type UserService interface {
	Signup(ctx context.Context, googleUser helper.GoogleUser) (models.User, string,error)
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

		token,err := generateToken(googleUser)

		if err != nil{
			return newUser, "",err
		}

		return newUser,token, nil
	}

	userFound, err := s.repository.FindByEmail(ctx,googleUser.Email)

	
	if err != nil{
		return userFound,"", err
	}

	token,err := generateToken(googleUser)

	if err != nil{
		return userFound, "",err
	}

	return userFound, token, nil
}

func generateToken(payload helper.GoogleUser)(string, error){
	claim := jwt.MapClaims{}
	claim["google_id"] = payload.Id
	claim["email"] = payload.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil{
		return signedToken, err
	}

	return signedToken, nil
}

func validateToken(encodedToken string)(*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}