package app

import (
	"net/http"
	"rest_api/controller"
	"rest_api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewCategoryRouter(r chi.Router, cat controller.CategoryController) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome v1"))
	})

	r.Route("/category", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(utils.TokenVerify)
		r.Post("/create", cat.Create())
		r.Get("/{id}", cat.FindById())
		r.Put("/{id}", cat.Update())
		r.Delete("/{id}", cat.Delete())
		r.Get("/", cat.FindAll())
	})

}

func NewUserRouter(r chi.Router, user controller.UserController) {
	r.Post("/register", user.Register())
	r.Post("/login", user.Login())
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome user"))
	})

}
