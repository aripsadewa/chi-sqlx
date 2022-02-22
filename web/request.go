package web

import (
	"encoding/json"
	"net/http"
	"time"
)

type CategoryCreateRequest struct {
	Name      string     `validate:"required,min=5,max=100" json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CategoryUpdateRequest struct {
	Id        int        `json:"id"`
	Name      string     `validate:"required,min=5,max=100" json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ParamRequest struct {
	Page  int `validate:"number" json:"page"`
	Limit int `validate:"number" json:"limit"`
	// Start time.Time `validate:"datetime"`
	// End   time.Time `validate:"datetime"`
	Sort      string `json:"sort"`
	SortValue string `json:"sort_value"`
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(result)
}
