package services

import (
	"context"
	"tracy-api/inputs"
	"tracy-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LokasiService interface {
	SaveLocation(ctx context.Context, input inputs.AddLokasiInput) (*mongo.InsertOneResult, error)
}

type lokasiService struct{
	repository repository.LokasiRepository
}

func NewLokasiService(repository repository.LokasiRepository) *lokasiService{
	return &lokasiService{repository}
}

func (s *lokasiService) SaveLocation(ctx context.Context, input inputs.AddLokasiInput) (*mongo.InsertOneResult, error){
	lokasi := bson.M{
		"idLaporan" : input.IdLaporan,
		"latitude" : input.Latitude,
		"longitude" : input.Longitude,
		"detailLokasi" : input.DetailLokasi,
	}

	result, err := s.repository.Save(ctx, lokasi)

	if err != nil{
		return result, err
	}

	return result, nil
}