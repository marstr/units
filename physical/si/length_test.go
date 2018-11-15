package si_test

import (
	"fmt"
	"github.com/marstr/units/physical/si"
)

func ExampleLength_Within() {
	raceDistance := 5 * si.Kilometre
	distanceRan := 4.99999 * si.Kilometre

	fmt.Println(distanceRan == raceDistance)
	fmt.Println(distanceRan.Within(raceDistance, si.Metre))

	// Output:
	// false
	// true
}
