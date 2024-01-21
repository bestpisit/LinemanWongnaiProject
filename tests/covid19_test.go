package tests

import (
    "testing"
    "covid-19/utils"
)

//Main test with proper Age and Province Format
func TestCovid19Main(t *testing.T) {
    covidSummary, err := utils.Covid19SummaryFromData(`{"Data":[{"Age": 1},{"Age":66,"Province":"Thai"}]}`)
    if err != nil || covidSummary.Province["Thai"] != 1 || covidSummary.AgeGroup["0-30"] != 1 || covidSummary.AgeGroup["61+"] != 1 {
        t.Fatal("Task#1 Failed")
    }
}

//Second test with improper Age and Province Format
func TestCovid19Second(t *testing.T) {
    covidSummary, err := utils.Covid19SummaryFromData(`{"Data":[{"Ages": 1},{"Age":66,"Provinces":"Thai"}]}`)
    if err != nil || covidSummary.Province["Thai"] != 0 || covidSummary.AgeGroup["0-30"] != 0 || covidSummary.AgeGroup["61+"] != 1 {
        t.Fatal("Task#2 Failed")
    }
}