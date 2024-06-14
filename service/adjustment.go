package service

import (
	"errors"
	"gts-dry/model"
	"gts-dry/repository"
	"strings"
)

type AdjustmentService interface {
	GetAdjustmentAll() ([]model.AdjustmentModel, error)
	GetAdjustmentByKode(kode string) (*model.AdjustmentModel, error)
	AddAdjustment(req model.AdjustmentRequestModel, id string) (model.AdjustmentRequestModel, error)
	UpdateAdjustment(kodeAdjustment string, req model.AdjustmentRequestModel, id string) (model.AdjustmentRequestModel, error)
	DeleteAdjustment(kodeAdjustment string, id string) error
}

type adjustmentService struct {
	repo       repository.AdjustmentRepository
	rakRepo    repository.RakRepository
	barangRepo repository.BarangRepository
}

func NewAdjustmentService(repo repository.AdjustmentRepository, rakRepo repository.RakRepository, barangRepo repository.BarangRepository) AdjustmentService {
	return &adjustmentService{repo: repo, rakRepo: rakRepo, barangRepo: barangRepo}
}

func (s *adjustmentService) GetAdjustmentAll() ([]model.AdjustmentModel, error) {
	result, err := s.repo.GetAdjustmentAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s *adjustmentService) GetAdjustmentByKode(kode string) (*model.AdjustmentModel, error) {
	result, err := s.repo.GetAdjustmentByKode(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *adjustmentService) AddAdjustment(req model.AdjustmentRequestModel, id string) (model.AdjustmentRequestModel, error) {
	if (strings.ToUpper(req.JenisRakAdjustment) != "STOK-FOOD") && (strings.ToUpper(req.JenisRakAdjustment) != "STOK-EQUIPMENT") && (strings.ToUpper(req.JenisRakAdjustment) != "TRANSIT") && (strings.ToUpper(req.JenisRakAdjustment) != "PREPARE") {
		return model.AdjustmentRequestModel{}, errors.New("jenis rak tidak memenuhi")
	}

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	if barang == nil {
		return model.AdjustmentRequestModel{}, errors.New("barang tidak ditemukan")
	}

	rak, err := s.rakRepo.GetRakByKode(req.RakCode)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	if rak == nil {
		return model.AdjustmentRequestModel{}, errors.New("rak tidak ditemukan")
	}

	if barang.Kategori == "FOOD" && rak.JenisRak == "STOK-EQUIPMENT" {
		return model.AdjustmentRequestModel{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
	}

	if barang.Kategori == "EQUIPMENT" && rak.JenisRak == "STOK-FOOD" {
		return model.AdjustmentRequestModel{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
	}

	if rak.JenisRak == "STOK-FOOD" {
		if req.ExpDate == "" {
			return model.AdjustmentRequestModel{}, errors.New("barang food mesti memiliki exp date")
		}

		err = s.rakRepo.CekRakisAvailable(req.ProductCode, req.RakCode, req.ExpDate)
		if err != nil {
			return model.AdjustmentRequestModel{}, err
		}
	}

	if req.ExpDate == "" {
		req.ExpDate = "0000-00-00"
	}

	result, err := s.repo.AddAdjustment(req, id)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	return result, nil
}

func (s *adjustmentService) UpdateAdjustment(kodeAdjustment string, req model.AdjustmentRequestModel, id string) (model.AdjustmentRequestModel, error) {
	adj, err := s.repo.GetAdjustmentByKode(kodeAdjustment)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	if adj == nil {
		return model.AdjustmentRequestModel{}, errors.New("kode adjustment tidak ditemukan")
	}

	if (strings.ToUpper(req.JenisRakAdjustment) != "STOK-FOOD") && (strings.ToUpper(req.JenisRakAdjustment) != "STOK-EQUIPMENT") && (strings.ToUpper(req.JenisRakAdjustment) != "TRANSIT") && (strings.ToUpper(req.JenisRakAdjustment) != "PREPARE") {
		return model.AdjustmentRequestModel{}, errors.New("jenis rak tidak memenuhi")
	}

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	if barang == nil {
		return model.AdjustmentRequestModel{}, errors.New("barang tidak ditemukan")
	}

	rak, err := s.rakRepo.GetRakByKode(req.RakCode)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	if rak == nil {
		return model.AdjustmentRequestModel{}, errors.New("rak tidak ditemukan")
	}

	if req.ExpDate == "" {
		req.ExpDate = "0000-00-00"
	}

	result, err := s.repo.UpdateAdjustment(req, *adj, id, *rak, *barang)
	if err != nil {
		return model.AdjustmentRequestModel{}, err
	}

	return result, nil
}

func (s *adjustmentService) DeleteAdjustment(kodeAdjustment string, id string) error {
	adj, err := s.repo.GetAdjustmentByKode(kodeAdjustment)
	if err != nil {
		return err
	}

	if adj == nil {
		return errors.New("kode adjustment tidak ditemukan")
	}

	err = s.repo.DeleteAdjustment(*adj, id)
	if err != nil {
		return err
	}

	return nil
}
