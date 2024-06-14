package service

import (
	"errors"
	"gts-dry/model"
	"gts-dry/repository"
)

type MutasiService interface {
	GetMutasiAll() ([]model.MutasiRak, error)
	AddMutasi(req model.MutasiRakRequest, id string) (model.MutasiRak, error)
}

type mutasiService struct {
	repo       repository.MutasiRepository
	rakRepo    repository.RakRepository
	barangRepo repository.BarangRepository
}

func NewMutasiService(repo repository.MutasiRepository, rakRepo repository.RakRepository, barangRepo repository.BarangRepository) MutasiService {
	return &mutasiService{repo: repo, rakRepo: rakRepo, barangRepo: barangRepo}
}

func (s *mutasiService) GetMutasiAll() ([]model.MutasiRak, error) {
	data, err := s.repo.GetMutasiAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *mutasiService) AddMutasi(req model.MutasiRakRequest, id string) (model.MutasiRak, error) {

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.MutasiRak{}, err
	}

	if barang == nil {
		return model.MutasiRak{}, errors.New("barang tidak ditemukan")
	}

	rakAsal, err := s.rakRepo.GetRakByKode(req.RakCodeAsal)
	if err != nil {
		return model.MutasiRak{}, err
	}

	if rakAsal == nil {
		return model.MutasiRak{}, errors.New("rak asal tidak ditemukan")
	}

	rakTujuan, err := s.rakRepo.GetRakByKode(req.RakCodeTujuan)
	if err != nil {
		return model.MutasiRak{}, err
	}

	if rakTujuan == nil {
		return model.MutasiRak{}, errors.New("rak tujuan tidak ditemukan")
	}

	if barang.Kategori == "FOOD" && rakTujuan.JenisRak == "STOK-EQUIPMENT" {
		return model.MutasiRak{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
	}

	if barang.Kategori == "EQUIPMENT" && rakTujuan.JenisRak == "STOK-FOOD" {
		return model.MutasiRak{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
	}

	if rakTujuan.JenisRak == "STOK-FOOD" {
		if req.ExpiredDate == "" {
			return model.MutasiRak{}, errors.New("barang food mesti memiliki exp date")
		}

		err = s.rakRepo.CekRakisAvailable(req.ProductCode, req.RakCodeTujuan, req.ExpiredDate)
		if err != nil {
			return model.MutasiRak{}, err
		}
	}

	request := model.MutasiRak{
		Tanggal:         req.Tanggal,
		ProductCode:     req.ProductCode,
		ProductName:     barang.Nama,
		ProductCategory: barang.Kategori,
		RakCodeAsal:     req.RakCodeAsal,
		JenisRakAsal:    rakAsal.JenisRak,
		RakCodeTujuan:   req.RakCodeTujuan,
		JenisRakTujuan:  rakTujuan.JenisRak,
		QtyMutasi:       req.QtyMutasi,
		Satuan:          req.Satuan,
		ExpiredDate:     req.ExpiredDate,
	}

	if request.ExpiredDate == "" {
		request.ExpiredDate = "0000-00-00"
	}

	result, err := s.repo.AddMutasi(request, id)
	if err != nil {
		return model.MutasiRak{}, err
	}
	return result, nil
}
