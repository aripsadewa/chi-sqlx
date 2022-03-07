package web

import (
	"rest_api/model/domain"
)

type CategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCategoryResponse(category domain.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func ToCategoriesResponse(category []*domain.Category) []*CategoryResponse {
	mapData := make([]*CategoryResponse, 0)
	for _, el := range category {
		responItem := &CategoryResponse{
			Id:          el.ID,
			Name:        el.Name,
			Description: el.Description,
		}
		mapData = append(mapData, responItem)
	}
	return mapData
}
