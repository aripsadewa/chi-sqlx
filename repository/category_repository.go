package repository

import (
	"context"
	"rest_api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, category domain.Category) (*domain.Category, error)
	Update(ctx context.Context, category domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, categoryId int) (int, error)
	FindById(ctx context.Context, categoryId int) (*domain.Category, error)
	FindAll(ctx context.Context, meta domain.CategoryMeta) ([]*domain.Category, *domain.CategoryMeta, error)
}
