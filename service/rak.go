package service

import (
	"errors"
	"gts-dry/model"
	"gts-dry/repository"
	"strings"
)

type RakService interface {
	GetRakAll() ([]model.RakModel, error)
	GetRakByType(types string) ([]model.RakModelResponse, error)
	GetRakByKodeRak(codeRak string) (*model.RakModelResponse, error)
	GetRakByJenis(jenis string) ([]model.RakModelResponse, error)
	GetRakIsiByProductCode(productCode string) ([]model.RakIsiModel, error)
	GetRakIsiByProductRakExp(productCode, rakCode, exp string) (*model.RakIsiModel, error)
	CekRakListAvailableIncoming(product, kategori, exp string) ([]model.RakIsiModel, error)
	AddRak(req model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error)
	UpdateRak(code string, req model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error)
	DeleteRak(code string) error
}

type rakService struct {
	repo repository.RakRepository
}

func NewRakService(repo repository.RakRepository) RakService {
	return &rakService{repo: repo}
}

func (s *rakService) GetRakAll() ([]model.RakModel, error) {
	// var raks []model.RakModelResponse

	rakData, err := s.repo.GetRakAll()
	if err != nil {
		return nil, err
	}

	// for _, rak := range rakData {
	// 	isiRak, err := s.repo.GetRakIsi(rak.Code)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	hasil := model.RakModelResponse{
	// 		Rak:    rak,
	// 		RakIsi: isiRak,
	// 	}

	// 	raks = append(raks, hasil)
	// }

	return rakData, nil
}

func (s *rakService) GetRakByType(types string) ([]model.RakModelResponse, error) {
	var raks []model.RakModelResponse

	rakData, err := s.repo.GetRakByType(types)
	if err != nil {
		return nil, err
	}

	for _, rak := range rakData {
		isiRak, err := s.repo.GetRakIsi(rak.Code)
		if err != nil {
			return nil, err
		}

		hasil := model.RakModelResponse{
			Rak:    rak,
			RakIsi: isiRak,
		}

		raks = append(raks, hasil)
	}

	return raks, nil
}

func (s *rakService) GetRakByJenis(jenis string) ([]model.RakModelResponse, error) {
	var raks []model.RakModelResponse

	rakData, err := s.repo.GetRakByJenis(jenis)
	if err != nil {
		return nil, err
	}

	for _, rak := range rakData {
		isiRak, err := s.repo.GetRakIsi(rak.Code)
		if err != nil {
			return nil, err
		}

		hasil := model.RakModelResponse{
			Rak:    rak,
			RakIsi: isiRak,
		}

		raks = append(raks, hasil)
	}

	return raks, nil
}

func (s *rakService) GetRakByKodeRak(codeRak string) (*model.RakModelResponse, error) {
	rak, err := s.repo.GetRakByKode(codeRak)
	if err != nil {
		return nil, err
	}

	if rak == nil {
		return nil, errors.New("kode rak tidak ditemukan")
	}

	isiRak, err := s.repo.GetRakIsi(rak.Code)
	if err != nil {
		return nil, err
	}

	hasil := model.RakModelResponse{
		Rak:    *rak,
		RakIsi: isiRak,
	}

	return &hasil, nil
}

func (s *rakService) GetRakIsiByProductCode(productCode string) ([]model.RakIsiModel, error) {
	isiRak, err := s.repo.GetRakIsiByProductCode(productCode)
	if err != nil {
		return nil, err
	}

	return isiRak, nil
}

func (s *rakService) GetRakIsiByProductRakExp(productCode, rakCode, exp string) (*model.RakIsiModel, error) {
	isiRak, err := s.repo.GetRakIsiByProductRakExp(productCode, rakCode, exp)
	if err != nil {
		return nil, err
	}

	return isiRak, nil
}

func (s *rakService) CekRakListAvailableIncoming(product, kategori, exp string) ([]model.RakIsiModel, error) {

	isiRak, err := s.repo.CekRakListAvailableIncoming(product, kategori, exp)
	if err != nil {
		return nil, err
	}

	return isiRak, nil
}

func (s *rakService) AddRak(req model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error) {
	if (strings.ToUpper(req.JenisRak) != "STOK-FOOD") && (strings.ToUpper(req.JenisRak) != "STOK-EQUIPMENT") && (strings.ToUpper(req.JenisRak) != "TRANSIT") && (strings.ToUpper(req.JenisRak) != "PREPARE") {
		return model.RakModelWithoutUser{}, errors.New("jenis rak tidak memenuhi")
	}

	result, err := s.repo.AddRak(req, id)
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}

	return result, nil
}

func (s *rakService) UpdateRak(code string, req model.RakModelWithoutUser, id string) (model.RakModelWithoutUser, error) {
	if (strings.ToUpper(req.JenisRak) != "STOK-FOOD") && (strings.ToUpper(req.JenisRak) != "STOK-EQUIPMENT") && (strings.ToUpper(req.JenisRak) != "TRANSIT") && (strings.ToUpper(req.JenisRak) != "PREPARE") {
		return model.RakModelWithoutUser{}, errors.New("jenis rak tidak memenuhi")
	}

	data, err := s.GetRakByKodeRak(code)
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}

	if data == nil {
		return model.RakModelWithoutUser{}, errors.New("kode rak tidak ditemukan")
	}

	if req.JenisRak == "STOK-FOOD" {
		checkedProducts := map[string]string{}

		isNonFood := false
		isDifferent := false

		for _, item := range data.RakIsi {
			if strings.ToUpper(item.ProductCategory) != "FOOD" {
				isNonFood = true
				break
			}

			productCode := item.ProductCode
			expDate := item.ExpDate

			key := productCode + "-" + expDate

			if _, exists := checkedProducts[key]; exists {
				isDifferent = true
				break
			}

			checkedProducts[key] = ""
		}

		if isNonFood {
			return model.RakModelWithoutUser{}, errors.New("ada barang selain food pada rak")
		}

		if isDifferent {
			return model.RakModelWithoutUser{}, errors.New("ada barang dan exp yg berbeda pada rak")
		}
	} else if req.JenisRak == "STOK-EQUIPMENT" {
		isNonEquipment := false
		for _, item := range data.RakIsi {
			if strings.ToUpper(item.ProductCategory) != "EQUIPMENT" {
				isNonEquipment = true
				break
			}
		}

		if isNonEquipment {
			return model.RakModelWithoutUser{}, errors.New("ada barang selain equipment pada rak")
		}
	}

	req.Code = code
	currentRak := model.RakModelWithoutUser{
		Type:     data.Rak.Code,
		Code:     data.Rak.Code,
		JenisRak: data.Rak.JenisRak,
	}

	result, err := s.repo.UpdateRak(req, currentRak, id)
	if err != nil {
		return model.RakModelWithoutUser{}, err
	}

	return result, nil
}

func (s *rakService) DeleteRak(code string) error {
	data, err := s.GetRakByKodeRak(code)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("kode rak tidak ditemukan")
	}

	if len(data.RakIsi) > 0 {
		return errors.New("masih ada barang di rak, mohon periksa kembali")
	}

	err = s.repo.DeleteRak(code)
	if err != nil {
		return err
	}

	return nil
}
