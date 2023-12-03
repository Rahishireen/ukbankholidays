package services

import (
	"encoding/json"
	"net/http"
	"ukbankholidays/utils"
)

// Event represents an event in the holiday data with details such as title, date, notes, and bunting status.
type Event struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Notes   string `json:"notes"`
	Bunting bool   `json:"bunting"`
}

// Division represents a division in the holiday data, including a division identifier and a slice of events.
type Division struct {
	Division string  `json:"division"`
	Events   []Event `json:"events"`
}

// Data represents the overall structure of the holiday JSON data, including divisions for England and Wales, Scotland, and Northern Ireland.
type Data struct {
	EnglandAndWales Division `json:"england-and-wales"`
	Scotland        Division `json:"scotland"`
	NorthernIreland Division `json:"northern-ireland"`
}

// FetchHolidaysDetails fetches holiday data from the Gov UK bank holidays API.
// It performs an HTTP GET request to the API endpoint and parses the JSON response into a Data struct.
// If the request or parsing encounters an error, the function logs the error and returns an internal server error.
// Otherwise, it returns the parsed holiday data or an internal server error if the HTTP response status is not OK.
func FetchHolidaysDetails() (*Data, *utils.AppErr) {
	resp, respErr := http.Get("https://www.gov.uk/bank-holidays.json")
	if respErr != nil {
		utils.Logger.Errorf("HTTP request failed: %s", respErr.Error())
		return nil, utils.NewInternalServerError("Internal server error while fetching holiday details")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		utils.Logger.Errorf("HTTP request failed with status: %v", resp.Status)
		return nil, utils.NewInternalServerError("Internal server error while fetching holiday details")
	}
	utils.Logger.Debug("Received a successful HTTP status code.")
	var result *Data
	decodeErr := json.NewDecoder(resp.Body).Decode(&result)
	if decodeErr != nil {
		utils.Logger.Errorf("Error decoding JSON response: %s", decodeErr.Error())
		return nil, utils.NewInternalServerError("Internal server error while fetching holiday details")
	}
	utils.Logger.Debug("Successfully decoded JSON response.")
	return result, nil
}

// ModifiedEvent represents a modified event structure in the holiday data with details such as title and date.
type ModifiedEvent struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

// ModifiedDivision represents a modified division in the holiday data, including a division identifier and a slice of modified events.
type ModifiedDivision struct {
	Division string          `json:"division"`
	Events   []ModifiedEvent `json:"events"`
}

// ModifiedData represents the overall structure of the modified holiday JSON data, including divisions for England and Wales, Scotland, and Northern Ireland.
type ModifiedData struct {
	EnglandAndWales ModifiedDivision `json:"england-and-wales"`
	Scotland        ModifiedDivision `json:"scotland"`
	NorthernIreland ModifiedDivision `json:"northern-ireland"`
}
