package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Devansh3712/minwp/models"
)

func AttendanceSemesterList(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	var semesters models.AttendenceSemesters
	payload := map[string]string{
		"clientid":    "JAYPEE",
		"instituteid": instituteId,
		"studentid":   userId,
		"membertype":  "S",
	}
	if statusCode := postRequest(
		token, baseURL+"/StudentClassAttendance/getstudentInforegistrationforattendence", payload, &semesters,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(semesters)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func Faculty(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	var faculty models.Faculty
	payload := map[string]string{
		"studentid":      userId,
		"instituteid":    instituteId,
		"registrationid": semesterId,
	}
	if statusCode := postRequest(
		token, baseURL+"/reqsubfaculty/getfaculties", payload, &faculty,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(faculty)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func Attendance(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	var attendance models.Attendance
	payload := map[string]string{
		"clientid":       "JAYPEE",
		"instituteid":    instituteId,
		"studentid":      userId,
		"stynumber":      "4",
		"registrationid": semesterId,
	}
	if statusCode := postRequest(
		token, baseURL+"/StudentClassAttendance/getstudentattendancedetail", payload, &attendance,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(attendance)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func AttendanceDetails(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	subjectId := r.URL.Query().Get("id")
	subjectComponentId := r.URL.Query().Get("componentId")
	if subjectId == "" || subjectComponentId == "" {
		http.Error(w, ErrEmptySubjectCode.Error(), http.StatusBadRequest)
		return
	}
	var details models.AttendanceDetail
	payload := map[string]any{
		"clientid":       "JAYPEE",
		"instituteid":    instituteId,
		"studentid":      userId,
		"subjectid":      subjectId,
		"registrationid": semesterId,
		"cmpidkey":       []map[string]string{{"subjectcomponentid": subjectComponentId}},
	}
	if statusCode := postRequest(
		token, baseURL+"/StudentClassAttendance/getstudentsubjectpersentage", payload, &details,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(details)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
