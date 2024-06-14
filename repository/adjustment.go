package repository

import (
	"database/sql"
	"errors"
	"gts-dry/model"
	"gts-dry/util"
	"time"
)

type AdjustmentRepository interface {
	GetAdjustmentAll() ([]model.AdjustmentModel, error)
	GetAdjustmentByKode(kode string) (*model.AdjustmentModel, error)
	GetAdjustmentByProduct(kode string) ([]model.AdjustmentModel, error)
	AddAdjustment(req model.AdjustmentRequestModel, id string) (model.AdjustmentRequestModel, error)
	UpdateAdjustment(req model.AdjustmentRequestModel, currentAdjust model.AdjustmentModel, id string, rak model.RakModel, brg model.BarangModel) (model.AdjustmentRequestModel, error)
	DeleteAdjustment(currentAdjust model.AdjustmentModel, id string) error
}

type adjustmentRepository struct {
	db *sql.DB
}

func NewAdjustmentRepository(db *sql.DB) AdjustmentRepository {
	return &adjustmentRepository{db: db}
}

func (r *adjustmentRepository) GetAdjustmentAll() ([]model.AdjustmentModel, error) {
	var adjusments []model.AdjustmentModel

	rows, err := r.db.Query(`SELECT a.KodeAdjustment,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.RakCode,COALESCE(c.JenisRak,'') as JenisRak, a.JenisRakAdjustment, a.Qty,a.Satuan,a.ExpiredDate,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_adjustment a 
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var adjus model.AdjustmentModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&adjus.KodeAdjustment, &adjus.Tanggal, &adjus.ProductCode, &adjus.ProductName, &adjus.ProductCategory, &adjus.RakCode, &adjus.JenisRak, &adjus.JenisRakAdjustment, &adjus.Qty, &adjus.Satuan, &adjus.ExpDate, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			adjus.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			adjus.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			adjus.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			adjus.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			adjus.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			adjus.DateTimeDelete = &parsedTime
		}

		adjusments = append(adjusments, adjus)

	}

	return adjusments, nil
}

func (r *adjustmentRepository) GetAdjustmentByKode(kode string) (*model.AdjustmentModel, error) {
	var adjus model.AdjustmentModel
	var dateTimeEntry sql.NullString
	var dateTimeUpdated sql.NullString
	var dateTimeDeleted sql.NullString
	var userEntry sql.NullString
	var userUpdate sql.NullString
	var userDelete sql.NullString

	row := r.db.QueryRow(`SELECT a.KodeAdjustment,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakAdjustment, a.Qty,a.Satuan,a.ExpiredDate,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_adjustment a 
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null and a.KodeAdjustment=?`, kode)
	err := row.Scan(&adjus.KodeAdjustment, &adjus.Tanggal, &adjus.ProductCode, &adjus.ProductName, &adjus.ProductCategory, &adjus.RakCode, &adjus.JenisRak, &adjus.JenisRakAdjustment, &adjus.Qty, &adjus.Satuan, &adjus.ExpDate, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("adjustment tidak ditemukan")
		}
		return nil, err
	}

	if userEntry.Valid {
		adjus.UserEntry = userEntry.String
	}
	if dateTimeEntry.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
		if err != nil {
			return nil, err
		}
		adjus.DateTimeEntry = &parsedTime
	}
	if userUpdate.Valid {
		adjus.UserUpdate = userUpdate.String
	}
	if dateTimeUpdated.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
		if err != nil {
			return nil, err
		}
		adjus.DateTimeUpdate = &parsedTime
	}
	if userDelete.Valid {
		adjus.UserDelete = userDelete.String
	}
	if dateTimeDeleted.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
		if err != nil {
			return nil, err
		}
		adjus.DateTimeDelete = &parsedTime
	}

	return &adjus, nil
}

func (r *adjustmentRepository) GetAdjustmentByProduct(kode string) ([]model.AdjustmentModel, error) {
	var adjusments []model.AdjustmentModel

	rows, err := r.db.Query(`SELECT a.KodeAdjustment,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, b.Kategori as ProductCategory, a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakAdjustment, a.Qty,a.Satuan,a.ExpiredDate,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_adjustment a 
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null and a.ProductCode=?`, kode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var adjus model.AdjustmentModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&adjus.KodeAdjustment, &adjus.Tanggal, &adjus.ProductCode, &adjus.ProductName, &adjus.ProductCategory, &adjus.RakCode, &adjus.JenisRak, &adjus.JenisRakAdjustment, &adjus.Qty, &adjus.Satuan, &adjus.ExpDate, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			adjus.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			adjus.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			adjus.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			adjus.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			adjus.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			adjus.DateTimeDelete = &parsedTime
		}

		adjusments = append(adjusments, adjus)

	}

	return adjusments, nil
}

func (r *adjustmentRepository) AddAdjustment(req model.AdjustmentRequestModel, id string) (model.AdjustmentRequestModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateinsert := time.Now()
	formattedDate := dateinsert.Format("2006-01-02 15:04:05")

	insertQuery := `
		INSERT INTO tr_adjustment (Tanggal,ProductCode,RakCode,JenisRakAdjustment,Qty,Satuan,ExpiredDate,UserEntry,DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.ProductCode, req.RakCode, req.JenisRakAdjustment, req.Qty, req.Satuan, req.ExpDate, id, formattedDate)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	var existingQty float64
	checkQuery := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQuery, req.RakCode, req.ProductCode, req.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.AdjustmentRequestModel{}, err
	}

	if err == sql.ErrNoRows {
		insertToRakIsi := `
			INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
			VALUES (?, ?, ?, ?)
		`
		_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.Qty, req.ExpDate)
		if err != nil {
			return model.AdjustmentRequestModel{}, err
		}
	} else {
		newQty := existingQty + req.Qty
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, req.RakCode, req.ProductCode, req.ExpDate)
		if err != nil {
			return model.AdjustmentRequestModel{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	return req, nil
}

func (r *adjustmentRepository) UpdateAdjustment(req model.AdjustmentRequestModel, currentAdjust model.AdjustmentModel, id string, rak model.RakModel, brg model.BarangModel) (model.AdjustmentRequestModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertQuery := `
		UPDATE tr_adjustment set Tanggal=?,ProductCode=?,RakCode=?,JenisRakAdjustment=?,Qty=?,Satuan=?,ExpiredDate=?,UserUpdate=?,DateTimeUpdate=? 
		WHERE KodeAdjustment=? and DateTimeDelete is null
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.ProductCode, req.RakCode, req.JenisRakAdjustment, req.Qty, req.Satuan, req.ExpDate, id, formattedDate, currentAdjust.KodeAdjustment)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	var existingQty float64
	checkQuery := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQuery, currentAdjust.RakCode, currentAdjust.ProductCode, currentAdjust.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.AdjustmentRequestModel{}, err
	}

	if err == sql.ErrNoRows {

		if brg.Kategori == "FOOD" && rak.JenisRak == "STOK-EQUIPMENT" {
			tx.Rollback()
			return model.AdjustmentRequestModel{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
		}

		if brg.Kategori == "EQUIPMENT" && rak.JenisRak == "STOK-FOOD" {
			tx.Rollback()
			return model.AdjustmentRequestModel{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
		}

		if rak.JenisRak == "STOK-FOOD" {
			if req.ExpDate == "" {
				tx.Rollback()
				return model.AdjustmentRequestModel{}, errors.New("barang food mesti memiliki exp date")
			}

			raks, err := r.GetRakIsiExcept(req.RakCode, req.ProductCode, req.ExpDate)
			if err != nil {
				return model.AdjustmentRequestModel{}, err
			}

			if len(raks) > 0 {
				tx.Rollback()
				return model.AdjustmentRequestModel{}, errors.New("rak tujuan memiliki barang dan exp yang berbeda")
			}
		}

		insertToRakIsi := `
			INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
			VALUES (?, ?, ?, ?)
		`
		_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.Qty, req.ExpDate)
		if err != nil {
			return model.AdjustmentRequestModel{}, err
		}
	} else {
		newQty := existingQty - currentAdjust.Qty
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, currentAdjust.RakCode, currentAdjust.ProductCode, currentAdjust.ExpDate)
		if err != nil {
			return model.AdjustmentRequestModel{}, err
		}

		if req.RakCode != currentAdjust.RakCode || req.ProductCode != currentAdjust.ProductCode || req.ExpDate != currentAdjust.ExpDate {

			if brg.Kategori == "FOOD" && rak.JenisRak == "STOK-EQUIPMENT" {
				tx.Rollback()
				return model.AdjustmentRequestModel{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
			}

			if brg.Kategori == "EQUIPMENT" && rak.JenisRak == "STOK-FOOD" {
				tx.Rollback()
				return model.AdjustmentRequestModel{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
			}

			if rak.JenisRak == "STOK-FOOD" {
				if req.ExpDate == "" {
					tx.Rollback()
					return model.AdjustmentRequestModel{}, errors.New("barang food mesti memiliki exp date")
				}

				if newQty != 0 {
					raks, err := r.GetRakIsiExcept(req.RakCode, req.ProductCode, req.ExpDate)
					if err != nil {
						return model.AdjustmentRequestModel{}, err
					}

					if len(raks) > 0 {
						tx.Rollback()
						return model.AdjustmentRequestModel{}, errors.New("rak tujuan memiliki barang dan exp yang berbeda")
					}
				}

			}

			var existingQty float64
			checkQuery := `SELECT Qty 
							FROM mst_rak_isi 
							WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
			err = tx.QueryRow(checkQuery, req.RakCode, req.ProductCode, req.ExpDate).Scan(&existingQty)
			if err != nil && err != sql.ErrNoRows {
				return model.AdjustmentRequestModel{}, err
			}

			if err == sql.ErrNoRows {
				insertToRakIsi := `
				INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
				VALUES (?, ?, ?, ?)`
				_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.Qty, req.ExpDate)
				if err != nil {
					return model.AdjustmentRequestModel{}, err
				}
			} else {
				newQtye := existingQty + req.Qty
				updateRakIsi = `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
				_, err = tx.Exec(updateRakIsi, newQtye, req.RakCode, req.ProductCode, req.ExpDate)
				if err != nil {
					return model.AdjustmentRequestModel{}, err
				}
			}

		} else {
			newQty = newQty + req.Qty
			updateRakIsi = `
				UPDATE mst_rak_isi 
				SET Qty = ? 
				WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
			`
			_, err = tx.Exec(updateRakIsi, newQty, currentAdjust.RakCode, currentAdjust.ProductCode, currentAdjust.ExpDate)
			if err != nil {
				return model.AdjustmentRequestModel{}, err
			}
		}
	}

	deleteNol := "DELETE from mst_rak_isi where Qty=0"
	_, err = tx.Exec(deleteNol)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	return req, nil
}

func (r *adjustmentRepository) DeleteAdjustment(currentAdjust model.AdjustmentModel, id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertQuery := `
		UPDATE tr_adjustment set UserDelete=?,DateTimeDelete=? 
		WHERE KodeAdjustment=? and DateTimeDelete is null
	`
	_, err = tx.Exec(insertQuery, id, formattedDate, currentAdjust.KodeAdjustment)
	if err != nil {
		return err
	}

	var existingQty float64
	checkQuery := `SELECT Qty 
					FROM mst_rak_isi 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
	err = tx.QueryRow(checkQuery, currentAdjust.RakCode, currentAdjust.ProductCode, currentAdjust.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		return errors.New("barang yg akan dihapus tidak ada di rak")
	} else {
		newQtye := existingQty - currentAdjust.Qty
		updateRakIsi := `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
		_, err = tx.Exec(updateRakIsi, newQtye, currentAdjust.RakCode, currentAdjust.ProductCode, currentAdjust.ExpDate)
		if err != nil {
			return err
		}
	}

	deleteNol := "DELETE from mst_rak_isi where Qty=0"
	_, err = tx.Exec(deleteNol)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *adjustmentRepository) GetRakIsiExcept(code, productCode, exp string) ([]model.RakIsiModel, error) {
	var raks []model.RakIsiModel

	rows, err := r.db.Query(`SELECT 
	ri.RakCode,ri.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,ri.Qty,ri.ExpiredDate 
	FROM mst_rak_isi ri 
	LEFT JOIN 
        mst_rak r ON ri.RakCode = r.Code and r.DateTimeDelete is null
    LEFT JOIN 
        mst_dry b ON ri.ProductCode = b.Code and b.DateTimeDelete is null
	where ri.RakCode=? and (ri.ProductCode<>? or ri.ExpiredDate<>?) order by ri.ExpiredDate asc`, code, productCode, exp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rak model.RakIsiModel

		if err := rows.Scan(&rak.RakCode, &rak.ProductCode, &rak.ProductName, &rak.ProductCategory, &rak.Qty, &rak.ExpDate); err != nil {
			return nil, err
		}

		if rak.ProductCategory == "FOOD" {
			tr, stts := util.CalculateRemainingExp(rak.ExpDate)
			rak.TimeRemaining = tr
			rak.Status = stts
		} else {
			rak.TimeRemaining = 0
			rak.Status = "TANPA EXPIRED DATE"
		}

		raks = append(raks, rak)
	}

	return raks, nil
}
