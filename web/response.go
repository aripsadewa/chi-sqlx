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
	categoryResponse := &CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
	if category.Description.Valid {
		categoryResponse.Description = category.Description.String
	}
	return categoryResponse
}

func ToCategoriesResponse(category []*domain.Category) []*CategoryResponse {
	mapData := make([]*CategoryResponse, 0)
	for _, el := range category {
		responItem := &CategoryResponse{
			Id:   el.ID,
			Name: el.Name,
		}
		if el.Description.Valid {
			responItem.Description = el.Description.String
		}
		mapData = append(mapData, responItem)
	}
	return mapData
}
