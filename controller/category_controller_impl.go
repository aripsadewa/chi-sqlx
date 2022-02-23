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

func (c *CategoryControllerImpl) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryCreateRequest := web.CategoryCreateRequest{}

		web.ReadFromRequestBody(r, &categoryCreateRequest)

		err := c.Validate.Struct(categoryCreateRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
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
			erorResponse := web.ErrorResponse{
				Code:   utils.GetCode(err),
				Status: http.StatusText(utils.GetCode(err)),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
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
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
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
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: []web.WebError{
					{
						Message: "param not int",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryResponseService, err := c.CategoryService.FindById(r.Context(), id)
		fmt.Println("response", categoryResponseService)
		if err != nil {
			w.WriteHeader(utils.GetCode(err))
			erorResponse := web.ErrorResponse{
				Code:   utils.GetCode(err),
				Status: http.StatusText(utils.GetCode(err)),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryUpdateRequest.Id = id
		categoryResponse, err := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
		if err != nil {
			w.WriteHeader(utils.GetCode(err))
			erorResponse := web.ErrorResponse{
				Code:   utils.GetCode(err),
				Status: http.StatusText(utils.GetCode(err)),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		categoryResponse.Id = id
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
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
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: []web.WebError{
					{
						Message: "Param not int",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		categoryResponse, err := c.CategoryService.FindById(r.Context(), id)
		if err != nil {
			w.WriteHeader(utils.GetCode(err))
			erorResponse := web.ErrorResponse{
				Code:   utils.GetCode(err),
				Status: http.StatusText(utils.GetCode(err)),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
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
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: []web.WebError{
					{
						Message: "Param not int",
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}

		resDelete, err := c.CategoryService.Delete(r.Context(), id)
		if err != nil || resDelete == "" {
			w.WriteHeader(utils.GetCode(err))
			erorResponse := web.ErrorResponse{
				Code:   utils.GetCode(err),
				Status: http.StatusText(utils.GetCode(err)),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   "Data Deleted",
		}

		web.WriteToResponseBody(w, webResponse)

	})
}

func (c *CategoryControllerImpl) FindAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decoder = schema.NewDecoder()
		paramRequest := web.GetParamRequest{}
		err := decoder.Decode(&paramRequest, r.URL.Query())
		// r.Body.Read()
		// fmt.Printf("value %s\n", reflect.TypeOf(r.URL.Query()))
		// fmt.Printf("param %#v\n", paramRequest)
		// fmt.Printf("param %v\n", paramRequest)
		// fmt.Printf("param %+v\n", paramRequest)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		err = c.Validate.Struct(paramRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			erorResponse := web.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		categoryResponses, meta, err := c.CategoryService.FindAll(r.Context(), paramRequest)
		if err != nil {
			w.WriteHeader(utils.GetCode(err))
			erorResponse := web.ErrorResponse{
				Code:   utils.GetCode(err),
				Status: http.StatusText(utils.GetCode(err)),
				Errors: []web.WebError{
					{
						Message: utils.GetMessage(err),
					},
				},
			}
			web.WriteToResponseBody(w, erorResponse)
			return
		}
		webResponse := web.GetAllCategory{
			Code:     http.StatusOK,
			Status:   http.StatusText(http.StatusOK),
			Data:     categoryResponses,
			MetaData: meta,
		}

		web.WriteToResponseBody(w, webResponse)

	})
}
