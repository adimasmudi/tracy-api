package models

import "time"

type PoliceStation struct {
	NamaKantor   string
	Username     string
	Password     string
	Email        string
	Alamat       string
	Telepon      string
	Picture       string
	KodeInstansi string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}