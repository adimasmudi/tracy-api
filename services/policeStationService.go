package services

import (
	"context"
	"math/rand"
	"time"
	"tracy-api/inputs"
	"tracy-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type PoliceStationService interface {
	Save(ctx context.Context, input inputs.PoliceStationInput, filename string) (*mongo.InsertOneResult, error)
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
		"nama_kantor" : input.NamaKantor,
		"username" : input.Username,
		"email" : input.Email,
		"password" : string(passwordHash),
		"telepon" : input.Telepon,
		"alamat" : input.Alamat,
		"picture" : filename,
		"kode_instansi" : kode_instansi,
		"createdAt": time.Now(),
		"updatedAt" : time.Now(),
	}

	result, err := s.repository.Create(ctx,registeredUser)

	if err!= nil{
		return result, err
	}

	return result, nil

}