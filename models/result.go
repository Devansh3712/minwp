package models

import (
	"encoding/json"
	"reflect"
)

type StudentInfo struct {
	Response struct {
		Info struct {
			BranchId    string `json:"branchid"`
			BranchCode  string `json:"branchcode"`
			ProgramId   string `json:"programid"`
			ProgramCode string `json:"programcode"`
		} `json:"studentinfo"`
	} `json:"response"`
}

type ResultSemesters struct {
	Response struct {
		SemesterCode []struct {
			Code string `json:"registrationcode"`
			Id   string `json:"registrationid"`
		} `json:"semestercode"`
	} `json:"response"`
}

func (r ResultSemesters) MarshalJSON() ([]byte, error) {
	semesters := make(map[string]string)
	response := r.Response.SemesterCode
	for index := range response {
		semesters[response[index].Code] = response[index].Id
	}
	return json.Marshal(semesters)
}

type GradeCardSemesters struct {
	Response struct {
		Registrations []struct {
			Code string `json:"registrationcode"`
			Id   string `json:"registrationid"`
		} `json:"registrations"`
	} `json:"response"`
}

func (g GradeCardSemesters) MarshalJSON() ([]byte, error) {
	semesters := make(map[string]string)
	response := g.Response.Registrations
	for index := range response {
		semesters[response[index].Code] = response[index].Id
	}
	return json.Marshal(semesters)
}

type ExamMarks struct {
	Response struct {
		Marks []struct {
			Subject       string  `json:"subjectdesc"`
			FullMarks     int     `json:"fullmarks"`
			ObtainedMarks float32 `json:"obtainedmarks"`
		} `json:"viewmarksbystudent"`
	} `json:"response"`
}

func (e ExamMarks) MarshalJSON() ([]byte, error) {
	response := make([]map[string]any, 0)
	marksInfo := e.Response.Marks
	for _, data := range marksInfo {
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

type GradeCard struct {
	Response struct {
		Grades []struct {
			Subject string `json:"subjectdesc"`
			Grade   string `json:"grade"`
		} `json:"gradecard"`
	} `json:"response"`
}

type Results struct {
	Response struct {
		SemesterList []struct {
			SGPA          float32 `json:"sgpa"`
			CGPA          float32 `json:"cgpa"`
			GradePoints   float32 `json:"totalgradepoints"`
			CourseCredits float32 `json:"totalcoursecredit"`
			EarnedCredits float32 `json:"totalearnedcredits"`
		} `json:"semesterList"`
	} `json:"response"`
}
