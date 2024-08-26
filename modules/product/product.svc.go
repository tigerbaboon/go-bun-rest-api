package product

import (
	"app/models"
	productdto "app/modules/product/dto"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type ProductService struct {
	db *bun.DB
}

func newService(db *bun.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

func (s *ProductService) Create(ctx context.Context, req productdto.CreateProductRequest) (*models.Product, error) {

	m := models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		IsActived:   req.IsActived,
	}

	_, err := s.db.NewInsert().Model(&m).Exec(ctx)

	return &m, err
}

func (s *ProductService) Update(ctx context.Context, id productdto.GetProductByIDRequest, req productdto.UpdateProductRequest) (*models.Product, error) {

	ex, err := s.db.NewSelect().Model((*models.Product)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("Product not found")
	}

	m := models.Product{
		ID:          id.ID,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		IsActived:   req.IsActived,
	}

	_, err = s.db.NewUpdate().Model(&m).
		Set("name = ?name").
		Set("name = ?name").
		Set("price = ?price").
		Set("description = ?description").
		Set("is_actived = ?is_actived").
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	return &m, err
}

func (s *ProductService) Delete(ctx context.Context, id productdto.GetProductByIDRequest) (*models.Product, error) {

	ex, err := s.db.NewSelect().Model((*models.Product)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("Product not found")
	}

	_, err = s.db.NewDelete().Model((*models.Product)(nil)).Where("id = ?", id.ID).Exec(ctx)
	return nil, err
}

func (s *ProductService) Get(ctx context.Context, id productdto.GetProductByIDRequest) (*models.Product, error) {

	m := models.Product{}

	err := s.db.NewSelect().Model(&m).Where("id = ?", id.ID).Scan(ctx)

	return &m, err
}

func (s *ProductService) List(ctx context.Context) ([]models.Product, error) {

	m := []models.Product{}
	err := s.db.NewSelect().Model(&m).Scan(ctx)

	return m, err
}
