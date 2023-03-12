package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Devansh3712/minwp/models"
)

func ResultSemesterList(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	var semesters models.ResultSemesters
	payload := map[string]string{"instituteid": instituteId, "studentid": userId}
	if statusCode := postRequest(
		token, baseURL+"/studentcommonsontroller/getsemestercode-exammarks", payload, &semesters,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(semesters)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GradeCardSemesterList(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	var semesters models.GradeCardSemesters
	payload := map[string]string{"instituteid": instituteId, "studentid": userId}
	if statusCode := postRequest(
		token, baseURL+"/studentgradecard/getregistrationList", payload, &semesters,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(semesters)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ExamMarks(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	var marks models.ExamMarks
	payload := map[string]any{
		"instituteid":    instituteId,
		"registrationid": semesterId,
		"memberid":       userId,
		"companyid":      nil,
	}
	if statusCode := postRequest(
		token, baseURL+"/studentsexamview/getstudent-exammarks", payload, &marks,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(marks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ExamReport(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterCode := r.URL.Query().Get("semesterCode")
	if semesterCode == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	url := baseURL + fmt.Sprintf(
		"/studentsexamview/printstudent-exammarks/%s/%s/%s/%s",
		userId, instituteId, semesterId, semesterCode,
	)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(request)
	if response.StatusCode != 200 || err != nil {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Add("Content-Disposition", "inline; filename=\"studentMarks.pdf\"")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func GradeCard(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	// Fetch Branch ID and Program ID
	var info models.StudentInfo
	payload := map[string]string{"instituteid": instituteId, "studentid": userId}
	if statusCode := postRequest(
		token, baseURL+"/studentgradecard/getstudentinfo", payload, &info,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	branchId := info.Response.Info.BranchId
	programId := info.Response.Info.ProgramId
	// Fetch GradeCard
	var grades models.GradeCard
	payload = map[string]string{
		"branchid":       branchId,
		"instituteid":    instituteId,
		"programid":      programId,
		"registrationid": semesterId,
		"studentid":      userId,
	}
	if statusCode := postRequest(
		token, baseURL+"/studentgradecard/showstudentgradecard", payload, &grades,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(grades)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GradeCardReport(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	semesterCode := r.URL.Query().Get("semesterCode")
	if semesterCode == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	semesterId := r.URL.Query().Get("semesterId")
	if semesterId == "" {
		http.Error(w, ErrEmptySemesterCode.Error(), http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, ErrEmptyName.Error(), http.StatusBadRequest)
		return
	}
	// Fetch Branch ID and Program ID
	var info models.StudentInfo
	payload := map[string]string{"instituteid": instituteId, "studentid": userId}
	if statusCode := postRequest(
		token, baseURL+"/studentgradecard/getstudentinfo", payload, &info,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	branchId := info.Response.Info.BranchId
	branchCode := info.Response.Info.BranchCode
	programId := info.Response.Info.ProgramId
	programCode := info.Response.Info.ProgramCode
	// Fetch grade card report as PDF
	url := baseURL + fmt.Sprintf(
		"/studentgradecard/printgradecard/%s/%s/%s/%s/%s/enrollmentno/%s/%s/%s/%s/%s",
		userId, instituteId, programId, branchId, semesterId,
		name, branchCode, programCode, programCode, semesterCode,
	)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(request)
	if response.StatusCode != 200 || err != nil {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Add("Content-Disposition", "inline; filename=\"gradeCard.pdf\"")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func SemesterResults(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	var results models.Results
	payload := map[string]string{
		"instituteid": instituteId,
		"studentid":   userId,
		"stynumber":   "4",
	}
	if statusCode := postRequest(
		token, baseURL+"/studentgradecard/showstudentgradecard", payload, &results,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
