package controller

import (
	"fmt"
	"net/http"
	"rest_api/service"
	"rest_api/utils"
	"rest_api/web"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
	Validate        *validator.Validate
}

func NewCategoryController(categoryService service.CategoryService, validate *validator.Validate) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
		Validate:        validate,
	}
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category with the input paylod
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body web.CategoryUpdateRequest true "Update Category"
// @Router /category/{id} [put]
func (c *CategoryControllerImpl) Update() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		categoryUpdateRequest := web.CategoryUpdateRequest{}
		web.ReadFromRequestBody(r, &categoryUpdateRequest)
		err := c.Validate.Struct(categoryUpdateRequest)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			web.WriteToResponseBody(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, erorResponse, nil)
			return
		}

		categoryId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(categoryId)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: "param is not int",
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}

		categoryUpdateRequest.Id = id
		categoryResponse, err := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}
		categoryResponse.Id = id
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse, nil, nil)

	})
}

// Show Category godoc
// @Summary      Show an Category
// @Description  get string by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Router       /category/{id} [get]
func (c *CategoryControllerImpl) FindById() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		categoryId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(categoryId)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: "param is not int",
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}

		categoryResponse, err := c.CategoryService.FindById(r.Context(), id)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse, nil, nil)
	})
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Create a new category with the input paylod
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Router /category/{id} [delete]
func (c *CategoryControllerImpl) Delete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(categoryId)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: "param is not int",
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}
		resDelete, err := c.CategoryService.Delete(r.Context(), id)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), resDelete, nil, nil)
	})
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the input paylod
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body web.CategoryCreateRequest true "Create Category"
// @Router /category/create [post]
func (c *CategoryControllerImpl) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryCreateRequest := web.CategoryCreateRequest{}
		web.ReadFromRequestBody(r, &categoryCreateRequest)

		err := c.Validate.Struct(categoryCreateRequest)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			web.WriteToResponseBody(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, erorResponse, nil)
			return
		}
		categoryResponse, err := c.CategoryService.Create(r.Context(), categoryCreateRequest)
		fmt.Println(err)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse, nil, nil)

	})
}

// Show Category godoc
// @Summary      Show an Category
// @Description  get categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        name  query string  false  "name"
// @Param        limit query int  false  "limit"
// @Param        page query int  false  "page"
// @Router       /category [get]
func (c *CategoryControllerImpl) FindAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decoder = schema.NewDecoder()
		paramRequest := web.GetParamRequest{}
		err := decoder.Decode(&paramRequest, r.URL.Query())
		fmt.Printf("request %+v \n", paramRequest)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			web.WriteToResponseBody(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, erorResponse, nil)
			return
		}
		err = c.Validate.Struct(paramRequest)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			web.WriteToResponseBody(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, erorResponse, nil)
			return
		}
		categoryResponses, metaData, err := c.CategoryService.FindAll(r.Context(), paramRequest)
		if err != nil {
			erorResponse := []web.WebError{
				{
					Message: utils.GetMessage(err),
				},
			}
			resCode := utils.GetCode(err)
			web.WriteToResponseBody(w, resCode, http.StatusText(resCode), nil, erorResponse, nil)
			return
		}
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), categoryResponses, nil, metaData)

	})
}
