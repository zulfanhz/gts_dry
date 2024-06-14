package service

import (
	"errors"
	"gts-dry/model"
	"gts-dry/repository"
	"gts-dry/util"
)

type UserService interface {
	LoginUser(email string, password string) (model.UserResponse, *string, error)
	GetUserByEmail(email string) (*model.UserResponse, error)
	ChangePassword(email string, passwordSekarang, passwordBaru, passwordBaruRepeat, id string) (bool, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s userService) LoginUser(email string, password string) (model.UserResponse, *string, error) {
	userDB, err := s.repo.GetByUserEmail(email)
	if err != nil {
		return model.UserResponse{}, nil, err
	}

	if userDB == nil {
		return model.UserResponse{}, nil, errors.New("akun tidak ditemukan")
	}

	passwordFromDB := userDB.Password

	isTrue := util.CheckPasswordHash(password, passwordFromDB)
	if isTrue {
		token, err := util.CreateJWT(userDB.Email)
		if err != nil {
			return model.UserResponse{}, nil, err
		}

		userResponse := model.UserResponse{
			Email:    userDB.Email,
			IsActive: userDB.IsActive,
			Unit:     userDB.Unit,
		}

		return userResponse, &token, nil
	} else {
		return model.UserResponse{}, nil, errors.New("invalid email or password")
	}

}

func (s *userService) GetUserByEmail(email string) (*model.UserResponse, error) {
	userDB, err := s.repo.GetByUserEmail(email)
	if err != nil {
		return nil, err
	}

	if userDB == nil {
		return nil, errors.New("akun tidak ditemukan")
	}

	userResponse := model.UserResponse{
		Email:    userDB.Email,
		IsActive: userDB.IsActive,
		Unit:     userDB.Unit,
	}

	return &userResponse, nil
}

func (s *userService) ChangePassword(email string, passwordSekarang, passwordBaru, passwordBaruRepeat, id string) (bool, error) {
	userDB, err := s.repo.GetByUserEmail(email)
	if err != nil {
		return false, err
	}

	if userDB == nil {
		return false, errors.New("akun tidak ditemukan")
	}

	isTrue := util.CheckPasswordHash(passwordSekarang, userDB.Password)

	if !isTrue {
		return false, errors.New("password lama anda salah")
	}

	if passwordBaru != passwordBaruRepeat {
		return false, errors.New("pengulangan password baru anda salah")
	}

	pasBar, err := util.HashPassword(passwordBaru)
	if err != nil {
		return false, err
	}

	result, err := s.repo.ChangePassword(email, pasBar, id)
	if err != nil {
		return false, err
	}

	return result, nil
}
