package inputs

type UpdateUserInput struct {
	UserName    string `json:"username" binding:"required"`
	NamaLengkap string `json:"namalengkap" binding:"required"`
	NoHp        string `json:"nohp" binding:"required"`
	DateOfBirth string `json:"dateofbirth" binding:"required"`
	Alamat      string `json:"alamat" binding:"required"`
}