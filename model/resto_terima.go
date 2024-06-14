package model

import "time"

type RestoTerimaModel struct {
	KodeRestoTerima string    `json:"kode_resto_terima" db:"KodeRestoTerima" validate:"required"`
	RestoCode       string    `json:"resto_code" db:"RestoCode" validate:"required"`
	RestoName       string    `json:"resto_name"`
	RestoCategory   string    `json:"resto_category"`
	Tanggal         string    `json:"tanggal" db:"Tanggal" validate:"required"`
	NoSJ            string    `json:"no_sj" db:"NoSJ" validate:"required"`
	ProductCode     string    `json:"product_code" db:"ProductCode" validate:"required"`
	ProductName     string    `json:"product_name"`
	ProductCategory string    `json:"product_category"`
	QtySJ           float64   `json:"qty_sj" db:"QtySJ" validate:"required"`
	QtyTerima       float64   `json:"qty_terima" db:"QtyTerima" validate:"required"`
	ExpDate         string    `json:"exp_date" db:"ExpiredDate"`
	Satuan          string    `json:"satuan" db:"Satuan" validate:"required"`
	UserEntry       string    `json:"user_entry" db:"UserEntry"`
	DateTimeEntry   *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate      string    `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate  *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete      string    `json:"user_delete" db:"UserDelete"`
	DateTimeDelete  *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}
