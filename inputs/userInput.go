package inputs

type RegisterUserInput struct {
	Username    string `json:"username" binding:"required"`
	NamaLengkap string `json:"namaLengkap" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NoHp        string `json:"noHp" binding:"required"`
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
	Alamat      string `json:"alamat" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	UserName     string `json:"username" binding:"required"`
	NamaLengkap  string `json:"namalengkap" binding:"required"`
	NoHp         string `json:"nohp" binding:"required"`
	DateOfBirth  string `json:"dateofbirth" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
	KodeInstansi string `json:"kodeInstansi"`
}