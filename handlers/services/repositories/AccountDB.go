package repositories

import (
	"errors"
	"fmt"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
)

type AccountDB struct {
	accounts map[string]*model.Account
}

func NewAccountDB() *AccountDB {
	// initiate accounts by test data
	accounts := make(map[string]*model.Account)
	return &AccountDB{accounts: accounts}
}

func (db *AccountDB) GetById(id int) (*model.Account, error) {
	name := db.getAccountNameById(id)
	if name == "" {
		return nil, fmt.Errorf("cannot find account by id %d", id)
	}
	acc, err := db.GetByName(name)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
func (db *AccountDB) GetByName(name string) (*model.Account, error) {
	val, ok := db.accounts[name]
	if !ok {
		return nil, errors.New("cannot find account by name " + name)
	}
	return val, nil
}

func (db *AccountDB) Save(account model.Account) (*model.Account, error) {
	acc, ok := db.accounts[account.UserName]
	if ok {
		return nil, errors.New("account " + acc.UserName + " already exists")
	}
	db.accounts[account.UserName] = &account
	return &account, nil
}

func (db *AccountDB) getAccountNameById(id int) string {
	var name string
	for _, acc := range db.accounts {
		if acc.ID == id {
			name = acc.UserName
		}
	}
	return name
}
