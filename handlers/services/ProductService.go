package services

import (
	"context"
	"errors"

	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/model"
	"github.com/hi-supergirl/go-microservice-template/logging"
	"go.uber.org/zap"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]*dto.ProductDTO, error)
	Add(ctx context.Context, productDTO dto.ProductDTO) (*dto.ProductDTO, error)
	Delete(ctx context.Context, id uint) error
}

func NewProductService(logger *zap.Logger, productDB repositories.ProductDB) ProductService {
	return &productService{productDB: productDB}
}

type productService struct {
	productDB repositories.ProductDB
}

func (productService *productService) GetAll(ctx context.Context) ([]*dto.ProductDTO, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("[productService]", "GetAll", "")
	products, err := productService.productDB.GetAll(ctx)
	if err != nil {
		logger.Infow("[productService]", "GetAll", "cannot get all products due to ", err.Error())
		return nil, errors.New("cannot get all products due to " + err.Error())
	}
	productDTOs := []*dto.ProductDTO{}
	for _, p := range products {
		productDTO := dto.ProductDTO{ID: p.ID, Name: p.Name, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
		productDTOs = append(productDTOs, &productDTO)
	}
	return productDTOs, nil
}

func (productService *productService) Add(ctx context.Context, productDTO dto.ProductDTO) (*dto.ProductDTO, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("[productService]", "Add", "")
	product, err := productService.productDB.Add(ctx, model.Product{Name: productDTO.Name})
	if err != nil {
		logger.Infow("[productService]", "Add", "cannot save product due to ", err.Error())
		return nil, errors.New("cannot save product due to " + err.Error())
	}

	savedProduct := dto.ProductDTO{ID: product.ID, Name: product.Name, CreatedAt: product.CreatedAt, UpdatedAt: product.UpdatedAt}
	return &savedProduct, nil
}
func (productService *productService) Delete(ctx context.Context, id uint) error {
	logger := logging.FromContext(ctx)
	logger.Infow("[productService]", "Delete", "")
	err := productService.productDB.Delete(ctx, id)
	if err != nil {
		logger.Infow("[productService]", "Delete", "cannot delete product due to ", err.Error())
		return errors.New("cannot delete product due to " + err.Error())
	}
	return nil
}
