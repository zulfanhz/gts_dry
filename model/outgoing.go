package model

import "time"

type OutgoingModel struct {
	KodeOutgoing     string     `json:"kode_outgoing" db:"KodeOutgoing"`
	Tanggal          string     `json:"tanggal" db:"Tanggal"`
	NoSJ             string     `json:"no_sj" db:"NoSJ"`
	RestoCode        string     `json:"resto_code" db:"RestoCode"`
	RestoName        string     `json:"resto_name"`
	RestoCategory    string     `json:"resto_category"`
	ProductCode      string     `json:"product_code" db:"ProductCode"`
	ProductName      string     `json:"product_name"`
	ProductCategory  string     `json:"product_category"`
	ExpDate          string     `json:"exp_date" db:"ExpiredDate"`
	RakCode          string     `json:"rak_code" db:"RakCode"`
	JenisRakOutgoing string     `json:"jenis_rak_outgoing" db:"JenisRakOutgoing"`
	QtySJ            float64    `json:"qty_sj" db:"QtySJ"`
	QtyOut           float64    `json:"qty_out" db:"QtyOut"`
	Satuan           string     `json:"satuan" db:"Satuan"`
	UserEntry        string     `json:"user_entry" db:"UserEntry"`
	DateTimeEntry    *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate       string     `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate   *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete       string     `json:"user_delete" db:"UserDelete"`
	DateTimeDelete   *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}
