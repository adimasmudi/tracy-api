package repository

import (
	"context"
	"tracy-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LokasiRepository interface {
	Save(ctx context.Context, lokasi primitive.M) (*mongo.InsertOneResult, error)
	GetByReportId(ctx context.Context, reportId primitive.ObjectID) (models.Lokasi, error)
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

func (r *lokasiRepository) GetByReportId(ctx context.Context, reportId primitive.ObjectID) (models.Lokasi, error){
	var lokasiModel models.Lokasi
	err := r.DB.FindOne(ctx,bson.M{"idLaporan" : reportId}).Decode(&lokasiModel)

	if err != nil{
		return lokasiModel, err
	}

	return lokasiModel, nil
}