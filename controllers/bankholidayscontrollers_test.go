package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Errorf("Assertion failed: %s\nExpected: %v\nActual: %v", message, expected, actual)
	}
}

func TestFetchHolidaysByYearForSuccessfulResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year?year=2023", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusOK, response.Code, "Response code")
}

func TestFetchHolidaysByYearForInvalidInputEmptyYear(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year?year=", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusBadRequest, response.Code, "Response code")
}

func TestFetchHolidaysByYearForInvalidInputWrongYearFormat(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year?year=2022.3", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusBadRequest, response.Code, "Response code")
}

func TestFetchHolidaysByYearForInvalidInputWrongYearRange(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year?year=1800", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusBadRequest, response.Code, "Response code")
}
func TestFetchTitleDateOfHolidaysByYearForSuccessfulResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year-title-date?year=2023", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchTitleDateOfHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusOK, response.Code, "Response code")
}

func TestFetchTitleDateOfHolidaysByYearForInvalidInputEmptyYear(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year-title-date?year=", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchTitleDateOfHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusBadRequest, response.Code, "Response code")
}

func TestFetchTitleDateOfHolidaysByYearForInvalidInputWrongYearFormat(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year-title-date?year=2022.3", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchTitleDateOfHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusBadRequest, response.Code, "Response code")
}

func TestFetchTitleDateOfHolidaysByYearForInvalidInputWrongYearRange(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-by-year-title-date?year=1800", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchTitleDateOfHolidaysByYear(response, request, nil)
	assertEqual(t, http.StatusBadRequest, response.Code, "Response code")
}

func TestFetchHolidaysOfEnglandWalesForSuccessfulResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "/holidays-england-wales", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()
	FetchHolidaysOfEnglandWales(response, request, nil)
	assertEqual(t, http.StatusOK, response.Code, "Response code")
}
