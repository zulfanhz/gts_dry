package repository

import (
	"database/sql"
	"gts-dry/model"
	"time"
)

type RestoTerimaRepository interface {
	GetTerimaRestoAll() ([]model.RestoTerimaModel, error)
	GetTerimaRestoByCodeResto(kode string) ([]model.RestoTerimaModel, error)
	AddRestoTerima(req model.RestoTerimaModel, id string) (model.RestoTerimaModel, error)
}

type restoTerimaRepository struct {
	db *sql.DB
}

func NewRestoTerimaRepository(db *sql.DB) RestoTerimaRepository {
	return &restoTerimaRepository{db: db}
}

func (r *restoTerimaRepository) GetTerimaRestoAll() ([]model.RestoTerimaModel, error) {
	var trms []model.RestoTerimaModel

	rows, err := r.db.Query(`SELECT a.KodeRestoTerima,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory, a.Tanggal,a.NoSJ,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.QtySJ,a.QtyTerima,a.ExpiredDate,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete FROM tr_resto_terima a
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete IS NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trm model.RestoTerimaModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&trm.KodeRestoTerima, &trm.RestoCode, &trm.RestoName, &trm.RestoCategory, &trm.Tanggal, &trm.NoSJ, &trm.ProductCode, &trm.ProductName, &trm.ProductCategory, &trm.QtySJ, &trm.QtyTerima, &trm.ExpDate, &trm.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			trm.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			trm.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			trm.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			trm.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			trm.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			trm.DateTimeDelete = &parsedTime
		}

		trms = append(trms, trm)

	}

	return trms, nil
}

func (r *restoTerimaRepository) GetTerimaRestoByCodeResto(kode string) ([]model.RestoTerimaModel, error) {
	var trms []model.RestoTerimaModel

	rows, err := r.db.Query(`SELECT a.KodeRestoTerima,a.RestoCode,COALESCE(c.Name,'') as RestoName, COALESCE(c.Kategori,'') as RestoCategory, a.Tanggal,a.NoSJ,a.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,a.QtySJ,a.QtyTerima,a.ExpiredDate,a.Satuan,a.UserEntry,a.DateTimeEntry,a.UserUpdate,a.DateTimeUpdate,a.UserDelete,a.DateTimeDelete FROM tr_resto_terima a
	LEFT JOIN mst_dry b on a.ProductCode = b.Code AND b.DateTimeDelete IS NULL
	LEFT JOIN mst_resto c on a.RestoCode = c.Kode AND c.Aktif = 1
	WHERE a.DateTimeDelete IS NULL and a.RestoCode=?
	`, kode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trm model.RestoTerimaModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&trm.KodeRestoTerima, &trm.RestoCode, &trm.RestoName, &trm.RestoCategory, &trm.Tanggal, &trm.NoSJ, &trm.ProductCode, &trm.ProductName, &trm.ProductCategory, &trm.QtySJ, &trm.QtyTerima, &trm.ExpDate, &trm.Satuan, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			trm.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			trm.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			trm.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			trm.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			trm.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			trm.DateTimeDelete = &parsedTime
		}

		trms = append(trms, trm)

	}

	return trms, nil
}

func (r *restoTerimaRepository) AddRestoTerima(req model.RestoTerimaModel, id string) (model.RestoTerimaModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.RestoTerimaModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateinsert := time.Now()
	formattedDate := dateinsert.Format("2006-01-02 15:04:05")

	insertQuery := `
		INSERT INTO tr_resto_terima (Tanggal,RestoCode,NoSJ,ProductCode,QtySJ,QtyTerima,ExpiredDate,Satuan,UserEntry,DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertQuery, req.Tanggal, req.RestoCode, req.NoSJ, req.ProductCode, req.QtySJ, req.QtyTerima, req.ExpDate, req.Satuan, id, formattedDate)
	if err != nil {
		return model.RestoTerimaModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.RestoTerimaModel{}, err
	}

	return req, nil
}
