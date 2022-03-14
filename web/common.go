package web

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	Status   string      `json:"status"`
	Errors   interface{} `json:"errors"`
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata"`
}

type PaginateMetaData struct {
	Offset    int     `json:"offset"`
	Limit     float64 `json:"limit"`
	Total     int     `json:"total"`
	Page      float64 `json:"page"`
	TotalPage int     `json:"total_page"`
}

type WebError struct {
	Message string `json:"message"`
}

func WriteToResponseBody(writer http.ResponseWriter, code int, status string, data interface{}, errors interface{}, metadata interface{}) {
	writer.WriteHeader(code)
	writer.Header().Add("Content-Type", "application/json")
	res := WebResponse{
		Status:   status,
		Data:     data,
		Errors:   errors,
		Metadata: metadata,
	}
	resByte, _ := json.Marshal(&res)

	writer.Write(resByte)
}
