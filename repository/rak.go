package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gts-dry/model"
	"gts-dry/util"
	"strings"
	"time"
)

type RakRepository interface {
	GetRakAll() ([]model.RakModel, error)
	GetRakByKode(kode string) (*model.RakModel, error)
	GetRakByType(types string) ([]model.RakModel, error)
	GetRakByJenis(jenis string) ([]model.RakModel, error)
	GetRakIsi(code string) ([]model.RakIsiModel, error)
	GetRakIsiByProductCode(productCode string) ([]model.RakIsiModel, error)
	GetRakIsiByProductRakExp(productCode, rakCode, exp string) (*model.RakIsiModel, error)
	AddRak(req model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error)
	UpdateRak(req model.RakModelWithoutUser, current model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error)
	DeleteRak(kode string) error
	CekRakisAvailable(productCode, rakCode string, expDate string) error
	CekRakListAvailableIncoming(product, kategori, exp string) ([]model.RakIsiModel, error)
	// pindah barang rak ke rak
}

type rakRepository struct {
	db *sql.DB
}

func NewRakRepository(db *sql.DB) RakRepository {
	return &rakRepository{db: db}
}

func (r *rakRepository) GetRakAll() ([]model.RakModel, error) {
	var raks []model.RakModel

	rows, err := r.db.Query("SELECT Type,Code,JenisRak,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_rak where DateTimeDelete is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rak model.RakModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&rak.Type, &rak.Code, &rak.JenisRak, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			rak.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			rak.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			rak.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeDelete = &parsedTime
		}

		raks = append(raks, rak)

	}

	return raks, nil
}

func (r *rakRepository) GetRakByKode(kode string) (*model.RakModel, error) {
	var rak model.RakModel
	var dateTimeEntry sql.NullString
	var dateTimeUpdated sql.NullString
	var dateTimeDeleted sql.NullString
	var userEntry sql.NullString
	var userUpdate sql.NullString
	var userDelete sql.NullString

	row := r.db.QueryRow("SELECT Type,Code,JenisRak,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_rak where DateTimeDelete is null and Code = ?", kode)
	err := row.Scan(&rak.Type, &rak.Code, &rak.JenisRak, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("rak tidak ditemukan")
		}
		return nil, err
	}

	if userEntry.Valid {
		rak.UserEntry = userEntry.String
	}
	if dateTimeEntry.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
		if err != nil {
			return nil, err
		}
		rak.DateTimeEntry = &parsedTime
	}
	if userUpdate.Valid {
		rak.UserUpdate = userUpdate.String
	}
	if dateTimeUpdated.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
		if err != nil {
			return nil, err
		}
		rak.DateTimeUpdate = &parsedTime
	}
	if userDelete.Valid {
		rak.UserDelete = userDelete.String
	}
	if dateTimeDeleted.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
		if err != nil {
			return nil, err
		}
		rak.DateTimeDelete = &parsedTime
	}

	return &rak, nil
}

func (r *rakRepository) GetRakByType(types string) ([]model.RakModel, error) {
	var raks []model.RakModel

	rows, err := r.db.Query("SELECT Type,Code,JenisRak,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_rak where Type='" + types + "' and DateTimeDelete is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rak model.RakModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&rak.Type, &rak.Code, &rak.JenisRak, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			rak.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			rak.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			rak.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeDelete = &parsedTime
		}

		raks = append(raks, rak)

	}

	return raks, nil
}

func (r *rakRepository) GetRakByJenis(jenis string) ([]model.RakModel, error) {
	var raks []model.RakModel

	rows, err := r.db.Query("SELECT Type,Code,JenisRak,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_rak where JenisRak='" + jenis + "' and DateTimeDelete is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rak model.RakModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&rak.Type, &rak.Code, &rak.JenisRak, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			rak.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			rak.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			rak.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			rak.DateTimeDelete = &parsedTime
		}

		raks = append(raks, rak)

	}

	return raks, nil
}

func (r *rakRepository) GetRakIsi(code string) ([]model.RakIsiModel, error) {
	var raks []model.RakIsiModel

	rows, err := r.db.Query(`SELECT 
	ri.RakCode,ri.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,ri.Qty,ri.ExpiredDate 
	FROM mst_rak_isi ri 
	LEFT JOIN 
        mst_rak r ON ri.RakCode = r.Code and r.DateTimeDelete is null
    LEFT JOIN 
        mst_dry b ON ri.ProductCode = b.Code and b.DateTimeDelete is null
	where ri.RakCode=? order by ri.ExpiredDate asc`, code)
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

func (r *rakRepository) GetRakIsiByProductCode(productCode string) ([]model.RakIsiModel, error) {
	var raks []model.RakIsiModel

	rows, err := r.db.Query(`SELECT 
	ri.RakCode,r.JenisRak,ri.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,ri.Qty,ri.ExpiredDate 
	FROM mst_rak_isi ri 
	LEFT JOIN 
        mst_rak r ON ri.RakCode = r.Code and r.DateTimeDelete is null
    LEFT JOIN 
        mst_dry b ON ri.ProductCode = b.Code and b.DateTimeDelete is null
	where ri.ProductCode=? order by ri.ExpiredDate asc`, productCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rak model.RakIsiModel

		if err := rows.Scan(&rak.RakCode, &rak.JenisRak, &rak.ProductCode, &rak.ProductName, &rak.ProductCategory, &rak.Qty, &rak.ExpDate); err != nil {
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

func (r *rakRepository) GetRakIsiByProductRakExp(productCode, rakCode, exp string) (*model.RakIsiModel, error) {
	var rak model.RakIsiModel

	row := r.db.QueryRow(`SELECT 
				ri.RakCode,ri.ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory,ri.Qty,ri.ExpiredDate 
				FROM mst_rak_isi ri 
				LEFT JOIN 
					mst_rak r ON ri.RakCode = r.Code and r.DateTimeDelete is null
				LEFT JOIN 
					mst_dry b ON ri.ProductCode = b.Code and b.DateTimeDelete is null
				where ri.ProductCode=? and ri.RakCode=? and ri.ExpiredDate=? order by ri.ExpiredDate asc`, productCode, rakCode, exp)
	err := row.Scan(&rak.RakCode, &rak.ProductCode, &rak.ProductName, &rak.ProductCategory, &rak.Qty, &rak.ExpDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tidak ditemukan")
		}
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

	return &rak, nil
}

func (r *rakRepository) CekRakListAvailableIncoming(product, kategori, exp string) ([]model.RakIsiModel, error) {
	var result []model.RakIsiModel

	query := ""
	if strings.ToUpper(kategori) == "FOOD" {
		query = fmt.Sprintf(`
		SELECT COALESCE(ri.RakCode,r.Code) as RakCode,r.JenisRak, COALESCE(ri.ProductCode,'') as ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, COALESCE(ri.Qty,0) as Qty, COALESCE(ri.ExpiredDate,'') as ExpiredDate
		FROM mst_rak r
		LEFT JOIN mst_rak_isi ri ON r.Code = ri.RakCode 
		LEFT JOIN mst_dry b ON ri.ProductCode = b.Code and b.DateTimeDelete is null
		WHERE r.DateTimeDelete IS NULL AND r.JenisRak <> 'STOK-EQUIPMENT' AND (COALESCE(ri.ProductCode,'')='%s' OR COALESCE(ri.ProductCode,'')='') AND (COALESCE(ri.ExpiredDate,'')='%s' OR COALESCE(ri.ExpiredDate,'')='')
		ORDER BY CASE 
		WHEN COALESCE(ri.ExpiredDate, '') = '' THEN 1 
		ELSE 0 
	  END, ri.ExpiredDate asc , r.Code ASC
	`, product, exp)
	} else {
		query = `
		SELECT COALESCE(ri.RakCode,r.Code) as RakCode,r.JenisRak, COALESCE(ri.ProductCode,'') as ProductCode,COALESCE(b.Name,'') as ProductName,COALESCE(b.Kategori,'') as ProductCategory, COALESCE(ri.Qty,0) as Qty, COALESCE(ri.ExpiredDate,'') as ExpiredDate
		FROM mst_rak r
		LEFT JOIN mst_rak_isi ri ON r.Code = ri.RakCode 
		LEFT JOIN mst_dry b ON ri.ProductCode = b.Code and b.DateTimeDelete is null
		WHERE r.DateTimeDelete IS NULL AND r.JenisRak <> 'STOK-FOOD'
		ORDER BY r.Code ASC
	`
	}

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var rak model.RakIsiModel

		if err := rows.Scan(&rak.RakCode, &rak.JenisRak, &rak.ProductCode, &rak.ProductName, &rak.ProductCategory, &rak.Qty, &rak.ExpDate); err != nil {
			return nil, err
		}

		if strings.ToUpper(kategori) == "FOOD" {
			tr, stts := util.CalculateRemainingExp(rak.ExpDate)
			rak.TimeRemaining = tr
			rak.Status = stts
		} else {
			rak.TimeRemaining = 0
			rak.Status = "TANPA EXPIRED DATE"
		}
		if rak.ProductCode == "" {
			rak.Status = ""
		}

		result = append(result, rak)

	}

	return result, nil
}

func (r *rakRepository) CekRakisAvailable(productCode, rakCode string, expDate string) error {
	isiRak, err := r.GetRakIsi(rakCode)
	if err != nil {
		return err
	}

	if len(isiRak) == 0 {
		return nil
	}

	// cek apakah ada product dan exp yg beda
	isAvailable := false
	for _, ir := range isiRak {
		if ir.ProductCode != productCode {
			return errors.New("ada barang berbeda pada rak ini")
		}

		if ir.ExpDate != expDate {
			return errors.New("ada exp date berbeda pada barang di rak ini")
		}

		isAvailable = true

	}

	if isAvailable {
		return nil
	} else {
		return errors.New("rak ini tidak tersedia")
	}
}

func (r *rakRepository) AddRak(req model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertHeaderQuery := `
		INSERT INTO mst_rak (Type, Code, JenisRak, UserEntry, DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertHeaderQuery, req.Type, req.Code, req.JenisRak, id, formattedDate)
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.RakModelWithoutUser{}, err
	}

	return req, nil
}

func (r *rakRepository) UpdateRak(req model.RakModelWithoutUser, current model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertHeaderQuery := `
		UPDATE mst_rak set Type=?, Code=?, JenisRak=?, UserUpdate=?, DateTimeUpdate=?
		WHERE Code=?
	`
	_, err = tx.Exec(insertHeaderQuery, req.Type, req.Code, req.JenisRak, id, formattedDate, current.Code)
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.RakModelWithoutUser{}, err
	}

	return req, nil
}

func (r *rakRepository) DeleteRak(kode string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	deleteRak := `
		Delete from mst_rak 
		WHERE code=?
	`
	_, err = tx.Exec(deleteRak, kode)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
