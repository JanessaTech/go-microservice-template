package repositories

import "github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"

type AccountDB struct{}

func NewAccountDB() *AccountDB {
	return &AccountDB{}
}

func (db *AccountDB) GetAll() ([]*model.Account, error) {
	return nil, nil
}

func (db *AccountDB) GetById(id int) (*model.Account, error) {
	return nil, nil
}
func (db *AccountDB) GetByName(name string) (*model.Account, error) {
	return nil, nil
}

func (db *AccountDB) Save(account model.Account) (*model.Account, error) {
	return nil, nil
}
