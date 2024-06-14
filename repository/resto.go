package repository

import (
	"database/sql"
	"errors"
	"gts-dry/model"
)

type RestoRepository interface {
	GetRestoAll() ([]model.RestoModel, error)
	GetRestoByKategori(kategori string) ([]model.RestoModel, error)
	GetRestoByKode(kode string) (*model.RestoModel, error)
}

type restoRepository struct {
	db *sql.DB
}

func NewRestoRepository(db *sql.DB) RestoRepository {
	return &restoRepository{db: db}
}

func (r *restoRepository) GetRestoAll() ([]model.RestoModel, error) {
	var result []model.RestoModel

	rows, err := r.db.Query("SELECT Kode,Name,Kategori FROM mst_resto where Aktif=1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resto model.RestoModel

		if err := rows.Scan(&resto.Kode, &resto.Nama, &resto.Kategori); err != nil {
			return nil, err
		}

		result = append(result, resto)

	}

	return result, nil
}

func (r *restoRepository) GetRestoByKategori(kategori string) ([]model.RestoModel, error) {
	var result []model.RestoModel

	rows, err := r.db.Query("SELECT Kode,Name,Kategori FROM mst_resto where Aktif=1 and Kategori=?", kategori)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resto model.RestoModel

		if err := rows.Scan(&resto.Kode, &resto.Nama, &resto.Kategori); err != nil {
			return nil, err
		}

		result = append(result, resto)

	}

	return result, nil
}

func (r *restoRepository) GetRestoByKode(kode string) (*model.RestoModel, error) {
	rows := r.db.QueryRow("SELECT Kode,Name,Kategori FROM mst_resto where Aktif=1 and Kode=?", kode)
	res := model.RestoModel{}
	err := rows.Scan(&res.Kode, &res.Nama, &res.Kategori)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("resto tidak ditemukan")
		}
		return nil, err
	}

	return &res, nil
}
