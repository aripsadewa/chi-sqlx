package web

import (
	"encoding/json"
	"net/http"
	"rest_api/model/domain"
)

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(category domain.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryMeta(category domain.CategoryMeta) *MetaData {
	return &MetaData{
		Limit:     int(category.Limit),
		Total:     category.Total,
		Page:      category.Page,
		TotalPage: category.TotalPage,
	}
}

func AllCategoryResponse(category []*domain.Category) []*CategoryResponse {
	mapData := make([]*CategoryResponse, 0)
	for _, el := range category {
		responItem := &CategoryResponse{
			Id:   el.ID,
			Name: el.Name,
		}
		mapData = append(mapData, responItem)
	}
	return mapData
}

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type GetAllCategory struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Data     interface{} `json:"data"`
	MetaData interface{} `json:"meta"`
}

type CategoryMetaResponse struct {
	Category  []CategoryResponse `json:"category"`
	Total     int                `db:"total" json:"total"`
	Page      int                `db:"page" json:"page"`
	TotalPage int                `db:"total_page" json:"total_page"`
}

type MetaData struct {
	Limit     int     `json:"limit"`
	Total     float64 `db:"total" json:"total"`
	Page      int     `db:"page" json:"page"`
	TotalPage float64 `db:"total_page" json:"total_page"`
}

type ErrorResponse struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Errors []WebError `json:"errors"`
}

type WebError struct {
	Message string `json:"message"`
}

type ParamResponse struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	// Start time.Time `json:"start_date"`
	// End   time.Time `json:"end_date"`
	// Sort  string    `json:"sort"`
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
