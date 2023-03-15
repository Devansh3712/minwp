package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Devansh3712/minwp/handlers"
	"github.com/Devansh3712/minwp/middlewares"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/student/", http.HandlerFunc(handlers.Student))
	mux.Handle("/attendance/", middlewares.Authorization(http.HandlerFunc(handlers.Attendance)))
	mux.Handle("/examination/", middlewares.Authorization(http.HandlerFunc(handlers.Examinations)))
	mux.Handle("/result/", middlewares.Authorization(http.HandlerFunc(handlers.Result)))
	port := os.Getenv("PORT")
	handler := cors.New(cors.Options{
		AllowedMethods:       []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:       []string{"*"},
		AllowCredentials:     true,
		OptionsPassthrough:   true,
		OptionsSuccessStatus: http.StatusOK,
	})
	if err := http.ListenAndServe("0.0.0.0:"+port, handler.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
