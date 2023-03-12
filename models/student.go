package models

import (
	"encoding/json"
	"reflect"
)

type Login struct {
	EnrollmentNumber string `json:"EnrollmentNumber"`
	Password         string `json:"Password"`
}

type Token struct {
	Response struct {
		RegData struct {
			MemberId string `json:"memberid"`
			Token    string `json:"token"`
		} `json:"regdata"`
	} `json:"response"`
}

func (tk Token) MarshalJSON() ([]byte, error) {
	response := make(map[string]any)
	tokenData := tk.Response.RegData
	v := reflect.ValueOf(tokenData)
	t := v.Type()
	for index := 0; index < v.NumField(); index++ {
		response[t.Field(index).Name] = v.Field(index).Interface()
	}
	return json.Marshal(response)
}

type Student struct {
	Response struct {
		GeneralInfo struct {
			Name             string `json:"studentname"`
			Batch            string `json:"batch"`
			EnrollmentNumber string `json:"registrationno"`
			PersonalEmail    string `json:"studentpersonalemailid"`
			InstituteEmail   string `json:"studentemailid"`
			DateOfBirth      string `json:"dateofbirth"`
			Branch           string `json:"sectioncode"`
			PhoneNumber      string `json:"studentcellno"`
			ProgramCode      string `json:"programcode"`
			Address1         string `json:"paddress1"`
			Address2         string `json:"paddress2"`
			Address3         string `json:"paddress3"`
			City             string `json:"pcityname"`
			State            string `json:"pstatename"`
		} `json:"generalinformation"`
	} `json:"response"`
}

func (s Student) MarshalJSON() ([]byte, error) {
	response := make(map[string]any)
	studentInfo := s.Response.GeneralInfo
	v := reflect.ValueOf(studentInfo)
	t := v.Type()
	for index := 0; index < v.NumField(); index++ {
		response[t.Field(index).Name] = v.Field(index).Interface()
	}
	return json.Marshal(response)
}
