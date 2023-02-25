package models

import (
	"time"
)

type PoliceStation struct {
	NamaKantor   	string `json:"namaKantor"`
	Username     	string `json:"username"`
	Password     	string `json:"password"`
	Email        	string `json:"email"`
	Alamat       	string `json:"alamat"`
	Telepon      	string `json:"noTelepon"`
	Picture      	string `json:"picture"`
	KodeInstansi 	string `json:"kodeInstansi"`
	CreatedAt 		time.Time `json:"createdAt"`
	UpdatedAt 		time.Time `json:"updatedAt"`
}