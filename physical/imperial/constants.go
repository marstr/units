package imperial

import (
	"github.com/marstr/units/physical/si"
)

// Imperial units of Length.
const (
	Thou = Inch / 1000
	Inch = Foot / 12
	Foot = Yard / 3
	Yard = 0.9144 * si.Metre // Made exact in 1959, before that this was closer to 0.914398415 metres.
	Mile = 1760 * Yard
)

// Imperial units of weight. Note: There is some conflation in the imperial system between mass and force. Essentially,
// "weight" is a measure of the force provided by gravity at the surface of the earth pulling on a set amount of mass.
// For that reason, we can compare weight and mass (unless you're an astronaut.)
const (
	Ounce = Pound / 16
	Pound = 0.45359237 * si.Kilogram // This is exact, by definition.
	Stone = 14 * Pound
	Ton   = 2240 * Pound
)

// Imperial units of volume.
const (
	FluidOunce = 28.4130625 * si.Millilitre
	Pint       = 20 * FluidOunce
	Quart      = 40 * FluidOunce
	Gallon     = 160 * FluidOunce
)
