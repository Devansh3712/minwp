package handlers

import (
	"fmt"
	"net/http"
	"strings"
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

func enableCors(w *http.ResponseWriter) {
	headers := []string{
		"Accept",
		"Content-Type",
		"Authorization",
		"User-Id",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Origin",
		// Pre-flight request headers
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
		"Origin",
	}
	(*w).Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
}

func get(f http.HandlerFunc) http.HandlerFunc  { return allowMethod(f, http.MethodGet) }
func post(f http.HandlerFunc) http.HandlerFunc { return allowMethod(f, http.MethodPost) }
