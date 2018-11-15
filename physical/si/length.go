// Package si allows for the measurement and mathematical interaction between all things related to the physical world.
//
// All distances in this package are represented according to the International System of Units (SI). You may be more
// familiar with the name "metric". You can learn more about it on Wikipedia:
// https://en.wikipedia.org/wiki/International_System_of_Units
//
// Because this library uses SI, despite the fact that the author of this library is American, the names of all units
// adhere to their official SI names.
package si

import "math"

// Length represents the distance between two points. Its base unit is a metre.
type Length float64

// SI unites of length.
const (
	Nanometre  = Micrometre / 1000
	Micrometre = Millimetre / 1000
	Millimetre = Metre / 1000
	Centimetre = Metre / 100
	// Metre is the distance traveled by light in a vacuum in 1/299792458 seconds.
	Metre     Length = 1
	Kilometre        = 1000 * Metre
)

// Within judges whether or not two lengths are equal to a specified window.
func (l Length) Within(other, tolerance Length) bool {
	return math.Abs(float64(l-other)) <= float64(tolerance)
}
