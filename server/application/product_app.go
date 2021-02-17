package application

import (
	"github.com/Prashant-sharma3012/ecom/tree/main/server/domain/entity"
	"github.com/Prashant-sharma3012/ecom/tree/main/server/domain/repository"
)

type ProductApp struct {
	pr repository.ProductRepository
}

var _ ProductAppInterface = &ProductApp{}

type ProductAppInterface interface {
	SaveProduct(*entity.Product) (*entity.Product, error)
	GetAllProduct(int64, int64) ([]entity.Product, error)
	GetProduct(string) (*entity.Product, error)
	UpdateProduct(*entity.Product) ([]byte, error)
	DeleteProduct(string) ([]byte, error)
}

func (p *ProductApp) SaveProduct(Product *entity.Product) (*entity.Product, error) {
	return p.pr.SaveProduct(Product)
}

func (p *ProductApp) GetAllProduct(skip, limit int64) ([]entity.Product, error) {
	return p.pr.GetAllProduct(skip, limit)
}

func (p *ProductApp) GetProduct(ProductId string) (*entity.Product, error) {
	return p.pr.GetProduct(ProductId)
}

func (p *ProductApp) UpdateProduct(Product *entity.Product) ([]byte, error) {
	return p.pr.UpdateProduct(Product)
}

func (p *ProductApp) DeleteProduct(ProductId string) ([]byte, error) {
	return p.pr.DeleteProduct(ProductId)
}
