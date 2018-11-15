package si

import (
	"fmt"
	"math"
)

// Volume represents the amount of space a three dimensional object populates.
type Volume float64

// SI units of volume.
const (
	Millilitre        = Litre / 1000
	Centilitre        = Litre / 100
	Litre      Volume = 1
	Kilolitre         = 1000 * Litre
)

// Within judges whether or not two volumes are equal, given a window of acceptable deviation.
func (v Volume) Within(other, tolerance Volume) bool {
	return within(
		float64(v),
		float64(other),
		float64(tolerance))
}

func within(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func (v Volume) String() string {
	return fmt.Sprintf("%gL", v)
}
