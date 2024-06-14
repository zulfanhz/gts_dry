package model

import "time"

type MutasiRak struct {
	KodeMutasiRak   int        `json:"kode" db:"KodeMutasiRak"`
	Tanggal         string     `json:"tanggal"`
	ProductCode     string     `json:"product_code"`
	ProductName     string     `json:"product_name"`
	ProductCategory string     `json:"product_category"`
	RakCodeAsal     string     `json:"rak_code_asal"`
	JenisRakAsal    string     `json:"jenis_rak_asal"`
	RakCodeTujuan   string     `json:"rak_code_tujuan"`
	JenisRakTujuan  string     `json:"jenis_rak_tujuan"`
	QtyMutasi       float64    `json:"qty_mutasi"`
	Satuan          string     `json:"satuan"`
	ExpiredDate     string     `json:"expired_date"`
	UserEntry       string     `json:"user_entry" db:"UserEntry"`
	DateTimeEntry   *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate      string     `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate  *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete      string     `json:"user_delete" db:"UserDelete"`
	DateTimeDelete  *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}

type MutasiRakRequest struct {
	Tanggal       string  `json:"tanggal"`
	ProductCode   string  `json:"product_code"`
	RakCodeAsal   string  `json:"rak_code_asal"`
	RakCodeTujuan string  `json:"rak_code_tujuan"`
	QtyMutasi     float64 `json:"qty_mutasi"`
	Satuan        string  `json:"satuan"`
	ExpiredDate   string  `json:"expired_date"`
}
