package repository

import (
	"context"
	"rest_api/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}
