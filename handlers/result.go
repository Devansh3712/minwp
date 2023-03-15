package handlers

import (
	"net/http"
	"regexp"

	"github.com/Devansh3712/minwp/controllers"
)

var (
	resultSemesters = regexp.MustCompile(`/result/semesters[/]?`)
	examMarks       = regexp.MustCompile(`/result/marks[/]?`)
	examReport      = regexp.MustCompile(`/result/report[/]?`)
	gradeCard       = regexp.MustCompile(`/result/grade-card[/]?`)
	gradeCardReport = regexp.MustCompile(`/result/grade-card/report[/]?`)
	results         = regexp.MustCompile(`/result/all[/]?`)
)

func Result(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var handler http.Handler
	path := r.URL.Path
	switch {
	case resultSemesters.MatchString(path):
		handler = get(controllers.ResultSemesterList)
	case examMarks.MatchString(path):
		handler = get(controllers.ExamMarks)
	case examReport.MatchString(path):
		handler = get(controllers.ExamReport)
	case gradeCard.MatchString(path):
		handler = get(controllers.GradeCard)
	case gradeCardReport.MatchString(path):
		handler = get(controllers.GradeCardReport)
	case results.MatchString(path):
		handler = get(controllers.SemesterResults)
	default:
		handler = notFound()
	}
	handler.ServeHTTP(w, r)
}
