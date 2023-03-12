package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Devansh3712/minwp/models"
)

func ExaminationSemesterList(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	var semesters models.ExaminationSemesters
	payload := map[string]string{
		"clientid":    "JAYPEE",
		"instituteid": instituteId,
		"memberid":    userId,
	}
	if statusCode := postRequest(
		token, baseURL+"/studentcommonsontroller/getsemestercode-withstudentexamevents", payload, &semesters,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(semesters)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ExaminationList(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	var examination models.Examination
	payload := map[string]string{"instituteid": instituteId, "registrationid": semesterId}
	if statusCode := postRequest(
		token, baseURL+"/studentcommonsontroller/getstudentexamevents", payload, &examination,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(examination)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ExaminationSchedule(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	examinationId := r.URL.Query().Get("examId")
	if examinationId == "" {
		http.Error(w, ErrEmptyExaminationCode.Error(), http.StatusBadRequest)
		return
	}
	var schedule models.ExaminationSchedule
	payload := map[string]string{
		"memberid":       userId,
		"instituteid":    instituteId,
		"exameventid":    examinationId,
		"registrationid": semesterId,
	}
	if statusCode := postRequest(
		token, baseURL+"/studentsttattview/getstudent-examschedule", payload, &schedule,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(schedule)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
