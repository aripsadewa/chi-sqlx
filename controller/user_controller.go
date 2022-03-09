package controller

import "net/http"

type UserController interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
}
