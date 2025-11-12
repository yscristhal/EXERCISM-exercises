//Package weather gives the forecast.
//of the weather.
package weather

//CurrentCondition it's a variable that tells how it´s the weather condition.
var CurrentCondition string
//CurrentLocation it´s a variable that tells from where it's being analised.
var	CurrentLocation  string

//Forecast gives a prediction of the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
