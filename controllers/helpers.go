package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	baseURL     = "https://webportal.jiit.ac.in:6011/StudentPortalAPI"
	instituteId = "11IN1902J000001"

	ErrNotAuthorized           = errors.New("student not logged in")
	ErrBadResponse             = errors.New("bad response from webportal server")
	ErrInvalidEnrollOrPassword = errors.New("invalid or empty enrollment number or password")
	ErrEmptySemesterCode       = errors.New("empty semester code")
	ErrEmptySubjectCode        = errors.New("empty subject code")
	ErrEmptyExaminationCode    = errors.New("empty examination code")
	ErrEmptyName               = errors.New("empty student name")
)

func decodeJSON(r *http.Request, model interface{}) string {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&model); err != nil {
		var res string
		var typeError *json.UnmarshalTypeError
		var syntaxError *json.SyntaxError

		switch {
		case errors.As(err, &syntaxError):
			res = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			res = "Request body contains badly-formed JSON"
		case errors.As(err, &typeError):
			res = fmt.Sprintf(
				"Request body contains an invalid value for the %q field (at position %d)",
				typeError.Field, typeError.Offset,
			)
		case errors.Is(err, io.EOF):
			res = "Request body must not be empty"
		default:
			res = "Unable to parse request body"
		}
		return res
	}
	return ""
}

func postRequest(token, url string, requestModel, responseModel interface{}) int {
	requestBody, _ := json.Marshal(requestModel)
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	if token != "" {
		request.Header.Add("Authorization", "Bearer "+token)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if response.StatusCode != 200 || err != nil {
		log.Println(err)
		return -1
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	json.Unmarshal(responseBody, &responseModel)
	return response.StatusCode
}
