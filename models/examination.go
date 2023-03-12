package models

import "encoding/json"

type ExaminationSemesters struct {
	Response struct {
		SemesterCodeInfo struct {
			SemesterCode []struct {
				Code string `json:"registrationcode"`
				Id   string `json:"registrationid"`
			} `json:"semestercode"`
		} `json:"semesterCodeinfo"`
	} `json:"response"`
}

func (e ExaminationSemesters) MarshalJSON() ([]byte, error) {
	semesters := make(map[string]string)
	response := e.Response.SemesterCodeInfo.SemesterCode
	for index := range response {
		semesters[response[index].Code] = response[index].Id
	}
	return json.Marshal(semesters)
}

type Examination struct {
	Response struct {
		EventCode struct {
			ExamEvent []struct {
				Event string `json:"exameventcode"`
				Id    string `json:"exameventid"`
			} `json:"examevent"`
		} `json:"eventcode"`
	} `json:"response"`
}

func (e Examination) MarshalJSON() ([]byte, error) {
	examination := make(map[string]string)
	response := e.Response.EventCode.ExamEvent
	for index := range response {
		examination[response[index].Event] = response[index].Id
	}
	return json.Marshal(examination)
}

type ExaminationSchedule struct {
	Response struct {
		SubjectInfo []struct {
			Timing string `json:"datetimeupto"`
			Date   string `json:"datetime"`
			Room   string `json:"roomcode"`
			Seat   string `json:"seatno"`
		} `json:"subjectinfo"`
	} `json:"response"`
}
