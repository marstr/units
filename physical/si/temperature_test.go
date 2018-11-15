package si_test

import (
	"fmt"

	"github.com/marstr/units/physical/si"
)

func ExampleTemperature_Fahrenheit() {
	fmt.Printf("%.0f\n", si.Freezing.Fahrenheit())
	fmt.Printf("%.0f\n", si.Boiling.Fahrenheit())
	// Output:
	// 32
	// 212
}

func ExampleTemperature_Celsius() {
	fmt.Printf("%.0f\n", si.Freezing.Celsius())
	fmt.Printf("%.0f\n", si.Boiling.Celsius())
	// Output:
	// 0
	// 100
}

func ExampleTemperature_conversion() {
	subject := si.NewTemperatureFromCelsius(100)
	fmt.Printf("%.0f\n", subject.Fahrenheit())
	fmt.Println(subject)

	// Output:
	// 212
	// 373.2K
}
