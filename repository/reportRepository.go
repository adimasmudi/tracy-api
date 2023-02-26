package repository

import (
	"context"
	"encoding/json"
	"tracy-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportRepository interface {
	Save(ctx context.Context, report primitive.M) (*mongo.InsertOneResult, error)
	GetById(ctx context.Context, id primitive.ObjectID) (models.Report, error)
	GetAll(ctx context.Context) ([]models.Report, error)
	GetAllCurrentUser(ctx context.Context, email string) ([]models.Report, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, newStatus primitive.M) (*mongo.UpdateResult, error)
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

func (r *reportRepository) GetAll(ctx context.Context) ([]models.Report, error){
	var results []models.Report

	cur, err := r.DB.Find(ctx,bson.D{{}})

	if err != nil{
		return results, err
	}

	if err = cur.All(ctx, &results);err != nil{
		return results, err
	}

	for _, result := range results{
		cur.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			return results, err
		}
		
	}

	return results, nil
}

func (r *reportRepository) GetAllCurrentUser(ctx context.Context, email string) ([]models.Report, error){
	var results []models.Report

	cur, err := r.DB.Find(ctx,bson.M{"emailUser" : email})

	if err != nil{
		return results, err
	}

	if err = cur.All(ctx, &results);err != nil{
		return results, err
	}

	for _, result := range results{
		cur.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			return results, err
		}
		
	}

	return results, nil
}

func (r *reportRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, newStatus primitive.M) (*mongo.UpdateResult, error){

	result, err := r.DB.UpdateOne(ctx,bson.M{"_id":id},bson.M{"$set" : newStatus})

	if err != nil{
		return result, err
	}

	return result, nil
}