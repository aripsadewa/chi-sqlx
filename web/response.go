package web

import (
	"rest_api/model"
	"rest_api/model/domain"
)

type CategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expired  string `json:"expired"`
}

func ToCategoryResponse(category domain.Category) *CategoryResponse {
	categoryResponse := &CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
	if category.Description.Valid {
		categoryResponse.Description = category.Description.String
	}
	return categoryResponse
}

func ToUserResponse(user model.User) *UserResponse {
	userResponse := &UserResponse{
		Id:       user.ID,
		Username: user.Username,
	}
	return userResponse
}

func ToCategoriesResponse(category []*domain.Category) []*CategoryResponse {
	mapData := make([]*CategoryResponse, 0)
	for _, el := range category {
		responItem := CategoryResponse{
			Id:   el.ID,
			Name: el.Name,
		}
		if el.Description.Valid {
			responItem.Description = el.Description.String
		}
		mapData = append(mapData, &responItem)
	}
	return mapData
}
