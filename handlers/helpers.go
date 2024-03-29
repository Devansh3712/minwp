package handlers

import (
	"fmt"
	"net/http"
)

func allowMethod(f http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			http.Error(
				w,
				fmt.Sprintf("endpoint does not allow %s method", r.Method),
				http.StatusMethodNotAllowed,
			)
			return
		}
		f(w, r)
	}
}

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "endpoint not found", http.StatusNotFound)
	}
}

func get(f http.HandlerFunc) http.HandlerFunc  { return allowMethod(f, http.MethodGet) }
func post(f http.HandlerFunc) http.HandlerFunc { return allowMethod(f, http.MethodPost) }
