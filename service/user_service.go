package service

import (
	"context"
	"rest_api/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserCreateRequest) (*web.UserResponse, error)
	Login(ctx context.Context, request web.UserCreateRequest) (*web.LoginResponse, error)
}
