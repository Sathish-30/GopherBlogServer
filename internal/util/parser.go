package util

import (
	"encoding/json"
	"net/http"

	"github.com/sathish-30/GopherBlogServer/internal/model"
)

func WriteJson(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func ConvertJson[T model.User](r *http.Request) (T, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	return data, err
}
