package services

import "ukbankholidays/utils"

// FetchHolidaysByYear fetches holiday data for the specified year and filters events for England, Scotland, and Northern Ireland.
// It returns the filtered holiday data or an error if there's an issue fetching or filtering the data.
func FetchHolidaysByYear(year string) (*Data, *utils.AppErr) {
	holidayResp, err := FetchHolidaysDetails()
	if err != nil {
		return nil, err
	}
	holidayResp.EnglandAndWales.Events = FilterHolidaysByYear(holidayResp.EnglandAndWales.Events, year)
	utils.Logger.Debug("Successfully filtered events for England and Wales by year.")

	holidayResp.Scotland.Events = FilterHolidaysByYear(holidayResp.Scotland.Events, year)
	utils.Logger.Debug("Successfully filtered events for Scotland by year.")

	holidayResp.NorthernIreland.Events = FilterHolidaysByYear(holidayResp.NorthernIreland.Events, year)
	utils.Logger.Debug("Successfully filtered events for Northern Ireland by year.")

	return holidayResp, nil
}

// FetchHolidaysByYearWithTitleDate fetches holiday data for the specified year and modifies the output by extracting titles and dates.
// It returns the modified holiday data or an error if there's an issue fetching or modifying the data.
func FetchHolidaysByYearWithTitleDate(year string) (*ModifiedData, *utils.AppErr) {
	utils.Logger.Debug("Fetching holidays by year with titles and dates.")
	respData, respErr := FetchHolidaysByYear(year)
	var ModifiedOutput ModifiedData
	if respErr != nil {
		return nil, respErr
	}
	ModifiedOutput.EnglandAndWales = ModifyOutput(respData.EnglandAndWales)
	ModifiedOutput.Scotland = ModifyOutput(respData.Scotland)
	ModifiedOutput.NorthernIreland = ModifyOutput(respData.NorthernIreland)
	return &ModifiedOutput, nil

}

// ModifyOutput takes a Division and extracts titles and dates, returning a ModifiedDivision.
// It creates a modified representation of the input division's events with only titles and dates.
func ModifyOutput(ip Division) ModifiedDivision {
	utils.Logger.Debug("Modifying output structure.")
	var output ModifiedDivision
	var outputEvents []ModifiedEvent
	output.Division = ip.Division
	for _, value := range ip.Events {
		var holiday ModifiedEvent
		holiday.Title = value.Title
		holiday.Date = value.Date
		outputEvents = append(outputEvents, holiday)
	}
	output.Events = outputEvents
	utils.Logger.Debug("Successfully modified output structure with titles and dates.")
	return output
}

// FilterHolidaysByYear filters a slice of events based on the specified year.
// It returns a new slice containing only events with dates matching the provided year.
func FilterHolidaysByYear(events []Event, year string) []Event {
	utils.Logger.Debug("Filtering events by year.")
	var filteredEvents []Event
	for _, value := range events {
		if value.Date[:4] == year {
			filteredEvents = append(filteredEvents, value)
		}
	}
	utils.Logger.Debug("Successfully filtered events by year.")
	return filteredEvents
}

// FetchHolidaysForEnglandAndWales fetches holiday data for England and Wales, filtering out events marked as non-bunting.
// It returns a Division containing the filtered holiday events or an error if there's an issue fetching or filtering the data.
func FetchHolidaysForEnglandAndWales() (*Division, *utils.AppErr) {
	utils.Logger.Debug("Fetching holidays for England and Wales.")
	var filteredDivision Division
	holidayResp, err := FetchHolidaysDetails()
	if err != nil {
		return nil, err
	}

	filteredDivision.Events = filterEventsByBuntingFalse(holidayResp.EnglandAndWales.Events)
	utils.Logger.Debug("Successfully filtered events by bunting status for England and Wales.")

	filteredDivision.Division = "england-and-wales"
	utils.Logger.Debug("Successfully set division identifier for England and Wales.")

	return &filteredDivision, nil
}

// filterEventsByBuntingFalse filters a slice of events, excluding those with the Bunting field set to true.
// It returns a new slice containing only events with Bunting set to false.
func filterEventsByBuntingFalse(events []Event) []Event {
	utils.Logger.Debug("Filtering events by bunting status.")
	var filteredEvents []Event
	for _, value := range events {
		if !value.Bunting {
			filteredEvents = append(filteredEvents, value)
		}
	}
	utils.Logger.Debug("successfully filtered events by bunting status.")
	return filteredEvents
}
