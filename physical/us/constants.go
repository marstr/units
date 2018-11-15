// Package us allows for comparison/conversion from SI units to ones Americans will be more familiar with. The constants
// defined in this package are still fundamentally represented in SI.
package us

import (
	"github.com/marstr/units/physical/imperial"
	"github.com/marstr/units/physical/si"
)

// American units of length.
const (
	Inch = imperial.Inch
	Foot = imperial.Foot
	Yard = imperial.Yard
	Mile = imperial.Mile
)

// American units of weight. Note: There is some conflation in the us system between mass and force. Essentially,
// "weight" is a measure of the force provided by gravity at the surface of the earth pulling on a set amount of mass.
// For that reason, we can compare weight and mass (unless you're an astronaut.)
const (
	Grain   = Pound / 7000
	Ounce   = imperial.Ounce
	Pound   = imperial.Pound
	Ton     = 2000 * Pound
	LongTon = 2240 * Pound
)

// American units of volume.
const (
	Teaspoon   = Tablespoon / 3
	Tablespoon = FluidOunce / 2
	FluidOunce = 29.5735295625 * si.Millilitre
	Shot       = 1.5 * FluidOunce
	Cup        = 8 * FluidOunce
	Pint       = Quart / 2
	Quart      = Gallon / 4
	Gallon     = 128 * FluidOunce
)
