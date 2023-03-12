package handlers

import (
	"net/http"
	"regexp"

	"github.com/Devansh3712/minwp/controllers"
)

var (
	examinationList      = regexp.MustCompile(`/examination[/]?`)
	examinationSemesters = regexp.MustCompile(`/examination/semesters[/]?`)
	examinationSchedule  = regexp.MustCompile(`/examination/schedule[/]?`)
)

func Examinations(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var handler http.Handler
	path := r.URL.Path
	switch {
	case examinationSemesters.MatchString(path):
		handler = get(controllers.ExaminationSemesterList)
	case examinationList.MatchString(path):
		handler = get(controllers.ExaminationList)
	case examinationSchedule.MatchString(path):
		handler = get(controllers.ExaminationSchedule)
	default:
		handler = notFound()
	}
	handler.ServeHTTP(w, r)
}
