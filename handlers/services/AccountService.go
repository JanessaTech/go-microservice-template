package services

import (
	"errors"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
)

type AccountService struct {
	accountDB *repositories.AccountDB
}

func NewAccountService(accountDB *repositories.AccountDB) *AccountService {
	return &AccountService{accountDB: accountDB}
}

func (accountService *AccountService) GetById(id int) (*dto.AccountDTO, error) {
	acc, err := accountService.accountDB.GetById(id)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: acc.ID, UserName: acc.UserName, Password: acc.Password}
	return &accDto, nil
}

func (accountService *AccountService) GetByName(name string) (*dto.AccountDTO, error) {
	acc, err := accountService.accountDB.GetByName(name)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: acc.ID, UserName: acc.UserName, Password: acc.Password}
	return &accDto, nil
}

func (accountService *AccountService) Save(accountDto dto.AccountDTO) (*dto.AccountDTO, error) {
	if accountDto.UserName == "" || accountDto.Password == "" {
		return nil, errors.New("UserName or Password cannot be empty")
	}
	acc := model.Account{ID: accountDto.ID, UserName: accountDto.UserName, Password: accountDto.Password}
	savedAcc, err := accountService.accountDB.Save(acc)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: savedAcc.ID, UserName: savedAcc.UserName, Password: savedAcc.Password}

	return &accDto, nil
}
