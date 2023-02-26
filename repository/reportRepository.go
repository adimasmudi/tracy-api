package repository

import (
	"context"
	"tracy-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportRepository interface {
	Save(ctx context.Context, report primitive.M) (*mongo.InsertOneResult, error)
	GetById(ctx context.Context, id primitive.ObjectID) (models.Report, error)
	// get all
	// get all for current user
}

type reportRepository struct{
	DB *mongo.Collection
}

func NewReportRepository(DB *mongo.Collection) *reportRepository{
	return &reportRepository{DB}
}

func (r *reportRepository) Save(ctx context.Context,report primitive.M) (*mongo.InsertOneResult, error) {
	
	result,err := r.DB.InsertOne(ctx, report)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *reportRepository) GetById(ctx context.Context, id primitive.ObjectID) (models.Report, error){
	var reportModel models.Report

	err := r.DB.FindOne(ctx, bson.M{"_id" : id}).Decode(&reportModel)

	if err != nil{
		return reportModel, err
	}

	return reportModel, nil
}