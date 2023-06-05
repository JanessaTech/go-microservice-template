package services

import (
	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
)

type AccountService struct {
}

func (accountService *AccountService) GetAll() ([]*dto.AccountDTO, error) {

	return nil, nil
}

func (accountService *AccountService) GetById(id int) (*dto.AccountDTO, error) {
	return nil, nil
}

func (accountService *AccountService) GetByName(name string) (*dto.AccountDTO, error) {
	return nil, nil
}

func (accountService *AccountService) Save(account dto.AccountDTO) (*dto.AccountDTO, error) {
	return nil, nil
}
