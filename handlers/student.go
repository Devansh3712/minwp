package handlers

import (
	"net/http"
	"regexp"

	"github.com/Devansh3712/minwp/controllers"
)

var (
	login       = regexp.MustCompile(`/student/login[/]?`)
	studentInfo = regexp.MustCompile(`/student/info[/]?`)
)

func Student(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	path := r.URL.Path
	switch {
	case login.MatchString(path):
		handler = post(controllers.Login)
	case studentInfo.MatchString(path):
		handler = get(controllers.StudentInformation)
	default:
		handler = notFound()
	}
	handler.ServeHTTP(w, r)
}
