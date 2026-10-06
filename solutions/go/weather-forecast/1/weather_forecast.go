// Package weather provides weather forecast functionality.
package weather

var (
	// CurrentCondition holds the current weather condition.
	CurrentCondition string
	// CurrentLocation holds the current location.
	CurrentLocation string
)

// Forecast returns the current weather forecast for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
