package service

import (
	"context"
	"fmt"
	"rest_api/helpers"
	"rest_api/model"
	"rest_api/repository"
	"rest_api/utils"
	"rest_api/web"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) (*web.UserResponse, error) {
	hash, _ := helpers.HashPassword(request.Password)

	user := model.User{
		Username: request.Username,
		Password: hash,
	}

	users, err := s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, utils.UnprocessableEntity(err)
	}
	res := web.ToUserResponse(*users)
	return res, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, request web.UserCreateRequest) (*web.LoginResponse, error) {
	user, err := s.UserRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return nil, utils.NotFoundError(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println("error ", err)
		return nil, err
	}
	data := map[string]string{
		"role":     fmt.Sprint(user.ID),
		"username": user.Username,
	}
	token, err := utils.GenerateToken(utils.EnvConfigs.SecretApp, utils.EnvConfigs.ExpToken, data)
	if err != nil {
		return nil, err
	}
	result := &web.LoginResponse{
		Username: user.Username,
		Expired:  utils.EnvConfigs.ExpToken.String(),
		Token:    *token,
	}
	return result, nil
}
