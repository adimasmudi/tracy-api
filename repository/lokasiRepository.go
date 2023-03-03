package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LokasiRepository interface {
	Save(ctx context.Context, lokasi primitive.M) (*mongo.InsertOneResult, error)
}

type lokasiRepository struct{
	DB *mongo.Collection
}

func NewLokasiRepository(DB *mongo.Collection) *lokasiRepository{
	return &lokasiRepository{DB}
}

func (r *lokasiRepository) Save(ctx context.Context, lokasi primitive.M) (*mongo.InsertOneResult, error){
	result, err := r.DB.InsertOne(ctx, lokasi)

	if err != nil{
		return result, err
	}

	return result, nil
}