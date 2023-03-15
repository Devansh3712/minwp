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

func enableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	if r.Method == http.MethodOptions {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

func get(f http.HandlerFunc) http.HandlerFunc  { return allowMethod(f, http.MethodGet) }
func post(f http.HandlerFunc) http.HandlerFunc { return allowMethod(f, http.MethodPost) }
