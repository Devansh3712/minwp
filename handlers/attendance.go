package handlers

import (
	"net/http"
	"regexp"

	"github.com/Devansh3712/minwp/controllers"
)

var (
	attendanceSemesters = regexp.MustCompile(`/attendance/semesters[/]?`)
	faculty             = regexp.MustCompile(`/attendance/faculty[/]?`)
	attendance          = regexp.MustCompile(`/attendance[/]?`)
	attendanceDetails   = regexp.MustCompile(`/attendance/details[/]?`)
)

func Attendance(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	var handler http.Handler
	path := r.URL.Path
	switch {
	case attendanceSemesters.MatchString(path):
		handler = get(controllers.AttendanceSemesterList)
	case faculty.MatchString(path):
		handler = get(controllers.Faculty)
	case attendance.MatchString(path):
		handler = get(controllers.Attendance)
	case attendanceDetails.MatchString(path):
		handler = get(controllers.AttendanceDetails)
	default:
		handler = notFound()
	}
	handler.ServeHTTP(w, r)
}
