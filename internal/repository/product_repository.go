package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/farizziezhi/layered/internal/domain"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]domain.Product, error)
	GetProduct(ctx context.Context, id string) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product) error
	UpdateProduct(ctx context.Context, id string, product domain.Product) error	
	DeleteProduct(ctx context.Context, id string) error
}

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (repo *ProductRepositoryImpl) GetProducts(ctx context.Context) ([]domain.Product, error) {
	result := []domain.Product{}
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, name, price FROM products")

	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		data := domain.Product{}
		err := rows.Scan(&data.ID, &data.Name, &data.Price)
		if err != nil {
			log.Println("ERROR:", err)
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (repo *ProductRepositoryImpl) GetProduct(ctx context.Context, id string) (domain.Product, error) {
	var result domain.Product
	err := repo.DB.QueryRowContext(ctx, "SELECT id, name, price FROM products WHERE id = ?", id).Scan(&result.ID, &result.Name, &result.Price)
	if err != nil {
		log.Println("ERROR:", err)
		return result, err
	}

	return result, nil
}

func (repo *ProductRepositoryImpl) CreateProduct(ctx context.Context, product domain.Product) error {
	_, err := repo.DB.ExecContext(ctx, "INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	return nil
}

func (repo *ProductRepositoryImpl) UpdateProduct(ctx context.Context, id string, product domain.Product) error {
	res, err := repo.DB.ExecContext(ctx, "UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, id)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepositoryImpl) DeleteProduct(ctx context.Context, id string) error {
	res, err := repo.DB.ExecContext(ctx, "DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}