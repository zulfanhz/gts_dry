package service

import (
	"gts-dry/model"
	"gts-dry/repository"
)

type RestoService interface {
	GetRestoAll() ([]model.RestoModel, error)
	GetRestoByKategori(kategori string) ([]model.RestoModel, error)
	GetRestoByKode(kode string) (*model.RestoModel, error)
}

type restoService struct {
	repo repository.RestoRepository
}

func NewRestoService(repo repository.RestoRepository) RestoService {
	return &restoService{repo: repo}
}

func (s *restoService) GetRestoAll() ([]model.RestoModel, error) {
	result, err := s.repo.GetRestoAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *restoService) GetRestoByKategori(kategori string) ([]model.RestoModel, error) {
	result, err := s.repo.GetRestoByKategori(kategori)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *restoService) GetRestoByKode(kode string) (*model.RestoModel, error) {
	result, err := s.repo.GetRestoByKode(kode)
	if err != nil {
		return nil, err
	}

	return result, nil
}
