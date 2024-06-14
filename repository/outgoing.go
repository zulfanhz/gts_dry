package repository

import (
	"database/sql"
	"errors"
	"gts-dry/model"
	"gts-dry/util"
	"time"
)

type OutgoingRepository interface {
	GetOutgoingAll() ([]model.OutgoingModel, error)
	GetOutgoingByKode(kode string) (*model.OutgoingModel, error)
	GetOutgoingBySJ(kode string) ([]model.OutgoingModel, error)
	GetOutgoingByProduct(kode string) ([]model.OutgoingModel, error)
	GetOutgoingByResto(kode string) ([]model.OutgoingModel, error)
	GetOutgoingByTanggal(tanggalAwal, tanggalAkhir string) ([]model.OutgoingModel, error)
	// GetOutgoingCheck(nosj, rakCode, productCode, expDate string) ([]model.OutgoingModel, error)
	AddOutgoing(req model.OutgoingModel, id string) (model.OutgoingModel, error)
	UpdateOutgoing(req model.OutgoingModel, currentOutgoing model.OutgoingModel, id string) (model.OutgoingModel, error)
	DeleteOutgoing(currentOutgoing model.OutgoingModel, id string) error
}

type outgoingRepository struct {
	db *sql.DB
}

func NewOutgoingRepository(db *sql.DB) OutgoingRepository {
	return &outgoingRepository{db: db}
}

// func (r *outgoingRepository) GetOutgoingCheck(nosj, rakCode, productCode, expDate string) ([]model.OutgoingModel, error) {
// 	var outs []model.OutgoingModel

// 	rows, err := r.db.Query("SELECT KodeOutgoing,Tanggal,RestoCode,ProductCode,RakCode,ExpiredDate,NoSJ,QtySJ,QtyOut,Satuan,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM tr_outgoing where DateTimeDelete is null and NoSJ=? and RakCode=? and ProductCode=? and ExpiredDate=?", nosj, rakCode, productCode, expDate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var out model.OutgoingModel
// 		var dateTimeEntry []byte
// 		var dateTimeUpdated []byte
// 		var dateTimeDeleted []byte
// 		var userEntry sql.NullString
// 		var userUpdate sql.NullString
// 		var userDelete sql.NullString

// 		if err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.ProductCode, &out.RakCode, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
// 			return nil, err
// 		}

// 		outs = append(outs, out)

// 	}

// 	return outs, nil
// }

func (r *outgoingRepository) GetOutgoingAll() ([]model.OutgoingModel, error) {
	var outs []model.OutgoingModel

	rows, err := r.db.Query(`
	SELECT a.KodeOutgoing,a.Tanggal,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.RakCode,a.JenisRakOutgoing,a.ExpiredDate,a.NoSJ,a.QtySJ,a.QtyOut,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_outgoing a 
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete is null`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var out model.OutgoingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.RestoName, &out.RestoCategory, &out.ProductCode, &out.ProductName, &out.ProductCategory, &out.RakCode, &out.JenisRakOutgoing, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			out.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			out.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			out.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeDelete = &parsedTime
		}

		outs = append(outs, out)

	}

	return outs, nil
}

func (r *outgoingRepository) GetOutgoingByKode(kode string) (*model.OutgoingModel, error) {
	var out model.OutgoingModel
	var dateTimeEntry sql.NullString
	var dateTimeUpdated sql.NullString
	var dateTimeDeleted sql.NullString
	var userEntry sql.NullString
	var userUpdate sql.NullString
	var userDelete sql.NullString

	rows := r.db.QueryRow(`SELECT a.KodeOutgoing,a.Tanggal,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.RakCode,a.JenisRakOutgoing,a.ExpiredDate,a.NoSJ,a.QtySJ,a.QtyOut,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_outgoing a 
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete is null and a.KodeOutgoing=?`, kode)

	err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.RestoName, &out.RestoCategory, &out.ProductCode, &out.ProductName, &out.ProductCategory, &out.RakCode, &out.JenisRakOutgoing, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("outgoing tidak ditemukan")
		}
		return nil, err
	}

	if userEntry.Valid {
		out.UserEntry = userEntry.String
	}
	if dateTimeEntry.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
		if err != nil {
			return nil, err
		}
		out.DateTimeEntry = &parsedTime
	}
	if userUpdate.Valid {
		out.UserUpdate = userUpdate.String
	}
	if dateTimeUpdated.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
		if err != nil {
			return nil, err
		}
		out.DateTimeUpdate = &parsedTime
	}
	if userDelete.Valid {
		out.UserDelete = userDelete.String
	}
	if dateTimeDeleted.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
		if err != nil {
			return nil, err
		}
		out.DateTimeDelete = &parsedTime
	}

	return &out, nil
}

func (r *outgoingRepository) GetOutgoingBySJ(kode string) ([]model.OutgoingModel, error) {
	var outs []model.OutgoingModel

	rows, err := r.db.Query(`SELECT a.KodeOutgoing,a.Tanggal,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.RakCode,a.JenisRakOutgoing,a.ExpiredDate,a.NoSJ,a.QtySJ,a.QtyOut,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_outgoing a 
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete is null and a.NoSJ=?`, kode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var out model.OutgoingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.RestoName, &out.RestoCategory, &out.ProductCode, &out.ProductName, &out.ProductCategory, &out.RakCode, &out.JenisRakOutgoing, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			out.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			out.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			out.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeDelete = &parsedTime
		}

		outs = append(outs, out)

	}

	return outs, nil
}

func (r *outgoingRepository) GetOutgoingByProduct(kode string) ([]model.OutgoingModel, error) {
	var outs []model.OutgoingModel

	rows, err := r.db.Query(`SELECT a.KodeOutgoing,a.Tanggal,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.RakCode,a.JenisRakOutgoing,a.ExpiredDate,a.NoSJ,a.QtySJ,a.QtyOut,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_outgoing a 
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete is null and a.ProductCode=?`, kode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var out model.OutgoingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.RestoName, &out.RestoCategory, &out.ProductCode, &out.ProductName, &out.ProductCategory, &out.RakCode, &out.JenisRakOutgoing, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			out.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			out.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			out.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeDelete = &parsedTime
		}

		outs = append(outs, out)

	}

	return outs, nil
}

func (r *outgoingRepository) GetOutgoingByResto(kode string) ([]model.OutgoingModel, error) {
	var outs []model.OutgoingModel

	rows, err := r.db.Query(`SELECT a.KodeOutgoing,a.Tanggal,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.RakCode,a.JenisRakOutgoing,a.ExpiredDate,a.NoSJ,a.QtySJ,a.QtyOut,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_outgoing a 
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete is null and a.RestoCode=?`, kode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var out model.OutgoingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.RestoName, &out.RestoCategory, &out.ProductCode, &out.ProductName, &out.ProductCategory, &out.RakCode, &out.JenisRakOutgoing, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			out.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			out.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			out.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeDelete = &parsedTime
		}

		outs = append(outs, out)

	}

	return outs, nil
}

func (r *outgoingRepository) GetOutgoingByTanggal(tanggalAwal, tanggalAkhir string) ([]model.OutgoingModel, error) {
	var outs []model.OutgoingModel

	rows, err := r.db.Query(`SELECT a.KodeOutgoing,a.Tanggal,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.RakCode,a.JenisRakOutgoing,a.ExpiredDate,a.NoSJ,a.QtySJ,a.QtyOut,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete 
	FROM tr_outgoing a 
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete is null and Tanggal between ? and ?`, tanggalAwal, tanggalAkhir)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var out model.OutgoingModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&out.KodeOutgoing, &out.Tanggal, &out.RestoCode, &out.RestoName, &out.RestoCategory, &out.ProductCode, &out.ProductName, &out.ProductCategory, &out.RakCode, &out.JenisRakOutgoing, &out.ExpDate, &out.NoSJ, &out.QtySJ, &out.QtyOut, &out.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			out.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			out.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			out.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			out.DateTimeDelete = &parsedTime
		}

		outs = append(outs, out)

	}

	return outs, nil
}

func (r *outgoingRepository) AddOutgoing(req model.OutgoingModel, id string) (model.OutgoingModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.OutgoingModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateinsert := time.Now()
	formattedDate := dateinsert.Format("2006-01-02 15:04:05")

	insertQuery := `
		INSERT INTO tr_outgoing (Tanggal,RestoCode,ProductCode,RakCode,JenisRakOutgoing,ExpiredDate,NoSJ,QtySJ,QtyOut,Satuan,UserEntry,DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.RestoCode, req.ProductCode, req.RakCode, req.JenisRakOutgoing, req.ExpDate, req.NoSJ, req.QtySJ, req.QtyOut, req.Satuan, id, formattedDate)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	var existingQty float64
	checkQuery := `SELECT Qty 
					FROM mst_rak_isi 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
	err = tx.QueryRow(checkQuery, req.RakCode, req.ProductCode, req.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.OutgoingModel{}, err
	}

	if err == sql.ErrNoRows {
		return model.OutgoingModel{}, errors.New("barang yg akan dikirim tidak ada di rak")
	} else {
		newQtye := existingQty - req.QtyOut
		updateRakIsi := `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
		_, err = tx.Exec(updateRakIsi, newQtye, req.RakCode, req.ProductCode, req.ExpDate)
		if err != nil {
			return model.OutgoingModel{}, err
		}
	}

	deleteNol := "DELETE from mst_rak_isi where Qty=0"
	_, err = tx.Exec(deleteNol)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.OutgoingModel{}, err
	}

	return req, nil
}

func (r *outgoingRepository) UpdateOutgoing(req model.OutgoingModel, currentOutgoing model.OutgoingModel, id string) (model.OutgoingModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.OutgoingModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertQuery := `
		UPDATE tr_outgoing set Tanggal=?,RestoCode=?,ProductCode=?,RakCode=?,JenisRakOutgoing=?,ExpiredDate=?,NoSJ=?,QtySJ=?,QtyOut=?,Satuan=?,UserUpdate=?,DateTimeUpdate=? 
		WHERE KodeOutgoing=? and DateTimeDelete is null
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.RestoCode, req.ProductCode, req.RakCode, req.JenisRakOutgoing, req.ExpDate, req.NoSJ, req.QtySJ, req.QtyOut, req.Satuan, id, formattedDate, currentOutgoing.KodeOutgoing)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	var existingQty float64
	checkQuery := `
		SELECT Qty 
		FROM mst_rak_isi 
		WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
	`
	err = tx.QueryRow(checkQuery, currentOutgoing.RakCode, currentOutgoing.ProductCode, currentOutgoing.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return model.OutgoingModel{}, err
	}

	if err == sql.ErrNoRows {
		raks, err := r.GetRakIsiExceptOut(req.RakCode, req.ProductCode, req.ExpDate)
		if err != nil {
			return model.OutgoingModel{}, err
		}

		if len(raks) > 0 {
			tx.Rollback()
			return model.OutgoingModel{}, errors.New("rak tujuan memiliki barang dan exp yang berbeda")
		}

		insertToRakIsi := `
			INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
			VALUES (?, ?, ?, ?)
		`
		_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.QtyOut, req.ExpDate)
		if err != nil {
			return model.OutgoingModel{}, err
		}
	} else {
		newQty := existingQty + currentOutgoing.QtyOut
		updateRakIsi := `
			UPDATE mst_rak_isi 
			SET Qty = ? 
			WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
		`
		_, err = tx.Exec(updateRakIsi, newQty, currentOutgoing.RakCode, currentOutgoing.ProductCode, currentOutgoing.ExpDate)
		if err != nil {
			return model.OutgoingModel{}, err
		}

		if req.RakCode != currentOutgoing.RakCode || req.ProductCode != currentOutgoing.ProductCode || req.ExpDate != currentOutgoing.ExpDate {

			raks, err := r.GetRakIsiExceptOut(req.RakCode, req.ProductCode, req.ExpDate)
			if err != nil {
				return model.OutgoingModel{}, err
			}

			if len(raks) > 0 {
				tx.Rollback()
				return model.OutgoingModel{}, errors.New("rak tujuan memiliki barang dan exp yang berbeda")
			}

			var existingQty float64
			checkQuery := `SELECT Qty 
							FROM mst_rak_isi 
							WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
			err = tx.QueryRow(checkQuery, req.RakCode, req.ProductCode, req.ExpDate).Scan(&existingQty)
			if err != nil && err != sql.ErrNoRows {
				return model.OutgoingModel{}, err
			}

			if err == sql.ErrNoRows {
				insertToRakIsi := `
				INSERT INTO mst_rak_isi (RakCode, ProductCode, Qty, ExpiredDate) 
				VALUES (?, ?, ?, ?)`
				_, err = tx.Exec(insertToRakIsi, req.RakCode, req.ProductCode, req.QtyOut, req.ExpDate)
				if err != nil {
					return model.OutgoingModel{}, err
				}
			} else {
				newQtye := existingQty - req.QtyOut
				updateRakIsi = `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
				_, err = tx.Exec(updateRakIsi, newQtye, req.RakCode, req.ProductCode, req.ExpDate)
				if err != nil {
					return model.OutgoingModel{}, err
				}
			}

		} else {
			newQty = newQty - req.QtyOut
			updateRakIsi = `
				UPDATE mst_rak_isi 
				SET Qty = ? 
				WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
			`
			_, err = tx.Exec(updateRakIsi, newQty, currentOutgoing.RakCode, currentOutgoing.ProductCode, currentOutgoing.ExpDate)
			if err != nil {
				return model.OutgoingModel{}, err
			}
		}
	}

	deleteNol := "DELETE from mst_rak_isi where Qty=0"
	_, err = tx.Exec(deleteNol)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.OutgoingModel{}, err
	}

	return req, nil
}

func (r *outgoingRepository) DeleteOutgoing(currentOutgoing model.OutgoingModel, id string) error {
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
		UPDATE tr_outgoing set UserDelete=?,DateTimeDelete=? 
		WHERE KodeOutgoing=? and DateTimeDelete is null
	`
	_, err = tx.Exec(insertQuery, id, formattedDate, currentOutgoing.KodeOutgoing)
	if err != nil {
		return err
	}

	var existingQty float64
	checkQuery := `SELECT Qty 
					FROM mst_rak_isi 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?`
	err = tx.QueryRow(checkQuery, currentOutgoing.RakCode, currentOutgoing.ProductCode, currentOutgoing.ExpDate).Scan(&existingQty)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		return errors.New("barang yg akan dihapus tidak ada di rak")
	} else {
		newQtye := existingQty + currentOutgoing.QtyOut
		updateRakIsi := `
					UPDATE mst_rak_isi 
					SET Qty = ? 
					WHERE RakCode = ? AND ProductCode = ? AND ExpiredDate = ?
				`
		_, err = tx.Exec(updateRakIsi, newQtye, currentOutgoing.RakCode, currentOutgoing.ProductCode, currentOutgoing.ExpDate)
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

func (r *outgoingRepository) GetRakIsiExceptOut(code, productCode, exp string) ([]model.RakIsiModel, error) {
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
