package repository

import (
	"context"
	"errors"
	"tracy-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PoliceStationRepository interface {
	Create(ctx context.Context, policeStation primitive.M) (*mongo.InsertOneResult, error)
	FindByEmail(ctx context.Context, email string) (models.PoliceStation,  error)
	IsPoliceStationExist(ctx context.Context, email string, username string) (bool, error)
}

type policeStationRepository struct{
	DB *mongo.Collection
}

func NewPoliceStationRepository(DB *mongo.Collection) *policeStationRepository{
	return &policeStationRepository{DB}
}

func (r *policeStationRepository) Create(ctx context.Context, policeStation primitive.M) (*mongo.InsertOneResult, error){
	r.DB.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys : bson.D{{Key: "email", Value: 1},{Key:"username", Value:1}},
			Options : options.Index().SetUnique(true),
		},
	)

	result, err := r.DB.InsertOne(ctx, policeStation)

	if err != nil{
		return result, err
	}

	return result, nil
}

func (r *policeStationRepository) FindByEmail(ctx context.Context, email string) (models.PoliceStation,  error){

	var policeStation models.PoliceStation

	err := r.DB.FindOne(ctx, bson.M{"email": email}).Decode(&policeStation)

	if err != nil{
		return policeStation, err
	}

	return policeStation, nil
}

func (r *policeStationRepository) IsPoliceStationExist(ctx context.Context, email string, username string) (bool, error){
	var policeStation models.PoliceStation

	err := r.DB.FindOne(ctx, bson.M{"email": email, "username" : username}).Decode(&policeStation)

	if err != nil{
		return false, err
	}

	return true, errors.New("user exist")
}