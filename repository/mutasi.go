package repository

import (
	"database/sql"
	"gts-dry/model"
	"time"
)

type MutasiRepository interface {
	// ambil data mutasi all
	GetMutasiAll() ([]model.MutasiRak, error)
	// mutasi rak
	AddMutasi(req model.MutasiRak, id string) (model.MutasiRak, error)
	// update mutasi
	// delete mutasi
}

type mutasiRepository struct {
	db *sql.DB
}

func NewMutasiRepository(db *sql.DB) MutasiRepository {
	return &mutasiRepository{db: db}
}

func (r *mutasiRepository) GetMutasiAll() ([]model.MutasiRak, error) {
	var mutasi []model.MutasiRak
	rows, err := r.db.Query(`SELECT a.KodeMutasiRak,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.RakCodeAsal,a.JenisRakAsal,a.RakCodeTujuan,a.JenisRakTujuan,a.QtyMutasi,a.Satuan,a.ExpiredDate,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete FROM tr_mutasi_rak a
	LEFT JOIN mst_dry b on a.ProductCode = b.Code
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mts model.MutasiRak
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&mts.KodeMutasiRak, &mts.Tanggal, &mts.ProductCode, &mts.ProductName, &mts.ProductCategory, &mts.RakCodeAsal, &mts.JenisRakAsal, &mts.RakCodeTujuan, &mts.JenisRakTujuan, &mts.QtyMutasi, &mts.Satuan, &mts.ExpiredDate, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			mts.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			mts.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			mts.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			mts.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			mts.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			mts.DateTimeDelete = &parsedTime
		}

		mutasi = append(mutasi, mts)

	}

	return mutasi, nil
}

func (r *mutasiRepository) AddMutasi(req model.MutasiRak, id string) (model.MutasiRak, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.MutasiRak{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateinsert := time.Now()
	formattedDate := dateinsert.Format("2006-01-02 15:04:05")

	insertQuery := `
		INSERT INTO tr_mutasi_rak (Tanggal,ProductCode,RakCodeAsal,JenisRakAsal,RakCodeTujuan,JenisRakTujuan,QtyMutasi,Satuan,ExpiredDate,UserEntry,DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.ProductCode, req.RakCodeAsal, req.JenisRakAsal, req.RakCodeTujuan, req.JenisRakTujuan, req.QtyMutasi, req.Satuan, req.ExpiredDate, id, formattedDate)
	if err != nil {
		return model.MutasiRak{}, err
	}

	var qtyLama float64
	checkQueryLama := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQueryLama, req.RakCodeAsal, req.ProductCode, req.ExpiredDate).Scan(&qtyLama)
	if err != nil && err != sql.ErrNoRows {
		return model.MutasiRak{}, err
	}

	if err == sql.ErrNoRows {
		tx.Rollback()
		return model.MutasiRak{}, err
	} else {
		newQty := qtyLama - req.QtyMutasi
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, req.RakCodeAsal, req.ProductCode, req.ExpiredDate)
		if err != nil {
			return model.MutasiRak{}, err
		}
	}

	var existingQty float64
	checkQuery := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQuery, req.RakCodeTujuan, req.ProductCode, req.ExpiredDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.MutasiRak{}, err
	}

	if err == sql.ErrNoRows {
		insertToRakIsi := `
			INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
			VALUES (?, ?, ?, ?)
		`
		_, err = tx.Exec(insertToRakIsi, req.RakCodeTujuan, req.ProductCode, req.QtyMutasi, req.ExpiredDate)
		if err != nil {
			return model.MutasiRak{}, err
		}
	} else {
		newQty := existingQty + req.QtyMutasi
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, req.RakCodeTujuan, req.ProductCode, req.ExpiredDate)
		if err != nil {
			return model.MutasiRak{}, err
		}
	}

	deleteNol := "DELETE from mst_rak_isi where Qty=0"
	_, err = tx.Exec(deleteNol)
	if err != nil {
		return model.MutasiRak{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.MutasiRak{}, err
	}

	return req, nil
}
