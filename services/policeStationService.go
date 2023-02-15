package services

import (
	"context"
	"errors"
	"math/rand"
	"time"
	"tracy-api/helper"
	"tracy-api/inputs"
	"tracy-api/models"
	"tracy-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type PoliceStationService interface {
	Save(ctx context.Context, input inputs.PoliceStationInput, filename string) (*mongo.InsertOneResult, error)
	Login(ctx context.Context, input inputs.PoliceStationLoginInput) (models.PoliceStation, string, error)
	GetProfile(ctx context.Context, email string) (models.PoliceStation, error)
}

type policeStationService struct{
	repository repository.PoliceStationRepository
}

func NewPoliceStationService(repository repository.PoliceStationRepository) *policeStationService{
	return &policeStationService{repository}
}

func (s *policeStationService) Save(ctx context.Context, input inputs.PoliceStationInput, filename string) (*mongo.InsertOneResult, error){

	policeStationExist, err := s.repository.IsPoliceStationExist(ctx, input.Email, input.Username)

	if policeStationExist{
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil{
		return nil, err
	}

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&?"

	kode := make([]byte, 15)
    for i := range kode {
        kode[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    
	kode_instansi := string(kode)

	registeredUser := bson.M{
		"namaKantor" : input.NamaKantor,
		"username" : input.Username,
		"email" : input.Email,
		"password" : string(passwordHash),
		"telepon" : input.Telepon,
		"alamat" : input.Alamat,
		"picture" : filename,
		"kodeInstansi" : kode_instansi,
		"createdAt": time.Now(),
		"updatedAt" : time.Now(),
	}

	result, err := s.repository.Create(ctx,registeredUser)

	if err!= nil{
		return result, err
	}

	return result, nil
}

func (s *policeStationService) Login(ctx context.Context, input inputs.PoliceStationLoginInput) (models.PoliceStation, string, error){

	police, err := s.repository.FindByEmail(ctx,input.Email)

	if err != nil{
		return police, "", errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(police.Password), []byte(input.Password))

	if err != nil{
		return police, "", errors.New("Wrong Password")
	}

	token, err := helper.GenerateToken(police.Email)

	if err != nil{
		return police, "", errors.New("Can't generate token")
	}

	return police, token, nil
}

func (s *policeStationService) GetProfile(ctx context.Context, email string) (models.PoliceStation, error){
	var police models.PoliceStation
	police, err := s.repository.FindByEmail(ctx,email)

	if err != nil{
		return police, err
	}

	return police, nil
}