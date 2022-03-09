package repository

import (
	"context"
	"fmt"
	"rest_api/model"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	query := "INSERT INTO users (username,password, created_at) VALUES (:username,:password, now())"
	fmt.Printf("repo %+v \n", user)

	rs, err := r.DB.NamedExec(query, user)

	if err != nil {
		return nil, err
	}

	insertId, err := rs.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(insertId)

	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT id,username,password FROM users WHERE username=?"
	rs := model.User{}
	err := r.DB.Get(&rs, query, username)
	if err != nil {
		return nil, err
	}

	return &rs, nil
}
