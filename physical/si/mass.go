package si

import "fmt"

// Mass represents the amount of matter. Its base unit is a Kilogram.
type Mass float64

// SI units of mass.
const (
	Milligram      = Gram / 1000
	Gram           = Kilogram / 1000
	Kilogram  Mass = 1
)

// Within judges whether two scalar masses are equal to one another, given a window of acceptable deviation.
func (m Mass) Within(other, tolerance Mass) bool {
	return within(
		float64(m),
		float64(other),
		float64(tolerance))
}

func (m Mass) String() string {
	return fmt.Sprintf("%gg", m)
}
