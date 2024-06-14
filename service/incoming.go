package service

import (
	"errors"
	"fmt"
	"gts-dry/model"
	"gts-dry/repository"

	"github.com/go-playground/validator/v10"
)

type IncomingService interface {
	GetIncomingAll() ([]model.IncomingModel, error)
	GetIncomingByKode(kode string) (*model.IncomingModel, error)
	GetIncomingBySJ(noSJ string) ([]model.IncomingModel, error)
	GetIncomingByPO(po string) ([]model.IncomingModel, error)
	GetIncomingByPOdanProduct(po, codeProduct string) ([]model.IncomingModel, error)
	GetIncomingByPOdanProductSUMQTY(po, codeProduct string) (*float64, *float64, error)
	AddIncoming(req model.IncomingModel, id string) (model.IncomingModel, error)
	UpdateIncoming(kodeIncoming string, req model.IncomingModel, id string) (model.IncomingModel, error)
	DeleteIncoming(kodeIncoming string, id string) error
}

type incomingService struct {
	repo       repository.IncomingRepository
	rakRepo    repository.RakRepository
	barangRepo repository.BarangRepository
}

func NewIncomingService(repo repository.IncomingRepository, rakRepo repository.RakRepository, barangRepo repository.BarangRepository) IncomingService {
	return &incomingService{repo: repo, rakRepo: rakRepo, barangRepo: barangRepo}
}

func (s *incomingService) GetIncomingAll() ([]model.IncomingModel, error) {
	result, err := s.repo.GetIncomingAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *incomingService) GetIncomingByKode(kode string) (*model.IncomingModel, error) {
	result, err := s.repo.GetIncomingByKode(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *incomingService) GetIncomingBySJ(noSJ string) ([]model.IncomingModel, error) {
	result, err := s.repo.GetIncomingBySJ(noSJ)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *incomingService) GetIncomingByPO(po string) ([]model.IncomingModel, error) {
	result, err := s.repo.GetIncomingByPO(po)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *incomingService) GetIncomingByPOdanProduct(po, codeProduct string) ([]model.IncomingModel, error) {
	result, err := s.repo.GetIncomingByPOdanProduct(po, codeProduct)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *incomingService) GetIncomingByPOdanProductSUMQTY(po, codeProduct string) (*float64, *float64, error) {
	qtySum, qtyPO, err := s.repo.GetIncomingByPOdanProductSUMQTY(po, codeProduct)
	if err != nil {
		return nil, nil, err
	}

	return qtySum, qtyPO, nil
}

func (s *incomingService) AddIncoming(req model.IncomingModel, id string) (model.IncomingModel, error) {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return model.IncomingModel{}, fmt.Errorf("error: Field '%s' failed validation with tag '%s'", err.Field(), err.Tag())
		}
	}

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.IncomingModel{}, err
	}

	if barang == nil {
		return model.IncomingModel{}, errors.New("barang tidak ditemukan")
	}

	rak, err := s.rakRepo.GetRakByKode(req.RakCode)
	if err != nil {
		return model.IncomingModel{}, err
	}

	if rak == nil {
		return model.IncomingModel{}, errors.New("rak tidak ditemukan")
	}

	if barang.Kategori == "FOOD" && rak.JenisRak == "STOK-EQUIPMENT" {
		return model.IncomingModel{}, errors.New("barang food tidak bisa ditaruh di rak equipment")
	}

	if barang.Kategori == "EQUIPMENT" && rak.JenisRak == "STOK-FOOD" {
		return model.IncomingModel{}, errors.New("barang equipment tidak bisa ditaruh di rak food")
	}

	if rak.JenisRak == "STOK-FOOD" {
		if req.ExpDate == "" {
			return model.IncomingModel{}, errors.New("barang food mesti memiliki exp date")
		}

		err = s.rakRepo.CekRakisAvailable(req.ProductCode, req.RakCode, req.ExpDate)
		if err != nil {
			return model.IncomingModel{}, err
		}
	}

	qtySum, qtyPO, err := s.GetIncomingByPOdanProductSUMQTY(req.NoPO, req.ProductCode)
	if err != nil {
		return model.IncomingModel{}, err
	}

	*qtySum += req.QtyOK

	if *qtyPO == 0 {
		*qtyPO += req.QtyPO
	} else {
		if *qtyPO != req.QtyPO {
			return model.IncomingModel{}, errors.New("qty PO pada barang ini berbeda dari qty po sebelumnya, mohon dicek kembali")
		}
	}

	if *qtySum > 0 && *qtyPO > 0 {
		if *qtySum > *qtyPO {
			return model.IncomingModel{}, errors.New("qty sudah melebihi qty pada PO")
		}
	}

	if req.ExpDate == "" {
		req.ExpDate = "0000-00-00"
	}

	result, err := s.repo.AddIncoming(req, id)
	if err != nil {
		return model.IncomingModel{}, err
	}

	return result, nil
}

func (s *incomingService) UpdateIncoming(kodeIncoming string, req model.IncomingModel, id string) (model.IncomingModel, error) {
	incoming, err := s.repo.GetIncomingByKode(kodeIncoming)
	if err != nil {
		return model.IncomingModel{}, err
	}

	if incoming == nil {
		return model.IncomingModel{}, errors.New("kode incoming tidak ditemukan")
	}

	barang, err := s.barangRepo.GetBarangByKode(req.ProductCode)
	if err != nil {
		return model.IncomingModel{}, err
	}

	if barang == nil {
		return model.IncomingModel{}, errors.New("barang tidak ditemukan")
	}

	rak, err := s.rakRepo.GetRakByKode(req.RakCode)
	if err != nil {
		return model.IncomingModel{}, err
	}

	if rak == nil {
		return model.IncomingModel{}, errors.New("rak tidak ditemukan")
	}

	if req.ExpDate == "" {
		req.ExpDate = "0000-00-00"
	}

	result, err := s.repo.UpdateIncoming(req, *incoming, id, *rak, *barang)
	if err != nil {
		return model.IncomingModel{}, err
	}

	return result, nil

}

func (s *incomingService) DeleteIncoming(kodeIncoming string, id string) error {
	incoming, err := s.repo.GetIncomingByKode(kodeIncoming)
	if err != nil {
		return err
	}

	if incoming == nil {
		return errors.New("kode incoming tidak ditemukan")
	}

	err = s.repo.DeleteIncoming(*incoming, id)
	if err != nil {
		return err
	}

	return nil
}
