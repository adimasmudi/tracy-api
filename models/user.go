package models

import (
	"time"
)

type User struct{
	Username string
	NamaLengkap string
	Email string
	Password string
	NoHp string
	DateOfBirth string
	Picture string
	IsDataValid bool
	Alamat string
	CreatedAt time.Time
	UpdatedAt time.Time
}