package services

import (
	"errors"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories"
)

type AccountService struct {
	accountDB *repositories.AccountDB
}

func NewAccountService(accountDB *repositories.AccountDB) *AccountService {
	return &AccountService{accountDB: accountDB}
}

func (accountService *AccountService) GetAll() ([]*dto.AccountDTO, error) {
	accs, err := accountService.accountDB.GetAll()
	if err != nil {
		return nil, err
	}
	dtoAccs := make([]*dto.AccountDTO, len(accs))
	for _, acc := range accs {
		dtoAcc := dto.AccountDTO{ID: acc.ID, UserName: acc.UserName, Password: acc.Password}
		dtoAccs = append(dtoAccs, &dtoAcc)
	}
	return dtoAccs, nil
}

func (accountService *AccountService) GetById(id int) (*dto.AccountDTO, error) {
	return nil, nil
}

func (accountService *AccountService) GetByName(name string) (*dto.AccountDTO, error) {
	acc, err := accountService.accountDB.GetByName(name)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: acc.ID, UserName: acc.UserName, Password: acc.Password}
	return &accDto, nil
}

func (accountService *AccountService) Save(account dto.AccountDTO) (*dto.AccountDTO, error) {
	if account.UserName == "" || account.Password == "" {
		return nil, errors.New("UserName or Password cannot be empty")
	}
	return nil, nil
}
