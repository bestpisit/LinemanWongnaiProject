package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type covidCase struct {
	Age      *int    `json:"Age,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type covidData struct {
	Data []covidCase `json:"Data"`
}

type CovidSummaryTemplate struct {
	Province map[string]int
	AgeGroup map[string]int
}

func getcovidData(url string) ([]covidCase, error) {
	response, apiError := http.Get(url)

	if apiError != nil {
		return nil, apiError
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data covidData
	if er := json.Unmarshal(body, &data); er != nil {
		return nil, er
	}

	return data.Data, nil
}

func composeProvince(covidCases []covidCase) map[string]int {
	provinceData := make(map[string]int)
	for _, covidCase := range covidCases {
		if covidCase.Province == nil || *covidCase.Province == "" || *covidCase.Province == "non" || *covidCase.Province == "none" {
			provinceData["N/A"]++
			continue
		}
		provinceData[*covidCase.Province]++
	}
	return provinceData
}

func composeAgeGroup(covidCases []covidCase) map[string]int {
	ageGroupData := make(map[string]int)
	for _, covidCase := range covidCases {
		if covidCase.Age == nil {
			ageGroupData["N/A"]++
		} else if *covidCase.Age >= 61 {
			ageGroupData["61+"]++
		} else if *covidCase.Age > 30 {
			ageGroupData["31-60"]++
		} else if *covidCase.Age >= 0 {
			ageGroupData["0-30"]++
		} else {
			ageGroupData["N/A"]++
		}
	}
	return ageGroupData
}

func Covid19SummaryFromURL(url string) (CovidSummaryTemplate, error) {
	var covidSummary CovidSummaryTemplate
	covidData, err := getcovidData(url)
	if err != nil {
		return covidSummary, err
	}
	covidSummary.Province = composeProvince(covidData)
	covidSummary.AgeGroup = composeAgeGroup(covidData)
	return covidSummary, nil
}

func Covid19SummaryFromData(datastring string) (CovidSummaryTemplate, error) {
	var covidSummary CovidSummaryTemplate
	var data covidData
	if er := json.Unmarshal([]byte(datastring), &data); er != nil {
		return covidSummary, er
	}
	covidSummary.Province = composeProvince(data.Data)
	covidSummary.AgeGroup = composeAgeGroup(data.Data)
	return covidSummary, nil
}