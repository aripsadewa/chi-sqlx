package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

type Standard struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta,omitempty"`
}

type Meta struct {
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

func Json(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		Error(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	if string(data) == "null" {
		_, _ = w.Write([]byte("[]"))
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		Error(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}
}
