package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Devansh3712/minwp/handlers"
	"github.com/Devansh3712/minwp/middlewares"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/student/", http.HandlerFunc(handlers.Student))
	mux.Handle("/attendance/", middlewares.Authorization(http.HandlerFunc(handlers.Attendance)))
	mux.Handle("/examination/", middlewares.Authorization(http.HandlerFunc(handlers.Examinations)))
	mux.Handle("/result/", middlewares.Authorization(http.HandlerFunc(handlers.Result)))
	port := os.Getenv("PORT")
	if err := http.ListenAndServe("0.0.0.0:"+port, mux); err != nil {
		log.Fatal(err)
	}
}
