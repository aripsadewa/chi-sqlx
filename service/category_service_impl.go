package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"rest_api/model/domain"
	"rest_api/repository"
	"rest_api/utils"
	"rest_api/web"
	"sync"
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
		Description: sql.NullString{
			Valid:  request.Description != "",
			String: request.Description,
		},
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
		ID:   request.Id,
		Name: request.Name,
		Description: sql.NullString{
			Valid:  request.Description != "",
			String: request.Description,
		},
	}
	categories, err := s.CategoryRepository.Update(ctx, category)
	if err != nil {
		return nil, utils.NotFoundError(err)
	}
	res := web.ToCategoryResponse(*categories)
	return res, nil
}

func (s *CategoryServiceImpl) FindData(ctx context.Context, request web.GetParamRequest) ([]*web.CategoryResponse, *web.PaginateMetaData, error) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	chanelCategory := make(chan []*domain.Category)
	chanelCountCategory := make(chan int)
	done := make(chan bool)
	chanelErr := make(chan error)

	defer close(chanelCategory)
	defer close(chanelCountCategory)
	defer close(done)
	defer close(chanelErr)

	go func() {
		wg.Wait()
		done <- true
	}()

	go s.findAll(ctx, request, wg, chanelCategory, chanelErr)

	filterPayload, _ := s.getFilterPayload(request)

	go s.getMetaCategoryService(filterPayload, request, wg, chanelCountCategory, chanelErr)

	categoriesData := make([]*domain.Category, 0)
	countData := 0
	var err error
L:
	for {
		select {
		case cData := <-chanelCategory:
			categoriesData = cData
		case coData := <-chanelCountCategory:
			countData = coData
		case coErr := <-chanelErr:
			err = coErr
		case <-done:
			break L
		}
	}

	if err != nil {
		return nil, nil, utils.InternalServerError(err)
	}
	resCategories := web.ToCategoriesResponse(categoriesData)

	resPaginateMetadata := web.PaginateMetaData{
		Page:      float64(utils.CekNulNumberRequest(request.Page.Int64, 1)),
		Limit:     float64(utils.CekNulNumberRequest(request.Limit.Int64, 5)),
		TotalPage: int(math.Ceil(float64(countData) / float64(utils.CekNulNumberRequest(request.Limit.Int64, 5)))),
		Total:     countData,
	}

	return resCategories, &resPaginateMetadata, nil
}

func (s *CategoryServiceImpl) findAll(ctx context.Context, request web.GetParamRequest, wg *sync.WaitGroup, chanelCategory chan []*domain.Category, chanErr chan error) {
	defer wg.Done()

	filterPayload, paginateParam := s.getFilterPayload(request)
	categories, _ := s.CategoryRepository.FindData(ctx, filterPayload, paginateParam)
	chanErr <- errors.New("error coy")

	// if err != nil {
	// }

	chanelCategory <- categories

}

func (s *CategoryServiceImpl) getMetaCategoryService(cty domain.CategoryFilter, request web.GetParamRequest, wg *sync.WaitGroup, chanelCount chan int, chanErr chan error) {
	defer wg.Done()

	count, err := s.CategoryRepository.GetCountCategory(cty)
	if err != nil {
		chanErr <- err

	}

	chanelCount <- count

}

func (s *CategoryServiceImpl) getFilterPayload(request web.GetParamRequest) (domain.CategoryFilter, *web.PaginateMetaData) {
	filterPayload := domain.CategoryFilter{
		StartDate: request.Start,
		EndDate:   request.End,
		Name:      request.Name,
	}
	filterPayload.SortValue = utils.CekNilParameter(request.SortValue.String, utils.EnvConfigs.SortCategoryValue)
	filterPayload.Sort = utils.CekNilParameter(request.Sort.String, "id")

	paginateParam := &web.PaginateMetaData{}

	paginateParam.Offset = int(utils.CekNulNumberRequest(request.Page.Int64, 1)-1) * int(utils.CekNulNumberRequest(request.Limit.Int64, 5))
	paginateParam.Page = float64(utils.CekNulNumberRequest(request.Page.Int64, 1))
	paginateParam.Limit = float64(utils.CekNulNumberRequest(request.Limit.Int64, 5))
	fmt.Printf("meta %+v ", paginateParam)

	return filterPayload, paginateParam
}
