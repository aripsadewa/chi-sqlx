package service

import (
	"context"
	"rest_api/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) (*web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (*web.CategoryResponse, error)
	FindById(ctx context.Context, categoryId int) (*web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId int) (string, error)
	FindAll(ctx context.Context, request web.ParamRequest) ([]*web.CategoryResponse, *web.MetaData, error)
}
