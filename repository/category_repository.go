package repository

import (
	"context"
	"rest_api/model/domain"
	"rest_api/web"
)

type CategoryRepository interface {
	Save(ctx context.Context, category domain.Category) (*domain.Category, error)
	Update(ctx context.Context, category domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, categoryId int) (int, error)
	FindById(ctx context.Context, categoryId int) (*domain.Category, error)
	FindAll(ctx context.Context, request web.GetParamRequest) ([]*domain.Category, *domain.CategoryMeta, error)
}
