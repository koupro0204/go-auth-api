package persistence

import (
	"database/sql"
	"go-auth-api/domain/entity"
	"go-auth-api/domain/repository"
	myerror "go-auth-api/error"
)

type userProductPersistence struct {
	DB *sql.DB
}

func NewUserProductPersistence(DB *sql.DB) repository.UserProduct {
	return &userProductPersistence{DB: DB}
}

func (up userProductPersistence) InsertUserProduct(userProduct *entity.UserProduct) (*entity.UserProduct, error) {
	if _, err := up.DB.Exec(
		"INSERT INTO UserProduct (user_id, product_id) values (?, ?)",
		userProduct.UserID,
		userProduct.ProductID,
	); err != nil {
		return nil, myerror.New(myerror.ErrorInsert, "Failed UserProduct insert")
	}
	return userProduct, nil
}

// SelectProductByNumber numberを条件にレコードを取得する
func (up userProductPersistence) SelectUserProduct(userID, productID int) error {
	// passwordが正しいかはusecaseで行う
	row := up.DB.QueryRow("SELECT * FROM UserProduct WHERE user_id=? AND product_id=?", userID, productID)
	_, err := convertToUserProduct(row)
	if err != nil {
		return err
	}
	return nil
}

// convertToUser rowデータをUserデータへ変換する
func convertToUserProduct(row *sql.Row) (*entity.UserProduct, error) {
	var userProduct entity.UserProduct
	if err := row.Scan(&userProduct.UserProductID, &userProduct.UserID, &userProduct.ProductID); err != nil {
		if err == sql.ErrNoRows {
			return nil, myerror.New(myerror.ErrorNotLinked, "User and product are not linked")
		}
		return nil, myerror.New(myerror.ErrorDB, "Failed convert to userProduct")
	}
	return &userProduct, nil
}
