package web

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/guregu/null.v4"
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

type GetParamRequest struct {
	Page      null.Float  `validate:"number" json:"page" schema:"page"`
	Limit     null.Float  `validate:"number" json:"limit" schema:"limit"`
	Start     null.Time   `validate:"datetime" schema:"start"`
	End       null.Time   `validate:"datetime" schema:"end"`
	Sort      null.String `json:"sort" schema:"sort"`
	SortValue null.String `json:"sort_value" schema:"sort_value"`
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(result)
}
