package repositories

import (
	"context"
	"fmt"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
	"gorm.io/gorm"
)

type AccountDB interface {
	GetById(ctx context.Context, id uint) (*model.Account, error)
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

func (accountDB *accountDB) GetById(ctx context.Context, id uint) (*model.Account, error) {
	var acc model.Account
	if err := accountDB.db.WithContext(ctx).Where("id = ?", id).First(&acc).Error; err != nil {
		fmt.Println("[accountDB.GetById] Failed to get account by id", id)
		return nil, err
	}

	return &acc, nil
}
func (accountDB *accountDB) GetByName(ctx context.Context, name string) (*model.Account, error) {
	var acc model.Account
	if err := accountDB.db.WithContext(ctx).Where("user_name = ?", name).First(&acc).Error; err != nil {
		fmt.Println("[accountDB.GetByName] Failed to save account due to ", err)
		return nil, err
	}

	return &acc, nil
}

func (accountDB *accountDB) Save(ctx context.Context, account model.Account) (*model.Account, error) {

	if err := accountDB.db.WithContext(ctx).Create(&account).Error; err != nil {
		fmt.Println("[accountDB.Save] Failed to save account due to ", err)
		return nil, err
	}
	return &account, nil
}
