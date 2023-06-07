package services

import (
	"context"
	"errors"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
)

type AccountService interface {
	GetById(ctx context.Context, id int) (*dto.AccountDTO, error)
	GetByName(ctx context.Context, name string) (*dto.AccountDTO, error)
	Save(ctx context.Context, accountDto dto.AccountDTO) (*dto.AccountDTO, error)
}

type accountService struct {
	accountDB repositories.AccountDB
}

func NewAccountService(accountDB repositories.AccountDB) AccountService {
	return &accountService{accountDB: accountDB}
}

func (accountService *accountService) GetById(ctx context.Context, id int) (*dto.AccountDTO, error) {
	acc, err := accountService.accountDB.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: acc.ID, UserName: acc.UserName, Password: acc.Password}
	return &accDto, nil
}

func (accountService *accountService) GetByName(ctx context.Context, name string) (*dto.AccountDTO, error) {
	acc, err := accountService.accountDB.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: acc.ID, UserName: acc.UserName, Password: acc.Password}
	return &accDto, nil
}

func (accountService *accountService) Save(ctx context.Context, accountDto dto.AccountDTO) (*dto.AccountDTO, error) {
	if accountDto.UserName == "" || accountDto.Password == "" {
		return nil, errors.New("UserName or Password cannot be empty")
	}
	acc := model.Account{ID: accountDto.ID, UserName: accountDto.UserName, Password: accountDto.Password}
	savedAcc, err := accountService.accountDB.Save(ctx, acc)
	if err != nil {
		return nil, err
	}
	accDto := dto.AccountDTO{ID: savedAcc.ID, UserName: savedAcc.UserName, Password: savedAcc.Password}

	return &accDto, nil
}
