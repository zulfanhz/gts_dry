package service

import (
	"errors"
	"gts-dry/model"
	"gts-dry/repository"
	"gts-dry/util"
	"strings"
)

type BarangService interface {
	GetBarangAll() ([]model.BarangResponseModel, error)
	GetBarangByKategori(kategori string) ([]model.BarangResponseModel, error)
	GetBarangByKode(code string) (model.BarangResponseModel, error)
	SaveBarang(req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error)
	UpdateBarang(code string, req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error)
	DeleteBarang(code string) error
	SaveSatuan(req model.BarangSatuanModel, id string) (model.BarangSatuanModel, error)
	UpdateSatuan(req model.BarangSatuanModel, kodeSatuan, namaSatuan, id string) (model.BarangSatuanModel, error)
	DeleteSatuan(kodeSatuan, namaSatuan string, IsHitung int) error
}

type barangService struct {
	repo    repository.BarangRepository
	repoRak repository.RakRepository
}

func NewBarangService(repo repository.BarangRepository, repoRak repository.RakRepository) BarangService {
	return &barangService{repo: repo, repoRak: repoRak}
}

func (s *barangService) GetBarangAll() ([]model.BarangResponseModel, error) {
	barang, err := s.repo.GetBarangAll()
	if err != nil {
		return nil, err
	}

	var result []model.BarangResponseModel
	for _, brg := range barang {
		satuan, err := s.repo.GetSatuanBarang(brg.Kode)
		if err != nil {
			return nil, err
		}

		for _, sat := range satuan {
			if sat.SatuanUtama == 1 {
				brg.Satuan = &sat.Satuan
			}
		}

		stok, err := s.repo.CekStokPerProduct(brg.Kode)
		if err != nil {
			return nil, err
		}

		brg.Stock = stok.Stok

		var hsl []model.BarangSatuanModel
		if len(satuan) == 0 {
			hsl = satuan
		} else {
			hsl = util.CalculateStock(stok.Stok, satuan)
		}

		hasil := model.BarangResponseModel{
			Barang: brg,
			Satuan: hsl,
		}

		result = append(result, hasil)
	}

	return result, nil
}

func (s *barangService) GetBarangByKategori(kategori string) ([]model.BarangResponseModel, error) {
	if (strings.ToUpper(kategori) != "FOOD") && (strings.ToUpper(kategori) != "EQUIPMENT") {
		return nil, errors.New("kategori barang tidak memenuhi")
	}

	barang, err := s.repo.GetBarangByKategori(kategori)
	if err != nil {
		return nil, err
	}

	var result []model.BarangResponseModel
	for _, brg := range barang {
		satuan, err := s.repo.GetSatuanBarang(brg.Kode)
		if err != nil {
			return nil, err
		}

		for _, sat := range satuan {
			if sat.SatuanUtama == 1 {
				brg.Satuan = &sat.Satuan
			}
		}

		stok, err := s.repo.CekStokPerProduct(brg.Kode)
		if err != nil {
			return nil, err
		}

		brg.Stock = stok.Stok

		var hsl []model.BarangSatuanModel
		if len(satuan) == 0 {
			hsl = satuan
		} else {
			hsl = util.CalculateStock(stok.Stok, satuan)
		}

		hasil := model.BarangResponseModel{
			Barang: brg,
			Satuan: hsl,
		}

		result = append(result, hasil)
	}

	return result, nil
}

func (s *barangService) GetBarangByKode(code string) (model.BarangResponseModel, error) {
	barang, err := s.repo.GetBarangByKode(code)
	if err != nil {
		return model.BarangResponseModel{}, err
	}

	if barang == nil {
		return model.BarangResponseModel{}, errors.New("kode barang tidak ditemukan")
	}

	satuan, err := s.repo.GetSatuanBarang(barang.Kode)
	if err != nil {
		return model.BarangResponseModel{}, err
	}

	for _, sat := range satuan {
		if sat.SatuanUtama == 1 {
			barang.Satuan = &sat.Satuan
		}
	}

	stok, err := s.repo.CekStokPerProduct(barang.Kode)
	if err != nil {
		return model.BarangResponseModel{}, err
	}

	barang.Stock = stok.Stok

	var hsl []model.BarangSatuanModel
	if len(satuan) == 0 {
		hsl = satuan
	} else {
		hsl = util.CalculateStock(stok.Stok, satuan)
	}

	hasil := model.BarangResponseModel{
		Barang: *barang,
		Satuan: hsl,
	}

	return hasil, nil
}

func (s *barangService) SaveBarang(req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error) {
	if (strings.ToUpper(req.Barang.Kategori) != "FOOD") && (strings.ToUpper(req.Barang.Kategori) != "EQUIPMENT") {
		return model.BarangWithoutUser{}, errors.New("kategori barang tidak memenuhi")
	}

	result, err := s.repo.SaveBarang(req, id)
	if err != nil {
		return model.BarangWithoutUser{}, err
	}

	return result, nil
}

func (s *barangService) UpdateBarang(code string, req model.BarangWithoutUser, id string) (model.BarangWithoutUser, error) {
	if (strings.ToUpper(req.Barang.Kategori) != "FOOD") && (strings.ToUpper(req.Barang.Kategori) != "EQUIPMENT") {
		return model.BarangWithoutUser{}, errors.New("kategori barang tidak memenuhi")
	}

	data, err := s.repo.GetBarangByKode(code)
	if err != nil {
		return model.BarangWithoutUser{}, err
	}

	if data == nil {
		return model.BarangWithoutUser{}, errors.New("kode barang tidak ditemukan")
	}

	req.Barang.Kode = code

	result, err := s.repo.UpdateBarang(req, id)
	if err != nil {
		return model.BarangWithoutUser{}, err
	}

	return result, nil
}

func (s *barangService) DeleteBarang(code string) error {
	barang, err := s.repo.GetBarangByKode(code)
	if err != nil {
		return err
	}

	if barang == nil {
		return errors.New("kode barang tidak ditemukan")
	}

	isiRak, err := s.repoRak.GetRakIsiByProductCode(code)
	if err != nil {
		return err
	}

	if len(isiRak) > 0 {
		return errors.New("kode barang masih ada di rak, mohon periksa kembali")
	}

	err = s.repo.DeleteBarang(code)
	if err != nil {
		return err
	}

	return nil
}

func (s *barangService) SaveSatuan(req model.BarangSatuanModel, id string) (model.BarangSatuanModel, error) {
	result, err := s.repo.SaveSatuan(req, id)
	if err != nil {
		return model.BarangSatuanModel{}, err
	}

	return result, nil
}

func (s *barangService) UpdateSatuan(req model.BarangSatuanModel, kodeSatuan, namaSatuan, id string) (model.BarangSatuanModel, error) {
	result, err := s.repo.UpdateSatuan(req, kodeSatuan, namaSatuan, id)
	if err != nil {
		return model.BarangSatuanModel{}, err
	}

	return result, nil
}

func (s *barangService) DeleteSatuan(kodeSatuan, namaSatuan string, IsHitung int) error {
	err := s.repo.DeleteSatuan(kodeSatuan, namaSatuan, IsHitung)
	if err != nil {
		return err
	}

	return nil
}
