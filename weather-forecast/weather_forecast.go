/* 
  Package weather
  This package provides a function that returns the current weather condition for a given city.
*/
package weather

// CurrentCondition holds the current weather condition for a given city.
var CurrentCondition string

// CurrentLocation holds the current location for which the weather condition is being reported.
var CurrentLocation string

/* 
  Forecast() returns the current weather condition for a given city.
  It takes two arguments: city and condition.
  It returns a string.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
