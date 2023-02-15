package inputs

type PoliceStationInput struct {
	NamaKantor string `json:"namaKantor" form:"namaKantor" binding:"required"`
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required"`
	Alamat     string `json:"alamat" form:"alamat" binding:"required"`
	Telepon    string `json:"telepon" form:"telepon" binding:"required"`
}

type PoliceStationLoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}