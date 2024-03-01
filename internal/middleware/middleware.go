package middleware

import (
	"errors"
	"net/http"

	"github.com/sathish-30/GopherBlogServer/internal/util"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandlerFunc(next apiFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			util.WriteJson(w, http.StatusBadRequest, errors.New("unable to parse request body into user object"))
		}
	})
}
