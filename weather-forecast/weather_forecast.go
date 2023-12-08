// Package weather return the Forecast and weather conditions for cities.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string

// CurrentLocation holds the current city/location.
var CurrentLocation string

// Forecast return a string value of the CurrentLocation + the CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
