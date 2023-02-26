package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Report struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	JenisKejahatan string    `json:"jenisKejahatan"`
	Uraian         string    `json:"uraian"`
	EmailUser      string    `json:"user"`
	EmailPolisi    string    `json:"polisi"`
	Status         string    `json:"status"`
	ReportedAt     time.Time `json:"reportedAt"`
}