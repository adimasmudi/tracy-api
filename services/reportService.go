package services

import (
	"context"
	"time"
	"tracy-api/inputs"
	"tracy-api/models"
	"tracy-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportService interface {
	CreateReport(ctx context.Context, email string,  input inputs.CreateReportInput) (*mongo.InsertOneResult, error)
	GetById(ctx context.Context, id primitive.ObjectID) (interface{}, error)
}

type reportService struct{
	repository repository.ReportRepository
	userRepository repository.UserRepository
	policeRepository repository.PoliceStationRepository
}

func NewReportService(repository repository.ReportRepository, userRepository repository.UserRepository, policeRepository repository.PoliceStationRepository) *reportService{
	return &reportService{repository, userRepository, policeRepository}
}

func (s *reportService) CreateReport(ctx context.Context, email string, input inputs.CreateReportInput) (*mongo.InsertOneResult, error) {
	
	report := bson.M{
		"jenisKejahatan" : input.JenisKejahatan,
		"uraian" : input.Uraian,
		"emailUser" : email,
		"emailPolisi" : input.EmailPolisi,
		"status" : "terkirim",
		"reportedAt" : time.Now(),
	}

	result, err := s.repository.Save(ctx, report)

	if err!= nil{
		return result, err
	}

	return result, nil
}

func (s *reportService) GetById(ctx context.Context, id primitive.ObjectID) (interface{}, error){
	var report models.Report

	report, err := s.repository.GetById(ctx, id)

	if err != nil{
		return report, err
	}

	user, err := s.userRepository.FindByEmail(ctx, report.EmailUser)

	if err != nil{
		return report, err
	}

	police, err := s.policeRepository.FindByEmail(ctx, report.EmailPolisi)

	if err != nil{
		return report, err
	}

	result := bson.M{
		"report" : report,
		"user" : user,
		"police" : police,
	}

	return result, nil
}