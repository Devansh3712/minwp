package models

import (
	"encoding/json"
	"reflect"
	"strings"
)

type AttendenceSemesters struct {
	Response struct {
		SemesterList []struct {
			Code string `json:"registrationcode"`
			Id   string `json:"registrationid"`
		} `json:"semlist"`
	} `json:"response"`
}

func (a AttendenceSemesters) MarshalJSON() ([]byte, error) {
	semesters := make(map[string]string)
	response := a.Response.SemesterList
	for index := range response {
		semesters[response[index].Code] = response[index].Id
	}
	return json.Marshal(semesters)
}

type Faculty struct {
	Response struct {
		Registrations []struct {
			Name         string `json:"employeename"`
			Subject      string `json:"subjectdesc"`
			Type         string `json:"subjectcomponentcode"`
			Credits      int    `json:"credits"`
			AuditSubject string `json:"audtsubject"`
			MinorSubject string `json:"minorsubject"`
			EmployeeCode string `json:"employeecode"`
			SubjectCode  string `json:"subjectcode"`
			SubjectId    string `json:"subjectid"`
		} `json:"registrations"`
	} `json:"response"`
}

func (f Faculty) MarshalJSON() ([]byte, error) {
	response := make(map[string][]map[string]any)
	facultyInfo := f.Response.Registrations
	for _, data := range facultyInfo {
		v := reflect.ValueOf(data)
		t := v.Type()
		info := make(map[string]any)
		for index := 0; index < v.NumField(); index++ {
			info[t.Field(index).Name] = v.Field(index).Interface()
		}
		if _, ok := response[data.Subject]; !ok {
			response[data.Subject] = make([]map[string]any, 0)
		}
		response[data.Subject] = append(response[data.Subject], info)
	}
	return json.Marshal(response)
}

type Attendance struct {
	Response struct {
		AttendanceList []struct {
			Subject                   string `json:"subjectcode"`
			SubjectId                 string `json:"subjectid"`
			Absent                    any    `json:"abseent"`
			LectureId                 string `json:"Lsubjectcomponentid"`
			LectureTutorialPercentage any    `json:"LTpercantage"`
			LecturePercentage         any    `json:"Lpercentage"`
			LectureTotalClasses       any    `json:"Ltotalclass"`
			LectureTotalPresent       any    `json:"Ltotalpres"`
			TutorialId                string `json:"Tsubjectcomponentid"`
			TutorialPercentage        any    `json:"Tpercentage"`
			TutorialTotalClasses      any    `json:"Ttotalclass"`
			TutorialTotalPresent      any    `json:"Ttotalpres"`
			PracticalId               string `json:"Psubjectcomponentid"`
			PracticalPercentage       any    `json:"Ppercentage"`
		} `json:"studentattendancelist"`
	} `json:"response"`
}

func (a Attendance) MarshalJSON() ([]byte, error) {
	response := make(map[string]map[string]any)
	facultyInfo := a.Response.AttendanceList
	for _, data := range facultyInfo {
		v := reflect.ValueOf(data)
		t := v.Type()
		info := make(map[string]any)
		for index := 0; index < v.NumField(); index++ {
			info[t.Field(index).Name] = v.Field(index).Interface()
		}
		key := strings.Split(data.Subject, "(")
		response[key[0]] = info
	}
	return json.Marshal(response)
}

type AttendanceDetail struct {
	Response struct {
		AttendanceList []struct {
			AttendanceBy string `json:"attendanceby"`
			ClassType    string `json:"classtype"`
			DateTime     string `json:"datetime"`
			Present      string `json:"present"`
		} `json:"studentAttdsummarylist"`
	} `json:"response"`
}

func (a AttendanceDetail) MarshalJSON() ([]byte, error) {
	response := make([]map[string]any, 0)
	facultyInfo := a.Response.AttendanceList
	for _, data := range facultyInfo {
		v := reflect.ValueOf(data)
		t := v.Type()
		info := make(map[string]any)
		for index := 0; index < v.NumField(); index++ {
			info[t.Field(index).Name] = v.Field(index).Interface()
		}
		response = append(response, info)
	}
	return json.Marshal(response)
}
