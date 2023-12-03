package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ukbankholidays/services"
	"ukbankholidays/utils"
)

// FetchHolidaysByYear handles HTTP requests to fetch holidays for a specific year.
//
// This function expects a query parameter "year" to be present in the request URL.
// It validates the year, fetches holidays for the specified year using the services.FetchHolidaysByYear
// function, and responds with the result or an error.
//
// If the year is not valid, the function returns a JSON response with a 400 Bad Request status.
// If there is an error fetching holidays, the function returns a JSON response with a 500 Internal Server Error status.
// If the request is successful, the function responds with a JSON representation of the fetched holidays
// and a 200 OK status.
func FetchHolidaysByYear(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.Logger.Infof("Received %s request for %s", r.Method, r.URL.Path)
	utils.Logger.Debug("Handling request to fetch holidays by year.")

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Extract year from query parameters
	year := r.URL.Query().Get("year")
	utils.Logger.Debugf("Received request for year: %s", year)

	// Validate the year
	valid, err := utils.IsValidYear(year)
	if !valid {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Fetch holidays by year
	serviceResp, err := services.FetchHolidaysByYear(year)
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Respond with the fetched holidays
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceResp)
	utils.Logger.Infof("request has been processed successfully")
}

// FetchHolidaysOfEnglandWales handles HTTP requests to fetch holidays for England and Wales.
//
// This function does not require any request parameters. It fetches holidays for
// England and Wales using the services.FetchHolidaysForEnglandAndWales function
// and responds with the result or an error.
//
// If there is an error fetching holidays, sets the appropriate HTTP status code in the response,
// and returns a JSON error response.
//
// If the request is successful, the function responds with a JSON representation
// of the fetched holidays and a 200 OK status.
func FetchHolidaysOfEnglandWales(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.Logger.Infof("Received %s request for %s", r.Method, r.URL.Path)
	utils.Logger.Debug("Handling request to fetch holidays for England and Wales.")

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Fetch holidays for England and Wales
	serviceResp, err := services.FetchHolidaysForEnglandAndWales()
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	// Respond with the fetched holidays
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceResp)
	utils.Logger.Infof("request has been processed successfully")
}

// FetchTitleDateOfHolidaysByYear handles HTTP requests to fetch holidays for a specific year
// along with their titles and dates.
// This function expects a query parameter "year" to be present in the request URL.
// It validates the year, fetches holidays with title and holiday only for the specified year using the services.FetchHolidaysByYearWithTitleDate
// function, and responds with the result or an error.
//
// If the year is not valid, the function returns a JSON response with a 400 Bad Request status.
// If there is an error fetching holidays, the function returns a JSON response with a 500 Internal Server Error status.
// If the request is successful, the function responds with a JSON representation of the fetched holidays
// and a 200 OK status.
func FetchTitleDateOfHolidaysByYear(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.Logger.Infof("Received %s request for %s", r.Method, r.URL.Path)
	utils.Logger.Debug("Handling request to fetch holidays with titles and dates by year.")

	w.Header().Set("Content-Type", "application/json")
	year := r.URL.Query().Get("year")
	utils.Logger.Debugf("Received request for year: %s", year)

	valid, err := utils.IsValidYear(year)
	if !valid {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	serviceResp, err := services.FetchHolidaysByYearWithTitleDate(year)
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceResp)
	utils.Logger.Infof("request has been processed successfully")
}
