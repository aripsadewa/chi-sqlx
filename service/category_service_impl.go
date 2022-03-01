package service

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"rest_api/model"
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
		Name:        request.Name,
		Description: request.Description,
	}

	categories, err := s.CategoryRepository.Save(ctx, category)
	if err != nil {
		return nil, utils.UnprocessableEntity(err)
	}
	res := web.ToCategoryResponse(*categories)
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
	_, err := s.CategoryRepository.Delete(ctx, categoryId)
	if err != nil {

		return "", utils.NotFoundError(err)
	}
	mes := fmt.Sprintf("category id %d is Deleted", categoryId)
	return mes, nil
}

func (s *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (*web.CategoryResponse, error) {
	category := domain.Category{
		ID:          request.Id,
		Name:        request.Name,
		Description: request.Description,
	}
	categories, err := s.CategoryRepository.Update(ctx, category)
	if err != nil {
		return nil, utils.NotFoundError(err)
	}
	res := web.ToCategoryResponse(*categories)
	return res, nil
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context, request web.GetParamRequest) ([]*web.CategoryResponse, *web.PaginateMetaData, error) {
	fmt.Printf("service %+v \n", request)

	filterPayload := domain.CategoryFilter{
		StartDate: request.Start,
		EndDate:   request.End,
		Name:      request.Name,
	}
	filterPayload.SortValue = utils.CekNilParameter(request.SortValue.String, utils.EnvConfigs.SortCategoryValue)
	filterPayload.Sort = utils.CekNilParameter(request.Sort.String, "id")

	paginateParam := model.PaginateParams{
		Offset: int(utils.CekNulNumberRequest(request.Page.Int64, 1)-1) * int(utils.CekNulNumberRequest(request.Limit.Int64, 5)),
		Limit:  int(utils.CekNulNumberRequest(request.Limit.Int64, 5)),
	}
	categories, meta, err := s.CategoryRepository.FindData(ctx, filterPayload, paginateParam)
	if err != nil {
		return nil, nil, utils.NotFoundError(err)
	}
	resCategories := web.ToCategoriesResponse(categories)

	resPaginateMetadata := web.PaginateMetaData{
		Page:      float64(utils.CekNulNumberRequest(request.Page.Int64, 1)),
		Limit:     float64(utils.CekNulNumberRequest(request.Limit.Int64, 5)),
		TotalPage: int(math.Ceil(float64(meta.Total) / float64(paginateParam.Limit))),
		Total:     meta.Total,
	}
	return resCategories, &resPaginateMetadata, nil
}
