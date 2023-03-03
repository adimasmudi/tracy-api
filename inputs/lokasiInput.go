package inputs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lokasi struct {
}

type AddLokasiInput struct {
	IdLaporan    primitive.ObjectID `json:"idLaporan" binding:"required"`
	Latitude     string             `json:"latitude" binding:"required"`
	Longitude    string             `json:"longitude" binding:"required"`
	DetailLokasi string             `json:"detailLokasi" binding:"required"`
}