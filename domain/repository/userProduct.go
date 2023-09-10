package repository

import (
	"go-auth-api/domain/entity"
)

type UserProduct interface {
	InsertUserProduct(userProduct *entity.UserProduct) (*entity.UserProduct, error)
	SelectUserProduct(userID, productID int) error
}
