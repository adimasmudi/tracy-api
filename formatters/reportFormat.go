package formatter

import (
	"time"
	"tracy-api/models"
)

type ReportFormatter struct {
	jenisKejahatan string    
	uraian         string   
	user      interface{}    
	polisi    interface{} 
	status         string    
	createdAt     time.Time
	updatedAt time.Time
}

func FormatReport(report models.Report, user models.User, police models.PoliceStation) interface{} {
	formatter := ReportFormatter{
		jenisKejahatan : report.JenisKejahatan,
		uraian : report.Uraian,
		user : user,
		polisi : police,
		status : report.Status,
		createdAt : report.CreatedAt,
		updatedAt : report.UpdatedAt,
	}

	return formatter
}