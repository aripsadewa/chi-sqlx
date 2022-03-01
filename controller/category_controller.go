package controller

import "net/http"

type CategoryController interface {
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	FindById() http.HandlerFunc
	FindAll() http.HandlerFunc
}
