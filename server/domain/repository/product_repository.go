package repository

import "github.com/Prashant-sharma3012/ecom/tree/main/server/domain/entity"

type ProductRepository interface {
	SaveProduct(*entity.Product) (*entity.Product, error)
	GetProduct(string) (*entity.Product, error)
	GetAllProduct(int64, int64) ([]entity.Product, error)
	UpdateProduct(*entity.Product) ([]byte, error)
	DeleteProduct(string) ([]byte, error)
}
