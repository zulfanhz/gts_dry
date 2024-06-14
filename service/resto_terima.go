package service

import (
	"gts-dry/model"
	"gts-dry/repository"
)

type RestoTerimaService interface {
	GetTerimaRestoAll() ([]model.RestoTerimaModel, error)
	GetTerimaRestoByCodeResto(kode string) ([]model.RestoTerimaModel, error)
	AddRestoTerima(req model.RestoTerimaModel, id string) (model.RestoTerimaModel, error)
}

type restoTerimaService struct {
	repo repository.RestoTerimaRepository
}

func NewRestoTerimaService(repo repository.RestoTerimaRepository) RestoTerimaService {
	return &restoTerimaService{repo: repo}
}

func (s *restoTerimaService) GetTerimaRestoAll() ([]model.RestoTerimaModel, error) {
	data, err := s.repo.GetTerimaRestoAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *restoTerimaService) GetTerimaRestoByCodeResto(kode string) ([]model.RestoTerimaModel, error) {
	data, err := s.repo.GetTerimaRestoByCodeResto(kode)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *restoTerimaService) AddRestoTerima(req model.RestoTerimaModel, id string) (model.RestoTerimaModel, error) {
	data, err := s.repo.AddRestoTerima(req, id)
	if err != nil {
		return model.RestoTerimaModel{}, err
	}

	return data, nil
}
