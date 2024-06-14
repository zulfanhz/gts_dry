package model

import "time"

type BarangModel struct {
	Kode           string     `json:"code" db:"Code"`
	Nama           string     `json:"nama" db:"Name"`
	Barcode        string     `json:"barcode" db:"Barcode"`
	Satuan         *string    `json:"satuan"`
	Stock          float64    `json:"stock"`
	Kategori       string     `json:"kategori"`
	UserEntry      string     `json:"user_entry" db:"UserEntry"`
	DateTimeEntry  *time.Time `json:"date_time_entry" db:"DateTimeEntry"`
	UserUpdate     string     `json:"user_update" db:"UserUpdate"`
	DateTimeUpdate *time.Time `json:"date_time_update" db:"DateTimeUpdate"`
	UserDelete     string     `json:"user_delete" db:"UserDelete"`
	DateTimeDelete *time.Time `json:"date_time_delete" db:"DateTimeDelete"`
}

type BarangModelWithoutUser struct {
	Kode     string              `json:"code" db:"Code"`
	Nama     string              `json:"nama" db:"Name"`
	Barcode  string              `json:"barcode" db:"Barcode"`
	Satuan   []BarangSatuanModel `json:"satuan"`
	Stock    float64             `json:"stock"`
	Kategori string              `json:"kategori"`
}

type BarangSatuanModel struct {
	Kode         string  `json:"code" db:"Code"`
	Satuan       string  `json:"satuan" db:"Satuan"`
	Qty          float64 `json:"qty" db:"Qty"`
	UrutanSatuan int     `json:"urutan_satuan" db:"Level"`
	SatuanUtama  int     `json:"satuan_utama" db:"IsHitung"`
	Stock        float64 `json:"stock"`
	Kategori     string  `json:"kategori"`
}

type BarangResponseModel struct {
	Barang BarangModel         `json:"barang"`
	Satuan []BarangSatuanModel `json:"satuan"`
}

type BarangWithoutUser struct {
	Barang BarangModelWithoutUser `json:"barang"`
	Satuan []BarangSatuanModel    `json:"satuan"`
}

type BarangModelStok struct {
	Kode     string  `json:"code" db:"Code"`
	Nama     string  `json:"nama" db:"Name"`
	Barcode  string  `json:"barcode" db:"Barcode"`
	Stok     float64 `json:"stok"`
	Kategori string  `json:"kategori"`
}
