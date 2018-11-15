package si

// Temperature holds a value indicating degrees Kelvin. Note: 0K is defined to be the absolute absence of thermal
// activity. Therefor, negative values of Temperature do not make any sense.
type Temperature float64

const (
	celsiusKelvinOffset    = 273.16
	fahrenheitKelvinOffset = 459.67
)

// NewTemperatureFromCelsius creates a Temperature from a scalar representing the temperature in Celsius.
func NewTemperatureFromCelsius(t float64) Temperature {
	return Temperature(t) + celsiusKelvinOffset
}

// NewTemperatureFromFahrenheit creates a Temperature from a scalar representing the temperature in Fahrenheit.
func NewTemperatureFromFahrenheit(t float64) Temperature {
	return (Temperature(t) + fahrenheitKelvinOffset) * 5 / 9
}

// Celsius returns a scalar representing the temperature in Celsius.
func (t Temperature) Celsius() float64 {
	return float64(t) - celsiusKelvinOffset
}

// Fahrenheit returns a scalar representing the temperature in Fahrenheit.
func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 - fahrenheitKelvinOffset
}
