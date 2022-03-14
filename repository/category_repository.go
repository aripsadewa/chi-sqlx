package repository

import (
	"context"
	"rest_api/model/domain"
	"rest_api/web"
)

type CategoryRepository interface {
	FindById(ctx context.Context, categoryId int) (*domain.Category, error)

	Save(ctx context.Context, category domain.Category) (*domain.Category, error)
	Update(ctx context.Context, category domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, categoryId int) (int, error)

	FindData(ctx context.Context, filter domain.CategoryFilter, paginate *web.PaginateMetaData) ([]*domain.Category, error)
	GetCountCategory(filter domain.CategoryFilter) (int64, error)

	// FindAll(ctx context.Context, request web.ValidateParamRequest) ([]*domain.Category, *model.PaginateParams, error)
}
