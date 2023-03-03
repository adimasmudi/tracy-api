package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lokasi struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	IdLaporan primitive.ObjectID `json:"idLaporan"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	DetailLokasi string `json:"detailLokasi"`
}