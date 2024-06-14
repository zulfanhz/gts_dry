package service

import (
	"errors"
	"gts-dry/model"
	"gts-dry/repository"
)

type OutgoingService interface {
	GetOutgoingAll() ([]model.OutgoingModel, error)
	GetOutgoingByKode(kode string) (*model.OutgoingModel, error)
	GetOutgoingBySJ(kode string) ([]model.OutgoingModel, error)
	GetOutgoingByProduct(kode string) ([]model.OutgoingModel, error)
	GetOutgoingByResto(kode string) ([]model.OutgoingModel, error)
	GetOutgoingByTanggal(tanggalAwal, tanggalAkhir string) ([]model.OutgoingModel, error)
	AddOutgoing(req model.OutgoingModel, id string) (model.OutgoingModel, error)
	UpdateOutgoing(kodeOutgoing string, req model.OutgoingModel, id string) (model.OutgoingModel, error)
	DeleteOutgoing(kodeOutgoing string, id string) error
}

type outgoingService struct {
	repo       repository.OutgoingRepository
	rakRepo    repository.RakRepository
	restoRepo  repository.RestoRepository
	barangRepo repository.BarangRepository
}

func NewOutgoingService(repo repository.OutgoingRepository, rakRepo repository.RakRepository, restoRepo repository.RestoRepository, barangRepo repository.BarangRepository) OutgoingService {
	return &outgoingService{repo: repo, rakRepo: rakRepo, restoRepo: restoRepo, barangRepo: barangRepo}
}

func (s *outgoingService) GetOutgoingAll() ([]model.OutgoingModel, error) {
	result, err := s.repo.GetOutgoingAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *outgoingService) GetOutgoingByKode(kode string) (*model.OutgoingModel, error) {
	result, err := s.repo.GetOutgoingByKode(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *outgoingService) GetOutgoingBySJ(kode string) ([]model.OutgoingModel, error) {
	result, err := s.repo.GetOutgoingBySJ(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *outgoingService) GetOutgoingByProduct(kode string) ([]model.OutgoingModel, error) {
	result, err := s.repo.GetOutgoingByProduct(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *outgoingService) GetOutgoingByResto(kode string) ([]model.OutgoingModel, error) {
	result, err := s.repo.GetOutgoingByResto(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *outgoingService) GetOutgoingByTanggal(tanggalAwal, tanggalAkhir string) ([]model.OutgoingModel, error) {
	result, err := s.repo.GetOutgoingByTanggal(tanggalAwal, tanggalAkhir)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *outgoingService) AddOutgoing(req model.OutgoingModel, id string) (model.OutgoingModel, error) {
	cekResto, err := s.restoRepo.GetRestoByKode(req.RestoCode)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if cekResto == nil {
		return model.OutgoingModel{}, errors.New("kode resto tidak ditemukan")
	}

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if barang == nil {
		return model.OutgoingModel{}, errors.New("barang tidak ditemukan")
	}

	rak, err := s.rakRepo.GetRakByKode(req.RakCode)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if rak == nil {
		return model.OutgoingModel{}, errors.New("rak tidak ditemukan")
	}

	cekRak, err := s.rakRepo.GetRakIsiByProductRakExp(req.ProductCode, req.RakCode, req.ExpDate)
	if err != nil {
		return model.OutgoingModel{}, nil
	}

	if cekRak == nil {
		return model.OutgoingModel{}, errors.New("barang tidak ditemukan di rak")
	}

	if req.QtyOut > cekRak.Qty {
		return model.OutgoingModel{}, errors.New("qty melebihi qty pada rak")
	}

	if req.ExpDate == "" {
		req.ExpDate = "0000-00-00"
	}

	result, err := s.repo.AddOutgoing(req, id)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	return result, nil
}

func (s *outgoingService) UpdateOutgoing(kodeOutgoing string, req model.OutgoingModel, id string) (model.OutgoingModel, error) {
	out, err := s.repo.GetOutgoingByKode(kodeOutgoing)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if out == nil {
		return model.OutgoingModel{}, errors.New("kode adjustment tidak ditemukan")
	}

	cekResto, err := s.restoRepo.GetRestoByKode(req.RestoCode)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if cekResto == nil {
		return model.OutgoingModel{}, errors.New("kode resto tidak ditemukan")
	}

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if barang == nil {
		return model.OutgoingModel{}, errors.New("barang tidak ditemukan")
	}

	rak, err := s.rakRepo.GetRakByKode(req.RakCode)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	if rak == nil {
		return model.OutgoingModel{}, errors.New("rak tidak ditemukan")
	}

	cekRak, err := s.rakRepo.GetRakIsiByProductRakExp(req.ProductCode, req.RakCode, req.ExpDate)
	if err != nil {
		return model.OutgoingModel{}, nil
	}

	if cekRak == nil {
		return model.OutgoingModel{}, errors.New("barang tidak ditemukan di rak")
	}

	if req.QtyOut > cekRak.Qty {
		return model.OutgoingModel{}, errors.New("qty melebihi qty pada rak")
	}

	if req.ExpDate == "" {
		req.ExpDate = "0000-00-00"
	}

	result, err := s.repo.UpdateOutgoing(req, *out, id)
	if err != nil {
		return model.OutgoingModel{}, err
	}

	return result, nil
}

func (s *outgoingService) DeleteOutgoing(kodeOutgoing string, id string) error {
	out, err := s.repo.GetOutgoingByKode(kodeOutgoing)
	if err != nil {
		return err
	}

	if out == nil {
		return errors.New("kode adjustment tidak ditemukan")
	}

	err = s.repo.DeleteOutgoing(*out, id)
	if err != nil {
		return err
	}

	return nil
}
