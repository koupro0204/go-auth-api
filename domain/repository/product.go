package repository

import (
	"go-auth-api/domain/entity"
)

type Product interface {
	InsertProduct(product *entity.Product) (*entity.Product, error)
	SelectProductByNumber(number string) (*entity.Product, error)
}
