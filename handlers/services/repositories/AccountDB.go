package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
	"gorm.io/gorm"
)

type AccountDB interface {
	GetById(ctx context.Context, id int) (*model.Account, error)
	GetByName(ctx context.Context, name string) (*model.Account, error)
	Save(ctx context.Context, account model.Account) (*model.Account, error)
}

type accountDB struct {
	db       *gorm.DB
	accounts map[string]*model.Account
}

func NewAccountDB(db *gorm.DB) AccountDB {
	// initiate accounts by test data
	accounts := make(map[string]*model.Account)
	return &accountDB{db: db, accounts: accounts}
}

func (db *accountDB) GetById(ctx context.Context, id int) (*model.Account, error) {
	name := db.getAccountNameById(id)
	if name == "" {
		return nil, fmt.Errorf("cannot find account by id %d", id)
	}
	acc, err := db.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
func (db *accountDB) GetByName(ctx context.Context, name string) (*model.Account, error) {
	val, ok := db.accounts[name]
	if !ok {
		return nil, errors.New("cannot find account by name " + name)
	}
	return val, nil
}

func (db *accountDB) Save(ctx context.Context, account model.Account) (*model.Account, error) {
	acc, ok := db.accounts[account.UserName]
	if ok {
		return nil, errors.New("account " + acc.UserName + " already exists")
	}
	db.accounts[account.UserName] = &account
	return &account, nil
}

func (db *accountDB) getAccountNameById(id int) string {
	var name string
	for _, acc := range db.accounts {
		if acc.ID == id {
			name = acc.UserName
		}
	}
	return name
}
