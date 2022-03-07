package app

import (
	"net/http"
	"rest_api/controller"

	"github.com/go-chi/chi/v5"
)

func NewRouter(r chi.Router, cat controller.CategoryController) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome v1"))
	})
	r.Route("/category", func(r chi.Router) {
		r.Post("/create", cat.Create())
		r.Get("/{id}", cat.FindById())
		r.Put("/{id}", cat.Update())
		r.Delete("/{id}", cat.Delete())
		r.Get("/", cat.FindAll())
	})

}
