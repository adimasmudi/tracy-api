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
	GetAll(ctx context.Context) ([]interface{}, error)
	GetAllByCurrentUser(ctx context.Context, email string) ([]interface{}, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, input inputs.UpdateStatusReport) (*mongo.UpdateResult, error)
}

type reportService struct{
	repository repository.ReportRepository
	userRepository repository.UserRepository
	policeRepository repository.PoliceStationRepository
	lokasiRepository repository.LokasiRepository
}

func NewReportService(repository repository.ReportRepository, userRepository repository.UserRepository, policeRepository repository.PoliceStationRepository, lokasiRepository repository.LokasiRepository) *reportService{
	return &reportService{repository, userRepository, policeRepository, lokasiRepository}
}

func (s *reportService) CreateReport(ctx context.Context, email string, input inputs.CreateReportInput) (*mongo.InsertOneResult, error) {
	
	report := bson.M{
		"jenisKejahatan" : input.JenisKejahatan,
		"uraian" : input.Uraian,
		"emailUser" : email,
		"emailPolisi" : input.EmailPolisi,
		"status" : "terkirim",
		"createdAt" : time.Now(),
		"updatedAt" : time.Now(),
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

	lokasi, err := s.lokasiRepository.GetByReportId(ctx, report.Id)

	if err != nil{
		return report, err
	}

	result := bson.M{
		"report" : report,
		"user" : user,
		"police" : police,
		"lokasi" : lokasi,
	}

	return result, nil
}

func (s *reportService) GetAll(ctx context.Context) ([]interface{}, error){
	var result []interface{}
	allReport, err := s.repository.GetAll(ctx)

	if err != nil{
		return result, err
	}

	for _, report := range allReport{
		user, _ := s.userRepository.FindByEmail(ctx,report.EmailUser)
		police, _ := s.policeRepository.FindByEmail(ctx, report.EmailPolisi)
		lokasi, _ := s.lokasiRepository.GetByReportId(ctx, report.Id)

		format := make([]interface{},4)

		format[0] = report
		format[1] = user
		format[2] = police
		format[3] = lokasi

		result = append(result, format)
	}

	return result, nil
}

func (s *reportService) GetAllByCurrentUser(ctx context.Context, email string) ([]interface{}, error){
	var result []interface{}

	reportsByCurrentUser, err := s.repository.GetAllCurrentUser(ctx, email)

	if err != nil{
		return result, err
	}

	for _, report := range reportsByCurrentUser{
		user, _ := s.userRepository.FindByEmail(ctx,report.EmailUser)
		police, _ := s.policeRepository.FindByEmail(ctx, report.EmailPolisi)
		lokasi, _ := s.lokasiRepository.GetByReportId(ctx, report.Id)

		format := make([]interface{},4)

		format[0] = report
		format[1] = user
		format[2] = police
		format[3] = lokasi

		result = append(result, format)
	}


	return result, nil
}

func (s *reportService) UpdateStatus(ctx context.Context, id primitive.ObjectID, input inputs.UpdateStatusReport) (*mongo.UpdateResult, error){

	newReportStatus := bson.M{
		"status" : input.Status,
		"updatedAt" : time.Now(),
	}

	updatedStatus, err := s.repository.UpdateStatus(ctx, id, newReportStatus)

	if err != nil{
		return updatedStatus, err
	}

	return updatedStatus, nil
}