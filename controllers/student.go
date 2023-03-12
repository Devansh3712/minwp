package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Devansh3712/minwp/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var student models.Login
	if err := decodeJSON(r, &student); err != "" {
		http.Error(w, err, http.StatusBadRequest)
		return
	}
	var token models.Token
	payload := map[string]string{
		"otppwd":           "PWD",
		"username":         student.EnrollmentNumber,
		"passwordotpvalue": student.Password,
		"Modulename":       "STUDENTMODULE",
	}
	if statusCode := postRequest("", baseURL+"/token/generate-token1", payload, &token); statusCode != 200 {
		http.Error(w, ErrInvalidEnrollOrPassword.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func StudentInformation(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId := r.Header.Get("User-Id")
	if token == "" || userId == "" {
		http.Error(w, ErrNotAuthorized.Error(), http.StatusUnauthorized)
		return
	}
	var student models.Student
	payload := map[string]string{
		"clientid":    "SOAU",
		"memberid":    userId,
		"instituteid": instituteId,
	}
	if statusCode := postRequest(
		token, baseURL+"/studentpersinfo/getstudent-personalinformation", payload, &student,
	); statusCode != 200 {
		http.Error(w, ErrBadResponse.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
