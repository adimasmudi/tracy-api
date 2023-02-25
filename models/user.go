package models

import (
	"time"
)

type User struct{
	Username 		string `json:"username"`
	NamaLengkap 	string `json:"namaLengkap"`
	Email 			string `json:"email"`
	Password 		string `json:"password"`
	NoHp 			string `json:"noHp"`
	DateOfBirth 	string `json:"dateOfBirth"`
	Picture 		string `json:"picture"`
	IsDataValid 	bool `json:"isDataValid"`
	IsPolice 		bool `json:"isPolice"`
	Alamat 			string `json:"alamat"`
	KodeInstansi 	string `json:"kodeInstansi"`
	CreatedAt 		time.Time `json:"createdAt"`
	UpdatedAt 		time.Time `json:"updatedAt"`
}