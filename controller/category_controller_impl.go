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

func (c *CategoryControllerImpl) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryCreateRequest := web.CategoryCreateRequest{}
		web.ReadFromRequestBody(r, &categoryCreateRequest)

		err := c.Validate.Struct(categoryCreateRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   400,
				Status: "Bad Request",
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryResponse, err := c.CategoryService.Create(r.Context(), categoryCreateRequest)
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(utils.GetCode(err))
			w.Write([]byte(utils.GetMessage(err)))
			return
		}

		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponse,
		}

		web.WriteToResponseBody(w, webResponse)

	})
}

func (c *CategoryControllerImpl) Update() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		categoryUpdateRequest := web.CategoryUpdateRequest{}
		web.ReadFromRequestBody(r, &categoryUpdateRequest)
		err := c.Validate.Struct(categoryUpdateRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   400,
				Status: "Bad Request",
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(categoryId)
		if err != nil {
			utils.NotFoundError(err)
			return
		}

		categoryResponseService, err := c.CategoryService.FindById(r.Context(), id)
		fmt.Println("response", categoryResponseService)
		if err != nil {
			w.WriteHeader(utils.GetCode(err))
			erorResponse := web.ErrorResponse{
				Code:   404,
				Status: "Not Found",
				Errors: []web.WebError{
					{
						Message: "Data not found",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryUpdateRequest.Id = id
		web.ReadFromRequestBody(r, &categoryUpdateRequest)

		categoryResponse, err := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
		if err != nil {

			w.WriteHeader(utils.GetCode(err))
			w.Write([]byte(utils.GetMessage(err)))
			return
		}
		categoryResponse.Id = id
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponse,
		}
		web.WriteToResponseBody(w, webResponse)

	})
}

func (c *CategoryControllerImpl) FindById() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		categoryId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(categoryId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			erorResponse := web.ErrorResponse{
				Code:   404,
				Status: "Not Found",
				Errors: []web.WebError{
					{
						Message: "Data not found",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryResponse, err := c.CategoryService.FindById(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			erorResponse := web.ErrorResponse{
				Code:   404,
				Status: "Not Found",
				Errors: []web.WebError{
					{
						Message: "Data not found",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponse,
		}

		web.WriteToResponseBody(w, webResponse)

	})
}

func (c *CategoryControllerImpl) Delete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		categoryId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(categoryId)
		if err != nil {
			// utils.NotFoundError(err)
			w.WriteHeader(http.StatusNotFound)
			erorResponse := web.ErrorResponse{
				Code:   404,
				Status: "Not Found",
				Errors: []web.WebError{
					{
						Message: "Data not found",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		resDelete, err := c.CategoryService.Delete(r.Context(), id)
		if err != nil || resDelete == "" {
			w.WriteHeader(http.StatusNotFound)
			erorResponse := web.ErrorResponse{
				Code:   404,
				Status: "Not Found",
				Errors: []web.WebError{
					{
						Message: "Data not found",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   "Data Deleted",
		}

		web.WriteToResponseBody(w, webResponse)

	})
}

func (c *CategoryControllerImpl) FindAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") != "" {
			par := r.URL.Query().Get("page")
			page, err := strconv.Atoi(par)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				webResponse := web.ErrorResponse{
					Code:   404,
					Status: "Not Found",
					Errors: []web.WebError{
						{
							Message: "Data not found",
						},
					},
				}
				web.WriteToResponseBody(w, webResponse)

				return
			}
			categoryResponses, meta, err := c.CategoryService.FindAll(r.Context(), page)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				webResponse := web.ErrorResponse{
					Code:   404,
					Status: "Not Found",
					Errors: []web.WebError{
						{
							Message: "Data not found",
						},
					},
				}
				web.WriteToResponseBody(w, webResponse)
				return
			}
			webResponse := web.GetAllCategory{
				Code:     200,
				Status:   "OK",
				Data:     categoryResponses,
				MetaData: meta,
			}

			web.WriteToResponseBody(w, webResponse)

		} else {
			categoryResponses, meta, err := c.CategoryService.FindAll(r.Context(), 0)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				webResponse := web.ErrorResponse{
					Code:   404,
					Status: "Not Found",
					Errors: []web.WebError{
						{
							Message: "Data not found",
						},
					},
				}
				web.WriteToResponseBody(w, webResponse)
				return
			}
			webResponse := web.GetAllCategory{
				Code:     200,
				Status:   "OK",
				Data:     categoryResponses,
				MetaData: meta,
			}

			web.WriteToResponseBody(w, webResponse)
		}

	})
}
