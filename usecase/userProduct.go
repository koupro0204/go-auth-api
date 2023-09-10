package usecase

import (
	"go-auth-api/domain/entity"
	"go-auth-api/domain/repository"
	myerror "go-auth-api/error"
)

// interfaceから呼び出せるように
type UserProduct interface {
	InsertUserProduct(email, password, name, number string) (*entity.UserProduct, error)
	SelectUserProduct(email, password, name, number string) error
}

type userProduct struct {
	userRepository        repository.User
	productRepository     repository.Product
	userProductRepository repository.UserProduct
}

// Productデータに対するusecaseを生成(依存関係の注入用)
func NewUserProductUseCase(ur repository.User, pr repository.Product, pur repository.UserProduct) UserProduct {
	return &userProduct{
		userRepository:        ur,
		productRepository:     pr,
		userProductRepository: pur,
	}
}
func (up userProduct) InsertUserProduct(email, password, name, number string) (*entity.UserProduct, error) {
	// Userが存在するか
	// データベースでユーザデータを検索
	user, err := up.userRepository.SelectUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// passwordが正しいか判断
	if password != user.Password {
		return nil, myerror.New(myerror.ErrorValidationPassword, "invalid password")
	}
	// Productが存在するか
	// データベースからProductデータを検索
	product, err := up.productRepository.SelectProductByNumber(number)
	if err != nil {
		return nil, err
	}
	// DB上のnameと受け取ったNameがあっているか
	if name != product.Name {
		return nil, myerror.New(myerror.ErrorValidationName, "invalid name")
	}
	// データベースにUserProductデータを登録する
	userProduct, err := up.userProductRepository.InsertUserProduct(&entity.UserProduct{
		UserID:    user.UserID,
		ProductID: product.ProductID,
	})
	if err != nil {
		return nil, err
	}

	return userProduct, nil
}

func (up userProduct) SelectUserProduct(email, password, name, number string) error {
	// Userが存在するか
	// データベースでユーザデータを検索
	user, err := up.userRepository.SelectUserByEmail(email)
	if err != nil {
		return err
	}
	// passwordが正しいか判断
	if password != user.Password {
		return myerror.New(myerror.ErrorValidationPassword, "invalid password")
	}
	// Productが存在するか
	// データベースからProductデータを検索
	product, err := up.productRepository.SelectProductByNumber(number)
	if err != nil {
		return err
	}
	// DB上のnameと受け取ったNameがあっているか
	if name != product.Name {
		return myerror.New(myerror.ErrorValidationName, "invalid name")
	}
	// DBに登録されている確認
	if err := up.userProductRepository.SelectUserProduct(user.UserID, product.ProductID); err != nil {
		return err
	}
	return nil
}

// user,productの存在確認は全く同じだからまとめたほうがいいかも
