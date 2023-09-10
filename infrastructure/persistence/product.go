package persistence

import (
	"database/sql"
	"go-auth-api/domain/entity"
	"go-auth-api/domain/repository"
	myerror "go-auth-api/error"
)

type productPersistence struct {
	DB *sql.DB
}

func NewProductPersistence(DB *sql.DB) repository.Product {
	return &productPersistence{DB: DB}
}

func (p productPersistence) InsertProduct(product *entity.Product) (*entity.Product, error) {
	if _, err := p.DB.Exec(
		"INSERT INTO Product (product_name, product_number) values (?, ?)",
		product.Name,
		product.Number,
	); err != nil {
		return nil, myerror.New(myerror.ErrorDB, "Failed product insert")
	}
	return product, nil
}

// SelectProductByNumber numberを条件にレコードを取得する
func (p productPersistence) SelectProductByNumber(number string) (*entity.Product, error) {
	// passwordが正しいかはusecaseで行う
	row := p.DB.QueryRow("SELECT * FROM Product WHERE product_number=?", number)
	return convertToProduct(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToProduct(row *sql.Row) (*entity.Product, error) {
	var product entity.Product
	if err := row.Scan(&product.ProductID, &product.Name, &product.Number); err != nil {
		if err == sql.ErrNoRows {
			return nil, myerror.New(myerror.ErrorDB, "Product is not exist")
		}
		return nil, myerror.New(myerror.ErrorDB, "Failed convert to Product")
	}
	return &product, nil
}
