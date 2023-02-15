package formatter

import (
	"tracy-api/models"
)

type UserFormatter struct {
	username string
	namaLengkap string
	email string
	password string
	noHp string
	dateOfBirth string
	picture string
	isDataValid bool
	alamat string
}

func FormatUser(user models.User) UserFormatter {
	formatter := UserFormatter{
		username : user.Username,
		namaLengkap : user.NamaLengkap,
		email : user.Email ,
		password : user.Password,
		noHp : user.NoHp ,
		dateOfBirth : user.DateOfBirth ,
		picture : user.Picture ,
		isDataValid : user.IsDataValid ,
		alamat : user.Alamat ,
	}

	return formatter
}