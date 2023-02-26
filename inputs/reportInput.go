package inputs

type CreateReportInput struct {
	JenisKejahatan string `json:"jenisKejahatan" binding:"required"`
	Uraian         string `json:"uraian" binding:"required"`
	EmailPolisi    string `json:"emailPolisi" binding:"required"`
}

type UpdateStatusReport struct {
	Status string `json:"status" binding:"required"`
}
