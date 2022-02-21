package web

import (
	"encoding/json"
	"net/http"
)

type CategoryCreateRequest struct {
	Name string `validate:"required,min=5,max=100" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required,min=5,max=100" json:"name"`
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(result)
}
