// Package weather provides tools to forecase weather.
package weather

// CurrentCondition represents what the current weather condition is.
var CurrentCondition string
// CurrentLocation represents the geographical location for the weather forecase.
var CurrentLocation string

// Forecast returns a string indicating what the weather prediction is. It takes the city
// and condition in as parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
