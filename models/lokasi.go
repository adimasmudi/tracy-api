package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lokasi struct {
	IdLaporan primitive.ObjectID `json:"idLaporan"`
	DetailLokasi string `json:"detailLokasi"`
}