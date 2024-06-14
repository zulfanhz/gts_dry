package repository

import (
	"database/sql"
	"errors"
	"gts-dry/model"
	"gts-dry/util"
	"time"
)

type IncomingRepository interface {
	GetIncomingAll() ([]model.IncomingModel, error)
	GetIncomingByKode(kode string) (*model.IncomingModel, error)
	GetIncomingBySJ(noSJ string) ([]model.IncomingModel, error)
	GetIncomingByPO(po string) ([]model.IncomingModel, error)
	GetIncomingByPOdanProduct(po, codeProduct string) ([]model.IncomingModel, error)
	GetIncomingByPOdanProductSUMQTY(po, codeProduct string) (*float64, *float64, error)
	AddIncoming(req model.IncomingModel, id string) (model.IncomingModel, error)
	UpdateIncoming(req model.IncomingModel, currentIncoming model.IncomingModel, id string, rak model.RakModel, brg model.BarangModel) (model.IncomingModel, error)
	DeleteIncoming(currentIncoming model.IncomingModel, id string) error
}

type incomingRepository struct {
	db *sql.DB
}

func NewIncomingRepository(db *sql.DB) IncomingRepository {
	return &incomingRepository{db: db}
}

func (r *incomingRepository) GetIncomingAll() ([]model.IncomingModel, error) {
	var ins []model.IncomingModel

	rows, err := r.db.Query(`SELECT a.KodeIncoming,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.NoPO,a.QtyPO,a.NoSJ,a.QtySJ,a.QtyOk,a.QtyBad,a.GAP,a.ExpiredDate,a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakIncoming,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_incoming a
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var in model.IncomingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&in.KodeIncoming, &in.Tanggal, &in.ProductCode, &in.ProductName, &in.ProductCategory, &in.NoPO, &in.QtyPO, &in.NoSJ, &in.QtySJ, &in.QtyOK, &in.QtyBad, &in.GAP, &in.ExpDate, &in.RakCode, &in.JenisRak, &in.JenisRakIncoming, &in.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			in.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			in.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			in.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeDelete = &parsedTime
		}

		ins = append(ins, in)

	}

	return ins, nil
}

func (r *incomingRepository) GetIncomingByKode(kode string) (*model.IncomingModel, error) {
	var in model.IncomingModel
	var dateTimeEntry sql.NullString
	var dateTimeUpdated sql.NullString
	var dateTimeDeleted sql.NullString
	var userEntry sql.NullString
	var userUpdate sql.NullString
	var userDelete sql.NullString

	rows := r.db.QueryRow(`SELECT a.KodeIncoming,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.NoPO,a.QtyPO,a.NoSJ,a.QtySJ,a.QtyOk,a.QtyBad,a.GAP,a.ExpiredDate,a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakIncoming,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_incoming a
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null and a.KodeIncoming=
	?`, kode)

	err := rows.Scan(&in.KodeIncoming, &in.Tanggal, &in.ProductCode, &in.ProductName, &in.ProductCategory, &in.NoPO, &in.QtyPO, &in.NoSJ, &in.QtySJ, &in.QtyOK, &in.QtyBad, &in.GAP, &in.ExpDate, &in.RakCode, &in.JenisRak, &in.JenisRakIncoming, &in.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("incoming tidak ditemukan")
		}
		return nil, err
	}

	if userEntry.Valid {
		in.UserEntry = userEntry.String
	}
	if dateTimeEntry.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
		if err != nil {
			return nil, err
		}
		in.DateTimeEntry = &parsedTime
	}
	if userUpdate.Valid {
		in.UserUpdate = userUpdate.String
	}
	if dateTimeUpdated.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
		if err != nil {
			return nil, err
		}
		in.DateTimeUpdate = &parsedTime
	}
	if userDelete.Valid {
		in.UserDelete = userDelete.String
	}
	if dateTimeDeleted.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
		if err != nil {
			return nil, err
		}
		in.DateTimeDelete = &parsedTime
	}

	return &in, nil
}

func (r *incomingRepository) GetIncomingBySJ(noSJ string) ([]model.IncomingModel, error) {
	var ins []model.IncomingModel

	rows, err := r.db.Query(`SELECT a.KodeIncoming,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.NoPO,a.QtyPO,a.NoSJ,a.QtySJ,a.QtyOk,a.QtyBad,a.GAP,a.ExpiredDate,a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakIncoming,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_incoming a
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null and a.NoSJ=?
	`, noSJ)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var in model.IncomingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&in.KodeIncoming, &in.Tanggal, &in.ProductCode, &in.ProductName, &in.ProductCategory, &in.NoPO, &in.QtyPO, &in.NoSJ, &in.QtySJ, &in.QtyOK, &in.QtyBad, &in.GAP, &in.ExpDate, &in.RakCode, &in.JenisRak, &in.JenisRakIncoming, &in.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			in.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			in.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			in.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeDelete = &parsedTime
		}

		ins = append(ins, in)

	}

	return ins, nil
}

func (r *incomingRepository) GetIncomingByPO(po string) ([]model.IncomingModel, error) {
	var ins []model.IncomingModel

	rows, err := r.db.Query(`SELECT a.KodeIncoming,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.NoPO,a.QtyPO,a.NoSJ,a.QtySJ,a.QtyOk,a.QtyBad,a.GAP,a.ExpiredDate,a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakIncoming,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_incoming a
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null and a.NoPO=?
	`, po)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var in model.IncomingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&in.KodeIncoming, &in.Tanggal, &in.ProductCode, &in.ProductName, &in.ProductCategory, &in.NoPO, &in.QtyPO, &in.NoSJ, &in.QtySJ, &in.QtyOK, &in.QtyBad, &in.GAP, &in.ExpDate, &in.RakCode, &in.JenisRak, &in.JenisRakIncoming, &in.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			in.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			in.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			in.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeDelete = &parsedTime
		}

		ins = append(ins, in)

	}

	return ins, nil
}

func (r *incomingRepository) GetIncomingByPOdanProduct(po, codeProduct string) ([]model.IncomingModel, error) {
	var ins []model.IncomingModel

	rows, err := r.db.Query(`SELECT a.KodeIncoming,a.Tanggal,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, a.NoPO,a.QtyPO,a.NoSJ,a.QtySJ,a.QtyOk,a.QtyBad,a.GAP,a.ExpiredDate,a.RakCode,COALESCE(c.JenisRak,'') as JenisRak,a.JenisRakIncoming,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_incoming a
	LEFT JOIN mst_dry b on a.ProductCode=b.Code and b.DateTimeDelete is NULL
	LEFT JOIN  mst_rak c on a.RakCode = c.Code and c.DateTimeDelete is NULL
	where a.DateTimeDelete is null and a.NoPO=? and a.ProductCode=?
	`, po, codeProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var in model.IncomingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&in.KodeIncoming, &in.Tanggal, &in.ProductCode, &in.ProductName, &in.ProductCategory, &in.NoPO, &in.QtyPO, &in.NoSJ, &in.QtySJ, &in.QtyOK, &in.QtyBad, &in.GAP, &in.ExpDate, &in.RakCode, &in.JenisRak, &in.JenisRakIncoming, &in.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			in.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			in.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			in.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			in.DateTimeDelete = &parsedTime
		}

		ins = append(ins, in)

	}

	return ins, nil
}
func (r *incomingRepository) GetIncomingByPOdanProductSUMQTY(po, codeProduct string) (*float64, *float64, error) {
	var qtySum, qtyPO float64

	rows := r.db.QueryRow("SELECT COALESCE(SUM(QtyOK), 0) as QtySum, COALESCE(MIN(QtyPO), 0) as QtyPO FROM tr_incoming where DateTimeDelete is null and NoPO=? and ProductCode=?", po, codeProduct)

	err := rows.Scan(&qtySum, &qtyPO)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, errors.New("data tidak ditemukan")
		}
		return nil, nil, err
	}

	return &qtySum, &qtyPO, nil
}

func (r *incomingRepository) AddIncoming(req model.IncomingModel, id string) (model.IncomingModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.IncomingModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateinsert := time.Now()
	formattedDate := dateinsert.Format("2006-01-02 15:04:05")

	insertQuery := `
		INSERT INTO tr_incoming (Tanggal,ProductCode,NoPO,QtyPO,NoSJ,QtySJ,QtyOk,QtyBad,GAP,ExpiredDate,RakCode,JenisRakIncoming,Satuan,UserEntry,DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.ProductCode, req.NoPO, req.QtyPO, req.NoSJ, req.QtySJ, req.QtyOK, req.QtyBad, req.GAP, req.ExpDate, req.RakCode, req.JenisRakIncoming, req.Satuan, id, formattedDate)
	if err != nil {
		return model.IncomingModel{}, err
	}

	var existingQty float64
	checkQuery := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQuery, req.RakCode, req.ProductCode, req.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.IncomingModel{}, err
	}

	if err == sql.ErrNoRows {
		insertToRakIsi := `
			INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
			VALUES (?, ?, ?, ?)
		`
		_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.QtyOK, req.ExpDate)
		if err != nil {
			return model.IncomingModel{}, err
		}
	} else {
		newQty := existingQty + req.QtyOK
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, req.RakCode, req.ProductCode, req.ExpDate)
		if err != nil {
			return model.IncomingModel{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return model.IncomingModel{}, err
	}

	return req, nil
}

func (r *incomingRepository) UpdateIncoming(req model.IncomingModel, currentIncoming model.IncomingModel, id string, rak model.RakModel, brg model.BarangModel) (model.IncomingModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.IncomingModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertQuery := `
		UPDATE tr_incoming set Tanggal=?,ProductCode=?,NoPO=?,QtyPO=?,NoSJ=?,QtySJ=?,QtyOk=?,QtyBad=?,GAP=?,ExpiredDate=?,RakCode=?,JenisRakIncoming=?,Satuan=?,UserUpdate=?,DateTimeUpdate=? 
		WHERE KodeIncoming=? and DateTimeDelete is null
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.ProductCode, req.NoPO, req.QtyPO, req.NoSJ, req.QtySJ, req.QtyOK, req.QtyBad, req.GAP, req.ExpDate, req.RakCode, req.JenisRakIncoming, req.Satuan, id, formattedDate, currentIncoming.KodeIncoming)
	if err != nil {
		return model.IncomingModel{}, err
	}

	var existingQty float64
	checkQuery := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQuery, currentIncoming.RakCode, currentIncoming.ProductCode, currentIncoming.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.IncomingModel{}, err
	}

	if err == sql.ErrNoRows {
		if brg.Kategori == "FOOD" && rak.JenisRak == "STOK-EQUIPMENT" {
			tx.Rollback()
			return model.IncomingModel{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
		}

		if brg.Kategori == "EQUIPMENT" && rak.JenisRak == "STOK-FOOD" {
			tx.Rollback()
			return model.IncomingModel{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
		}

		if rak.JenisRak == "STOK-FOOD" {
			if req.ExpDate == "" {
				tx.Rollback()
				return model.IncomingModel{}, errors.New("barang food mesti memiliki exp date")
			}

			raks, err := r.GetRakIsiExcept(req.RakCode, req.ProductCode, req.ExpDate)
			if err != nil {
				return model.IncomingModel{}, err
			}

			if len(raks) > 0 {
				tx.Rollback()
				return model.IncomingModel{}, errors.New("rak tujuan memiliki barang dan exp yang berbeda")
			}
		}

		insertToRakIsi := `
			INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
			VALUES (?, ?, ?, ?)
		`
		_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.QtyOK, req.ExpDate)
		if err != nil {
			return model.IncomingModel{}, err
		}
	} else {
		newQty := existingQty - currentIncoming.QtyOK
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, currentIncoming.RakCode, currentIncoming.ProductCode, currentIncoming.ExpDate)
		if err != nil {
			return model.IncomingModel{}, err
		}

		if req.RakCode != currentIncoming.RakCode || req.ProductCode != currentIncoming.ProductCode || req.ExpDate != currentIncoming.ExpDate {

			qtySum, qtyPO, err := r.GetIncomingByPOdanProductSUMQTY(req.NoPO, req.ProductCode)
			if err != nil {
				return model.IncomingModel{}, err
			}

			*qtySum += req.QtyOK

			if *qtyPO == 0 {
				*qtyPO += req.QtyPO
			} else {
				if *qtyPO != req.QtyPO {
					tx.Rollback()
					return model.IncomingModel{}, errors.New("qty PO pada barang ini berbeda dari qty po sebelumnya, mohon dicek kembali")
				}
			}

			if *qtySum > 0 && *qtyPO > 0 {
				if *qtySum > *qtyPO {
					tx.Rollback()
					return model.IncomingModel{}, errors.New("qty sudah melebihi qty pada PO")
				}
			}

			if brg.Kategori == "FOOD" && rak.JenisRak == "STOK-EQUIPMENT" {
				tx.Rollback()
				return model.IncomingModel{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
			}

			if brg.Kategori == "EQUIPMENT" && rak.JenisRak == "STOK-FOOD" {
				tx.Rollback()
				return model.IncomingModel{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
			}

			if rak.JenisRak == "STOK-FOOD" {
				if req.ExpDate == "" {
					tx.Rollback()
					return model.IncomingModel{}, errors.New("barang food mesti memiliki exp date")
				}

				if newQty != 0 {
					raks, err := r.GetRakIsiExcept(req.RakCode, req.ProductCode, req.ExpDate)
					if err != nil {
						return model.IncomingModel{}, err
					}

					if len(raks) > 0 {
						tx.Rollback()
						return model.IncomingModel{}, errors.New("rak tujuan memiliki barang dan exp yang berbeda")
					}
				}

			}

			var existingQty float64
			checkQuery := `SELECT Qty 
							FROM mst_rak_isi 
							WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
			err = tx.QueryRow(checkQuery, req.RakCode, req.ProductCode, req.ExpDate).Scan(&existingQty)
			if err != nil && err != sql.ErrNoRows {
				return model.IncomingModel{}, err
			}

			if err == sql.ErrNoRows {
				insertToRakIsi := `
				INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
				VALUES (?, ?, ?, ?)`
				_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.QtyOK, req.ExpDate)
				if err != nil {
					return model.IncomingModel{}, err
				}
			} else {
				newQtye := existingQty + req.QtyOK
				updateRakIsi = `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
				_, err = tx.Exec(updateRakIsi, newQtye, req.RakCode, req.ProductCode, req.ExpDate)
				if err != nil {
					return model.IncomingModel{}, err
				}
			}

		} else {
			qtySum, qtyPO, err := r.GetIncomingByPOdanProductSUMQTY(currentIncoming.NoPO, currentIncoming.ProductCode)
			if err != nil {
				return model.IncomingModel{}, err
			}

			*qtySum = *qtySum - currentIncoming.QtyOK + req.QtyOK

			if *qtyPO == 0 {
				*qtyPO += req.QtyPO
			} else {
				if *qtyPO != req.QtyPO {
					tx.Rollback()
					return model.IncomingModel{}, errors.New("qty PO pada barang ini berbeda dari qty po sebelumnya, mohon dicek kembali")
				}
			}

			if *qtySum > 0 && *qtyPO > 0 {
				if *qtySum > *qtyPO {
					tx.Rollback()
					return model.IncomingModel{}, errors.New("qty sudah melebihi qty pada PO")
				}
			}

			newQty = newQty + req.QtyOK
			updateRakIsi = `
				UPDATE mst_rak_isi 
				SET Qty = ? 
				WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
			`
			_, err = tx.Exec(updateRakIsi, newQty, currentIncoming.RakCode, currentIncoming.ProductCode, currentIncoming.ExpDate)
			if err != nil {
				return model.IncomingModel{}, err
			}
		}
	}

	deleteNol := "DELETE from mst_rak_isi where Qty=0"
	_, err = tx.Exec(deleteNol)
	if err != nil {
		return model.IncomingModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.IncomingModel{}, err
	}

	return req, nil
}

func (r *incomingRepository) DeleteIncoming(currentIncoming model.IncomingModel, id string) error {
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
		UPDATE tr_incoming set UserDelete=?,DateTimeDelete=? 
		WHERE KodeIncoming=? and DateTimeDelete is null
	`
	_, err = tx.Exec(insertQuery, id, formattedDate, currentIncoming.KodeIncoming)
	if err != nil {
		return err
	}

	var existingQty float64
	checkQuery := `SELECT Qty 
					FROM mst_rak_isi 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
	err = tx.QueryRow(checkQuery, currentIncoming.RakCode, currentIncoming.ProductCode, currentIncoming.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		return errors.New("barang yg akan dihapus tidak ada di rak")
	} else {
		newQtye := existingQty - currentIncoming.QtyOK
		updateRakIsi := `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
		_, err = tx.Exec(updateRakIsi, newQtye, currentIncoming.RakCode, currentIncoming.ProductCode, currentIncoming.ExpDate)
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

func (r *incomingRepository) GetRakIsiExcept(code, productCode, exp string) ([]model.RakIsiModel, error) {
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
