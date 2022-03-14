package controller

import "net/http"

type CategoryController interface {
	FindById() http.HandlerFunc

	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	FindAll() http.HandlerFunc
}
