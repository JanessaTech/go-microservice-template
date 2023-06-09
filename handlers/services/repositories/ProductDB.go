package repositories

import (
	"context"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
	"github.com/hi-supergirl/go-microservice-template/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductDB interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	Add(ctx context.Context, product model.Product) (*model.Product, error)
	Delete(ctx context.Context, id uint) error
}

type productDB struct {
	db *gorm.DB
}

func NewProductDB(logger *zap.Logger, db *gorm.DB) ProductDB {
	return &productDB{db: db}
}

func (productDB *productDB) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	logger := logging.FromContext(ctx)
	if err := productDB.db.WithContext(ctx).Find(&products).Error; err != nil {
		logger.Debugw("[productDB]", "GetAll", "Failed to get all products due to ", err.Error())
		return nil, err
	}
	return products, nil
}

func (productDB *productDB) Add(ctx context.Context, product model.Product) (*model.Product, error) {
	logger := logging.FromContext(ctx)
	if err := productDB.db.WithContext(ctx).Create(&product).Error; err != nil {
		logger.Debugw("[productDB]", "Add", "Failed to add product due to ", err.Error())
		return nil, err
	}
	return &product, nil
}

func (productDB *productDB) Delete(ctx context.Context, id uint) error {
	logger := logging.FromContext(ctx)
	if err := productDB.db.Delete(model.Product{ID: id}).Error; err != nil {
		logger.Debugw("[productDB]", "Delete", "Failed to delete product id ", id, " due to ", err.Error())
		return err
	}
	return nil
}
