package services

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Errorf("Assertion failed: %s\nExpected: %v\nActual: %v", message, expected, actual)
	}
}
func assertNotEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected == actual {
		t.Errorf("Assertion failed: %s\nExpected: values to be different\nActual: %v", message, actual)
	}
}

func assertType(t *testing.T, expectedType interface{}, actual interface{}, message string) {
	expected := reflect.TypeOf(expectedType)
	actualType := reflect.TypeOf(actual)

	if expected != actualType {
		t.Errorf("Assertion failed: %s\nExpected type: %v\nActual type: %v", message, expected, actualType)
	}
}

func TestFetchHolidaysByYear(t *testing.T) {
	// Create a ResponseRecorder to capture the response
	respData, _ := FetchHolidaysByYear("2023")

	// Asserting whether the results contains all the region or not
	assertEqual(t, "england-and-wales", respData.EnglandAndWales.Division, "England and Wales")
	assertEqual(t, "scotland", respData.Scotland.Division, "Scotland")
	assertEqual(t, "northern-ireland", respData.NorthernIreland.Division, "Northern Ireland")

	// Assering whether all the region has the only the given year holidays
	assertNotEqual(t, "2022", respData.EnglandAndWales.Events[0].Date[:4], "Asserting year")
	assertNotEqual(t, "2022", respData.Scotland.Events[0].Date[:4], "Asserting year")
	assertNotEqual(t, "2022", respData.NorthernIreland.Events[0].Date[:4], "Asserting year")

	assertEqual(t, "2023", respData.EnglandAndWales.Events[0].Date[:4], "Asserting year")
	assertEqual(t, "2023", respData.Scotland.Events[0].Date[:4], "Asserting year")
	assertEqual(t, "2023", respData.NorthernIreland.Events[0].Date[:4], "Asserting year")

}

func TestFetchHolidaysForEnglandAndWales(t *testing.T) {
	//setup
	respData, _ := FetchHolidaysForEnglandAndWales()

	//assertion
	assertEqual(t, "england-and-wales", respData.Division, "Validating division name")
	assertNotEqual(t, "scotland", respData.Division, "Validating division name")
	assertNotEqual(t, "northern-ireland", respData.Division, "Validating division name")

	// randomly validating anyone of the events
	assertEqual(t, false, respData.Events[0].Bunting, "Validating bunt value")
	assertNotEqual(t, true, respData.Events[1].Bunting, "Validating bunt value")

}

func TestFetchHolidaysByYearWithTitleDate(t *testing.T) {
	// Create a ResponseRecorder to capture the response
	respData, _ := FetchHolidaysByYearWithTitleDate("2023")

	// Asserting whether the results contains all the region or not
	assertEqual(t, "england-and-wales", respData.EnglandAndWales.Division, "England and Wales")
	assertEqual(t, "scotland", respData.Scotland.Division, "Scotland")
	assertEqual(t, "northern-ireland", respData.NorthernIreland.Division, "Northern Ireland")

	// Assering whether all the region has the only the given year holidays
	assertNotEqual(t, "2022", respData.EnglandAndWales.Events[0].Date[:4], "Asserting year")
	assertNotEqual(t, "2022", respData.Scotland.Events[0].Date[:4], "Asserting year")
	assertNotEqual(t, "2022", respData.NorthernIreland.Events[0].Date[:4], "Asserting year")

	assertEqual(t, "2023", respData.EnglandAndWales.Events[0].Date[:4], "Asserting year")
	assertEqual(t, "2023", respData.Scotland.Events[0].Date[:4], "Asserting year")
	assertEqual(t, "2023", respData.NorthernIreland.Events[0].Date[:4], "Asserting year")

	assertType(t, ModifiedEvent{}, respData.Scotland.Events[0], "Validating the type of output")

	//randomly validating anyone of the events for the attributes

}
