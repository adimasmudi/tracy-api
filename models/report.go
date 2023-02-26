package models

import "time"

type Report struct {
	JenisKejahatan string    `json:"jenisKejahatan"`
	Uraian         string    `json:"uraian"`
	EmailUser      string    `json:"user"`
	EmailPolisi    string    `json:"polisi"`
	Status         string    `json:"status"`
	ReportedAt     time.Time `json:"reportedAt"`
}