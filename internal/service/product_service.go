package service

import (
	"context"
	"log"

	"github.com/farizziezhi/layered/internal/domain"
	"github.com/farizziezhi/layered/internal/repository"
)

type ProductService interface {
	GetProducts(ctx context.Context) ([]domain.Product, error)
	GetProduct(ctx context.Context, id string) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product) error
	UpdateProduct(ctx context.Context, id string, product domain.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type ProductServiceImpl struct {
	Repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{Repo: repo}
}

func (s *ProductServiceImpl) GetProducts(ctx context.Context) ([]domain.Product, error) {
	result, err := s.Repo.GetProducts(ctx)

	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}
		
	return result, nil
}

func (s *ProductServiceImpl) GetProduct(ctx context.Context, id string) (domain.Product, error) {
	result, err := s.Repo.GetProduct(ctx, id)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}
	return result, nil
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, product domain.Product) error {
	err := s.Repo.CreateProduct(ctx, product)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return err
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, id string, product domain.Product) error {
	err := s.Repo.UpdateProduct(ctx, id, product)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return err
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	err := s.Repo.DeleteProduct(ctx, id)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return err
}