package server

import (
	"gRPC/internal/pkg/auth"
	"gRPC/internal/pkg/constErr"
	"gRPC/internal/pkg/data"
	"gRPC/internal/pkg/storage"
)

type Service struct {
	IStorage storage.InterfaceStorage
}

type InterfaceServer interface {
	SignUp(acc *data.Account) (string, error)
	Login(acc *data.Account) (string, error)
	Logout(token string) error
	Add(ad *data.Ad) error
	Get(id uint) (*data.Ad, error)
	GetAll() ([]*data.Ad, error)
	Delete(id uint) error
	Update(id uint, ad *data.Ad) error
	GetStorage() storage.InterfaceStorage
}

func (s *Service) SignUp(acc *data.Account) (string, error) {
	// Получаем базу данных с Storage
	baseAcc, err := s.IStorage.GetAccounts()
	if err != nil {
		return "", err
	}
	// Если есть такой пользователь возвращать account и true, в противном случае false
	account, err := auth.IsAccountExists(acc, baseAcc)
	if err != nil {
		return "", err
	}
	// генерация токена
	token, err := auth.GenerateToken(account.GetUserName(), account.GetPassword())
	if err != nil {
		return "", err
	}

	account.SetToken(token)

	return token, nil
}

func (s *Service) Login(acc *data.Account) (string, error) {
	baseAcc, _ := s.IStorage.GetAccounts()

	_, err := auth.IsAccountExists(acc, baseAcc)
	if err == nil {
		return "", constErr.YouRat
	}

	token, err := auth.GenerateToken(acc.GetUserName(), acc.GetPassword())
	if err != nil {
		return "", constErr.FailedToGenerateAToken
	}

	acc.SetToken(token)

	if err := s.IStorage.AddAccount(acc); err != nil {
		return "", constErr.ErrorToAddAcc
	}

	return token, nil
}

//TODO
func (s *Service) Logout(token string) error {
	return nil
}

func (s *Service) Add(ad *data.Ad) error {
	if err := s.IStorage.Add(ad); err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(id uint) (*data.Ad, error) {
	ad, err := s.IStorage.Get(id) //TODO bool or err add
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (s *Service) GetAll() ([]*data.Ad, error) {
	baseAd, err := s.IStorage.GetAll()
	if err != nil {
		return nil, err
	}
	return baseAd, nil
}

func (s *Service) Delete(id uint) error {
	if err := s.IStorage.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(id uint, ad *data.Ad) error {
	if len(ad.GetBrand()) == 0 && len(ad.GetModel()) == 0 &&
		len(ad.GetColor()) == 0 && ad.GetPrice() <= 0 {
		return constErr.EmptyFields
	}
	if err := s.IStorage.Update(ad, id); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetStorage() storage.InterfaceStorage {
	return s.IStorage
}

func (s *Service) SetStorage(stg storage.InterfaceStorage) {
	s.IStorage = stg
}
