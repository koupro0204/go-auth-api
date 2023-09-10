package usecase

import (
	"go-auth-api/domain/entity"
	"go-auth-api/domain/repository"
	myerror "go-auth-api/error"
)

// interfaceから呼び出せるように
type Product interface {
	InsertProduct(name, number string) (*entity.Product, error)
	SelectProduct(name, number string) (*entity.Product, error)
}

type product struct {
	productRepository repository.Product
}

// Productデータに対するusecaseを生成(依存関係の注入用)
func NewProductUseCase(pr repository.Product) Product {
	return &product{
		productRepository: pr,
	}
}
func (p product) InsertProduct(name, number string) (*entity.Product, error) {
	// データベースにProductデータを登録する
	product, err := p.productRepository.InsertProduct(&entity.Product{
		Name:   name,
		Number: number,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}
func (p product) SelectProduct(name, number string) (*entity.Product, error) {
	// データベースからProductデータを検索
	product, err := p.productRepository.SelectProductByNumber(number)
	if err != nil {
		return nil, err
	}
	// DB上のnameと受け取ったNameがあっているか
	if name != product.Name {
		return nil, myerror.New(myerror.ErrorValidation, "invalid name")
	}

	return product, nil
}
