package repository

import (
	"context"
	"rest_api/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	// DB *sqlx.DB
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	rs := model.User{}
	err := r.DB.Where("username = ?", username).First(&rs).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil

	// query := "SELECT id,username,password FROM users WHERE username=?"
	// rs := model.User{}
	// err := r.DB.Get(&rs, query, username)
	// if err != nil {
	// 	return nil, err
	// }

	// return &rs, nil
}
