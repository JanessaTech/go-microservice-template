package repositories

import (
	"context"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
	"github.com/hi-supergirl/go-microservice-template/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountDB interface {
	GetById(ctx context.Context, id uint) (*model.Account, error)
	GetByName(ctx context.Context, name string) (*model.Account, error)
	Save(ctx context.Context, account model.Account) (*model.Account, error)
}

type accountDB struct {
	db *gorm.DB
}

func NewAccountDB(logger *zap.Logger, db *gorm.DB) AccountDB {
	return &accountDB{db: db}
}

func (accountDB *accountDB) GetById(ctx context.Context, id uint) (*model.Account, error) {
	var acc model.Account
	logger := logging.FromContext(ctx)
	if err := accountDB.db.WithContext(ctx).Where("id = ?", id).First(&acc).Error; err != nil {
		logger.Errorw("[accountDB]", "GetById", "Failed to get account by id", id)
		return nil, err
	}

	return &acc, nil
}
func (accountDB *accountDB) GetByName(ctx context.Context, name string) (*model.Account, error) {
	var acc model.Account
	logger := logging.FromContext(ctx)
	if err := accountDB.db.WithContext(ctx).Where("user_name = ?", name).First(&acc).Error; err != nil {
		logger.Errorw("[accountDB]", "GetByName", "Failed to save account due to ", err)
		return nil, err
	}

	return &acc, nil
}

func (accountDB *accountDB) Save(ctx context.Context, account model.Account) (*model.Account, error) {
	logger := logging.FromContext(ctx)
	if err := accountDB.db.WithContext(ctx).Create(&account).Error; err != nil {
		logger.Errorw("[accountDB]", "Save", "Failed to save account due to ", err)
		return nil, err
	}
	return &account, nil
}
