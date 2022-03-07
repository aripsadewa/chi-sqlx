package web

import (
	"encoding/json"
	"net/http"

	"gopkg.in/guregu/null.v4"
)

// Create category the model for an category
type CategoryCreateRequest struct {
	Name        string `validate:"required,min=5,max=100" json:"name"`
	Description string `validate:"required,min=5,max=100" json:"description"`
}

type CategoryUpdateRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetParamRequest struct {
	Page      null.Int    `validate:"number" json:"page" schema:"page"`
	Limit     null.Int    `validate:"number" json:"limit" schema:"limit"`
	Start     null.Time   `validate:"datetime" schema:"start"`
	End       null.Time   `validate:"datetime" schema:"end"`
	Sort      null.String `json:"sort" schema:"sort"`
	Name      null.String `json:"name" schema:"name"`
	SortValue null.String `json:"sort_value" schema:"sort_value"`
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(result)
}
