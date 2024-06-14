package model

import "time"

type AdjustmentModel struct {
	KodeAdjustment     string    `json:"kode_adjustment" db:"KodeAdjustment"`
	Tanggal            string    `json:"tanggal" db:"Tanggal"`
	ProductCode        string    `json:"product_code" db:"ProductCode"`
	ProductName        string    `json:"product_name"`
	ProductCategory    string    `json:"product_category"`
	JenisRakAdjustment string    `json:"jenis_rak_adjustment"`
	RakCode            string    `json:"rak_code" db:"RakCode"`
	JenisRak           string    `json:"jenis_rak"`
	Qty                float64   `json:"qty" db:"qty"`
	Satuan             string    `json:"satuan" db:"satuan"`
	ExpDate            string    `json:"exp_date" db:"ExpiredDate"`
	UserEntry          string    `json:"user_entry" db:"UserEntry"`
	DateTimeEntry      *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate         string    `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate     *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete         string    `json:"user_delete" db:"UserDelete"`
	DateTimeDelete     *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}

type AdjustmentRequestModel struct {
	Tanggal            string  `json:"tanggal" db:"Tanggal"`
	ProductCode        string  `json:"product_code" db:"ProductCode"`
	RakCode            string  `json:"rak_code" db:"RakCode"`
	JenisRakAdjustment string  `json:"jenis_rak_adjustment" db:"JenisRakAdjustment"`
	Qty                float64 `json:"qty" db:"qty"`
	Satuan             string  `json:"satuan" db:"satuan"`
	ExpDate            string  `json:"exp_date" db:"ExpiredDate"`
}
