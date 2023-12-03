package app

import (
	"github.com/julienschmidt/httprouter"
	"ukbankholidays/controllers"
)

// It creates a new httprouter.Router and registers several endpoints for
// handling different holiday-related queries. The configured router is then
// returned, ready to be used in the application's server setup.
//
// Endpoints:
// - "/holidays-by-year": Fetches holidays for a specific year.
// - "/holidays-england-wales": Fetches holidays specific to England and Wales.
// - "/holidays-by-year-title-date": Fetches titles and dates of holidays for a specific year.
func setupRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/holidays-by-year", controllers.FetchHolidaysByYear)
	router.GET("/holidays-england-wales", controllers.FetchHolidaysOfEnglandWales)
	router.GET("/holidays-by-year-title-date", controllers.FetchTitleDateOfHolidaysByYear)
	return router
}
