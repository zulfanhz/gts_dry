package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gts-dry/model"
	"strings"
	"time"
)

type BarangRepository interface {
	GetBarangAll() ([]model.BarangModel, error)
	GetBarangByKode(code string) (*model.BarangModel, error)
	GetBarangByKategori(kategori string) ([]model.BarangModel, error)
	GetSatuanBarang(code string) ([]model.BarangSatuanModel, error)
	SaveBarang(req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error)
	UpdateBarang(req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error)
	DeleteBarang(code string) error
	SaveSatuan(req model.BarangSatuanModel, id string) (model.BarangSatuanModel, error)
	UpdateSatuan(req model.BarangSatuanModel, kodeSatuan, namaSatuan, id string) (model.BarangSatuanModel, error)
	DeleteSatuan(kodeSatuan, namaSatuan string, IsHitung int) error
	// CekStokAllProduct() (float64, error)
	CekStokPerProduct(kode string) (*model.BarangModelStok, error)
	// CekStokPerProductPerExp(kode, exp string) (float64, error)
}

type barangRepository struct {
	db *sql.DB
}

func NewBarangRepository(db *sql.DB) BarangRepository {
	return &barangRepository{db: db}
}

func (r *barangRepository) GetBarangAll() ([]model.BarangModel, error) {
	var barangs []model.BarangModel

	rows, err := r.db.Query("SELECT Code,Name,Barcode,Kategori,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_dry where DateTimeDelete is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var barang model.BarangModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&barang.Kode, &barang.Nama, &barang.Barcode, &barang.Kategori, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			barang.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			barang.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			barang.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			barang.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			barang.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			barang.DateTimeDelete = &parsedTime
		}

		barangs = append(barangs, barang)

	}

	return barangs, nil
}

func (r *barangRepository) GetBarangByKategori(kategori string) ([]model.BarangModel, error) {
	var barangs []model.BarangModel

	rows, err := r.db.Query("SELECT Code,Name,Barcode,Kategori,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_dry where DateTimeDelete is null and kategori=?", kategori)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var barang model.BarangModel
		var dateTimeEntry sql.NullString
		var dateTimeUpdated sql.NullString
		var dateTimeDeleted sql.NullString
		var userEntry sql.NullString
		var userUpdate sql.NullString
		var userDelete sql.NullString

		if err := rows.Scan(&barang.Kode, &barang.Nama, &barang.Barcode, &barang.Kategori, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted); err != nil {
			return nil, err
		}

		if userEntry.Valid {
			barang.UserEntry = userEntry.String
		}
		if dateTimeEntry.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
			if err != nil {
				return nil, err
			}
			barang.DateTimeEntry = &parsedTime
		}
		if userUpdate.Valid {
			barang.UserUpdate = userUpdate.String
		}
		if dateTimeUpdated.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
			if err != nil {
				return nil, err
			}
			barang.DateTimeUpdate = &parsedTime
		}
		if userDelete.Valid {
			barang.UserDelete = userDelete.String
		}
		if dateTimeDeleted.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
			if err != nil {
				return nil, err
			}
			barang.DateTimeDelete = &parsedTime
		}

		barangs = append(barangs, barang)

	}

	return barangs, nil
}

func (r *barangRepository) GetBarangByKode(code string) (*model.BarangModel, error) {
	var barang model.BarangModel
	var dateTimeEntry sql.NullString
	var dateTimeUpdated sql.NullString
	var dateTimeDeleted sql.NullString
	var userEntry sql.NullString
	var userUpdate sql.NullString
	var userDelete sql.NullString

	row := r.db.QueryRow("SELECT Code,Name,Barcode,Kategori,UserEntry,DateTimeEntry,UserUpdate,DateTimeUpdate,UserDelete,DateTimeDelete FROM mst_dry where DateTimeDelete is null and (Code = ? OR Barcode = ?)", code, code)
	err := row.Scan(&barang.Kode, &barang.Nama, &barang.Barcode, &barang.Kategori, &userEntry, &dateTimeEntry, &userUpdate, &dateTimeUpdated, &userDelete, &dateTimeDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("barang tidak ditemukan")
		}
		return nil, err
	}

	if userEntry.Valid {
		barang.UserEntry = userEntry.String
	}
	if dateTimeEntry.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeEntry.String)
		if err != nil {
			return nil, err
		}
		barang.DateTimeEntry = &parsedTime
	}
	if userUpdate.Valid {
		barang.UserUpdate = userUpdate.String
	}
	if dateTimeUpdated.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeUpdated.String)
		if err != nil {
			return nil, err
		}
		barang.DateTimeUpdate = &parsedTime
	}
	if userDelete.Valid {
		barang.UserDelete = userDelete.String
	}
	if dateTimeDeleted.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeDeleted.String)
		if err != nil {
			return nil, err
		}
		barang.DateTimeDelete = &parsedTime
	}

	return &barang, nil
}

func (r *barangRepository) SaveBarang(req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.BarangWithoutUser{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	insertHeaderQuery := `
		INSERT INTO mst_dry (Code, Name, Barcode, Kategori, UserEntry, DateTimeEntry) 
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertHeaderQuery, req.Barang.Kode, req.Barang.Nama, req.Barang.Barcode, req.Barang.Kategori, id, formattedDate)
	if err != nil {
		return model.BarangWithoutUser{}, err
	}

	if len(req.Satuan) > 0 {
		for _, detail := range req.Barang.Satuan {
			insertSatuanQuery := `
				INSERT INTO mst_dry_satuan (Code, Satuan, Qty, Level, IsHitung) 
				VALUES (?, ?, ?, ?, ?)
			`
			_, err = tx.Exec(insertSatuanQuery, detail.Kode, strings.ToUpper(detail.Satuan), detail.Qty, detail.UrutanSatuan, detail.SatuanUtama)
			if err != nil {
				return model.BarangWithoutUser{}, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return model.BarangWithoutUser{}, err
	}

	return req, nil
}

func (r *barangRepository) UpdateBarang(req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.BarangWithoutUser{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	updateHeaderQuery := `
		UPDATE mst_dry set Name=?, Barcode=?, Kategori=?, UserUpdate=?, DateTimeUpdate=?
		where Code=? and DateTimeDelete is null
	`
	_, err = tx.Exec(updateHeaderQuery, req.Barang.Nama, req.Barang.Barcode, req.Barang.Kategori, id, formattedDate, req.Barang.Kode)
	if err != nil {
		return model.BarangWithoutUser{}, err
	}

	deleteSatuan := `delete from mst_dry_satuan where Code=?`

	_, err = tx.Exec(deleteSatuan, req.Barang.Kode)
	if err != nil {
		return model.BarangWithoutUser{}, err
	}

	fmt.Println(len(req.Satuan))

	if len(req.Satuan) > 0 {
		for _, detail := range req.Barang.Satuan {
			insertSatuanQuery := `
				INSERT INTO mst_dry_satuan (Code, Satuan, Qty, Level, IsHitung) 
				VALUES (?, ?, ?, ?, ?)
			`
			_, err = tx.Exec(insertSatuanQuery, detail.Kode, strings.ToUpper(detail.Satuan), detail.Qty, detail.UrutanSatuan, detail.SatuanUtama)
			if err != nil {
				return model.BarangWithoutUser{}, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return model.BarangWithoutUser{}, err
	}

	return req, nil
}

func (r *barangRepository) DeleteBarang(code string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	deleteSatuan := `delete from mst_dry_satuan where Code=?`

	_, err = tx.Exec(deleteSatuan, code)
	if err != nil {
		return err
	}

	deleteBarang := `
		Delete from mst_dry 
		WHERE code=?
	`
	_, err = tx.Exec(deleteBarang, code)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *barangRepository) GetSatuanBarang(code string) ([]model.BarangSatuanModel, error) {
	var satuans []model.BarangSatuanModel

	rows, err := r.db.Query("SELECT Code,Satuan,Qty,Level,IsHitung FROM mst_dry_satuan where Code='" + code + "' order by isHitung desc, Level")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var satuan model.BarangSatuanModel

		if err := rows.Scan(&satuan.Kode, &satuan.Satuan, &satuan.Qty, &satuan.UrutanSatuan, &satuan.SatuanUtama); err != nil {
			return nil, err
		}

		satuans = append(satuans, satuan)

	}

	return satuans, nil
}

func (r *barangRepository) SaveSatuan(req model.BarangSatuanModel, id string) (model.BarangSatuanModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.BarangSatuanModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	if req.SatuanUtama == 1 {
		updateSatuanQuery := `UPDATE mst_dry_satuan SET isHitung=0 WHERE code=?`
		_, err := tx.Exec(updateSatuanQuery, req.Kode)
		if err != nil {
			return model.BarangSatuanModel{}, err
		}
	}

	insertSatuanQuery := `
		INSERT INTO mst_dry_satuan (Code, Satuan, Qty, Level, IsHitung) 
		VALUES (?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(insertSatuanQuery, req.Kode, strings.ToUpper(req.Satuan), req.Qty, req.UrutanSatuan, req.SatuanUtama)
	if err != nil {
		return model.BarangSatuanModel{}, err
	}

	updateDryQuery := `
		UPDATE mst_dry 
		SET UserEntry=?, DateTimeEntry=?
		WHERE code=? and DateTimeDelete is null
	`
	_, err = tx.Exec(updateDryQuery, id, formattedDate, req.Kode)
	if err != nil {
		return model.BarangSatuanModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.BarangSatuanModel{}, err
	}

	return req, nil
}

func (r *barangRepository) UpdateSatuan(req model.BarangSatuanModel, kodeSatuan, namaSatuan, id string) (model.BarangSatuanModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.BarangSatuanModel{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dateUpdate := time.Now()
	formattedDate := dateUpdate.Format("2006-01-02 15:04:05")

	if req.SatuanUtama == 1 {
		updateIsHitungQuery := `UPDATE mst_dry_satuan SET isHitung=0 WHERE code=?`
		_, err := tx.Exec(updateIsHitungQuery, kodeSatuan)
		if err != nil {
			return model.BarangSatuanModel{}, err
		}
	}

	updateSatuanQuery := `
		UPDATE mst_dry_satuan 
		SET Satuan=?, Qty=?, Level=?, IsHitung=? 
		WHERE code=? AND satuan=?
	`
	_, err = tx.Exec(updateSatuanQuery, strings.ToUpper(req.Satuan), req.Qty, req.UrutanSatuan, req.SatuanUtama, kodeSatuan, namaSatuan)
	if err != nil {
		return model.BarangSatuanModel{}, err
	}

	updateDryQuery := `
		UPDATE mst_dry 
		SET UserUpdate=?, DateTimeUpdate=? 
		WHERE code=? and DateTimeDelete is null
	`
	_, err = tx.Exec(updateDryQuery, id, formattedDate, kodeSatuan)
	if err != nil {
		return model.BarangSatuanModel{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.BarangSatuanModel{}, err
	}

	return req, nil
}

func (r *barangRepository) DeleteSatuan(kodeSatuan, namaSatuan string, IsHitung int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if IsHitung == 1 {
		var brs model.BarangSatuanModel
		row := tx.QueryRow(`
			SELECT Code, Satuan, Qty, Level, IsHitung 
			FROM mst_dry_satuan 
			WHERE Code = ? AND Satuan NOT IN (?) 
			ORDER BY IsHitung DESC, Level 
			LIMIT 1
		`, kodeSatuan, namaSatuan)

		err := row.Scan(&brs.Kode, &brs.Satuan, &brs.Qty, &brs.UrutanSatuan, &brs.SatuanUtama)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if brs.Kode != "" {
			updateIsHitungQuery := `UPDATE mst_dry_satuan SET IsHitung=1 WHERE Code=? AND Satuan=?`
			_, err := tx.Exec(updateIsHitungQuery, brs.Kode, brs.Satuan)
			if err != nil {
				return err
			}
		}
	}

	deleteSatuanQuery := `DELETE FROM mst_dry_satuan WHERE code=? AND satuan=? and DateTimeDelete is null`
	_, err = tx.Exec(deleteSatuanQuery, kodeSatuan, namaSatuan)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

//	func (r *barangRepository) CekStokAllProduct() (model.BarangModelStok, error) {
//		// SELECT a.`Code`,a.`Name`,a.Barcode, COALESCE((SELECT SUM(x.Qty) FROM mst_rak_isi x WHERE x.ProductCode=a.Code),0)  as Qty FROM mst_dry a
//	}
func (r *barangRepository) CekStokPerProduct(kode string) (*model.BarangModelStok, error) {
	var stok model.BarangModelStok

	row := r.db.QueryRow("SELECT a.Code,a.Name,a.Barcode, COALESCE((SELECT SUM(x.Qty) FROM mst_rak_isi x WHERE x.ProductCode=a.Code),0)  as Qty FROM mst_dry a WHERE a.Code=?", kode)
	err := row.Scan(&stok.Kode, &stok.Nama, &stok.Barcode, &stok.Stok)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("barang tidak ditemukan")
		}
		return nil, err
	}

	return &stok, nil
}

// func (r *barangRepository) CekStokPerProductPerExp(kode, exp string) (float64, error) {
// 	// 	SELECT ProductCode,ExpiredDate, SUM(Qty) as Qty FROM mst_rak_isi WHERE ProductCode='' AND ExpiredDate=''
// 	// GROUP BY ProductCode,ExpiredDate
// }
