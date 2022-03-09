package controller

import (
	"fmt"
	"net/http"
	"rest_api/service"
	"rest_api/utils"
	"rest_api/web"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the input paylod
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body web.UserCreateRequest true "Create Category"
// @Router /register [post]
func (c *UserControllerImpl) Register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCreateRequest := web.UserCreateRequest{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		// web.ReadFromRequestBody(r, &userCreateRequest)
		fmt.Printf("request %+v \n", userCreateRequest)

		userResponse, err := c.UserService.Register(r.Context(), userCreateRequest)
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
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), userResponse, nil, nil)

	})
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the input paylod
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body web.UserCreateRequest true "Create Category"
// @Router /login [post]
func (c *UserControllerImpl) Login() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCreateRequest := web.UserCreateRequest{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		// web.ReadFromRequestBody(r, &userCreateRequest)

		userResponse, err := c.UserService.Login(r.Context(), userCreateRequest)
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
		web.WriteToResponseBody(w, http.StatusOK, http.StatusText(http.StatusOK), userResponse, nil, nil)

	})
}
