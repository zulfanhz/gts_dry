package model

import "time"

type IncomingModel struct {
	KodeIncoming     string     `json:"kode_incoming" db:"KodeIncoming"`
	Tanggal          string     `json:"tanggal" db:"Tanggal" validate:"required"`
	ProductCode      string     `json:"product_code" db:"ProductCode" validate:"required"`
	ProductName      string     `json:"product_name"`
	ProductCategory  string     `json:"product_category"`
	NoPO             string     `json:"no_po" db:"NoPO" validate:"required"`
	QtyPO            float64    `json:"qty_po" db:"QtyPO" validate:"required"`
	NoSJ             string     `json:"no_sj" db:"NoSJ" validate:"required"`
	QtySJ            float64    `json:"qty_sj" db:"QtySJ" validate:"required"`
	QtyOK            float64    `json:"qty_ok" db:"QtyOK"`
	QtyBad           float64    `json:"qty_bad" db:"QtyBad"`
	GAP              float64    `json:"gap" db:"GAP"`
	ExpDate          string     `json:"exp_date" db:"ExpiredDate"`
	RakCode          string     `json:"rak_code" db:"RakCode" validate:"required"`
	JenisRakIncoming string     `json:"jenis_rak_incoming" db:"JenisRakIncoming" validate:"required"`
	JenisRak         string     `json:"jenis_rak"`
	Satuan           string     `json:"satuan" db:"Satuan" validate:"required"`
	UserEntry        string     `json:"user_entry" db:"UserEntry"`
	DateTimeEntry    *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate       string     `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate   *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete       string     `json:"user_delete" db:"UserDelete"`
	DateTimeDelete   *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}
