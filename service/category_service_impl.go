package service

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/model/domain"
	"rest_api/repository"
	"rest_api/utils"
	"rest_api/web"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) (*web.CategoryResponse, error) {
	// return nil, utils.BadRequest(errors.New("test error"))

	category := domain.Category{
		Name: request.Name,
	}

	categories, err := s.CategoryRepository.Save(ctx, category)
	if err != nil {
		return nil, utils.UnprocessableEntity(err)
	}
	res := web.ToCategoryResponse(*categories)
	fmt.Println(res)
	return res, nil
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (*web.CategoryResponse, error) {
	category, err := s.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NotFoundError(err)
		}
		return nil, utils.InternalServerError(err)

	}

	res := web.ToCategoryResponse(*category)
	return res, nil
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) (string, error) {
	status, err := s.CategoryRepository.Delete(ctx, categoryId)
	if err != nil || int(status) == 0 {

		return "", utils.NotFoundError(err)
	}
	mes := fmt.Sprintf("%d not found", categoryId)
	return mes, nil

}

func (s *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (*web.CategoryResponse, error) {
	category := domain.Category{
		ID:   request.Id,
		Name: request.Name,
	}
	categories, err := s.CategoryRepository.Update(ctx, category)

	if err != nil {

		return nil, utils.UnprocessableEntity(err)
	}

	res := web.ToCategoryResponse(*categories)
	return res, nil
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context, request web.GetParamRequest) ([]*web.CategoryResponse, *web.MetaData, error) {
	// param := domain.CategoryMeta{
	// 	Start:     request.Start,
	// 	End:       request.End,
	// 	Page:      request.Page,
	// 	Limit:     request.Limit,
	// 	Sort:      request.Sort,
	// 	SortValue: request.SortValue,
	// }
	categories, metaData, err := s.CategoryRepository.FindAll(ctx, request)
	if err != nil || metaData == nil {
		return nil, nil, utils.NotFoundError(err)
	}
	res := web.ToCategoryMeta(*metaData)

	return web.AllCategoryResponse(categories), res, nil
}
