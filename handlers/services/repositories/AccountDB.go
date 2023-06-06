package repositories

import (
	"errors"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
)

type AccountDB struct {
	accounts map[string]*model.Account
}

func NewAccountDB() *AccountDB {
	// initiate accounts by test data
	accounts := make(map[string]*model.Account)
	accounts["JanessaTech1"] = &model.Account{ID: 1, UserName: "JanessaTech1", Password: "12345"}
	accounts["JanessaTech2"] = &model.Account{ID: 2, UserName: "JanessaTech2", Password: "12345"}
	return &AccountDB{accounts: accounts}
}

func (db *AccountDB) GetAll() ([]*model.Account, error) {
	accs := make([]*model.Account, len(db.accounts))
	for _, value := range db.accounts {
		accs = append(accs, value)
	}
	return accs, nil
}

func (db *AccountDB) GetById(id int) (*model.Account, error) {
	return nil, nil
}
func (db *AccountDB) GetByName(name string) (*model.Account, error) {
	val, ok := db.accounts[name]
	if !ok {
		return nil, errors.New("cannot find account by name " + name)
	}
	return val, nil
}

func (db *AccountDB) Save(account model.Account) (*model.Account, error) {
	return nil, nil
}
func (db *AccountDB) DeleteById(id int) error {
	return nil
}
