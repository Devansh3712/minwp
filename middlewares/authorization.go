package middlewares

import (
	"errors"
	"net/http"
)

var ErrNotAuthorized = errors.New("student not logged in")

func Authorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		userId := r.Header.Get("User-Id")
		if token == "" || userId == "" {
			http.Error(w, ErrNotAuthorized.Error(), http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
