package formatter

import (
	"tracy-api/models"
)

type UserFormatter struct {
	Username string
	NamaLengkap string
	Email string
	Password string
	NoHp string
	DateOfBirth string
	Picture string
	IsDataValid bool
	Alamat string
}

func FormatUser(user models.User) UserFormatter {
	formatter := UserFormatter{
		Username : user.Username,
		NamaLengkap : user.NamaLengkap,
		Email : user.Email ,
		Password : user.Password,
		NoHp : user.NoHp ,
		DateOfBirth : user.DateOfBirth ,
		Picture : user.Picture ,
		IsDataValid : user.IsDataValid ,
		Alamat : user.Alamat ,
	}

	return formatter
}