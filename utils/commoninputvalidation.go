package utils

import (
	"fmt"
	"strconv"
)

// IsValidYear validates a year parameter, ensuring it is not empty and falls within a specified range.
// It takes a string representing the year and returns a boolean indicating validity and an AppErr if an error occurs.
//
// If the year parameter is empty, the function returns false and a BadRequestError indicating that the year is required.
// If the year parameter is not a valid integer, the function returns false and a BadRequestError indicating an invalid format.
// If the parsed year is within the valid range, the function returns true and nil.
// Otherwise, it returns false and a BadRequestError specifying the valid range.
func IsValidYear(yearParam string) (bool, *AppErr) {
	Logger.Debug("Validating year parameter.")
	if yearParam == "" {
		Logger.Error("Year parameter is required and cannot be empty")
		return false, NewBadRequestError("Year parameter is required and cannot be empty")
	}
	Logger.Debugf("Received non-empty year parameter: %s", yearParam)
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		Logger.Errorf("Invalid year format: %s", err.Error())
		return false, NewBadRequestError("Invalid year format, not a valid integer")
	}
	Logger.Debugf("Successfully converted year parameter to integer: %d", year)
	if year >= minValidYear && year <= maxValidYear {
		Logger.Debugf("Year %d is valid.", year)
		return true, nil
	}
	Logger.Errorf("The year is not valid. Please provide a year between %d and %d", minValidYear, maxValidYear)
	return false, NewBadRequestError(fmt.Sprintf("The year is not valid,Please provide a year between % d and %d", minValidYear, maxValidYear))
}
