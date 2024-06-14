package model

import "time"

type RakModel struct {
	Type           string     `json:"type"`
	Code           string     `json:"code"`
	JenisRak       string     `json:"jenis_rak"`
	UserEntry      string     `json:"user_entry" db:"UserEntry"`
	DateTimeEntry  *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate     string     `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete     string     `json:"user_delete" db:"UserDelete"`
	DateTimeDelete *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}

type RakModelWithoutUser struct {
	Type     string `json:"type"`
	Code     string `json:"code"`
	JenisRak string `json:"jenis_rak"`
}

type RakModelResponse struct {
	Rak    RakModel      `json:"rak"`
	RakIsi []RakIsiModel `json:"rak_isi"`
}

type RakIsiModel struct {
	RakCode         string  `json:"rak_code"`
	JenisRak        string  `json:"jenis_rak"`
	ProductCode     string  `json:"product_code"`
	ProductName     string  `json:"product_name"`
	ProductCategory string  `json:"product_category"`
	Qty             float64 `json:"qty"`
	ExpDate         string  `json:"exp_date" db:"ExpiredDate"`
	TimeRemaining   int     `json:"time_remaining"`
	Status          string  `json:"status"`
}
