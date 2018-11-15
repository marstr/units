package data

import (
	"fmt"
	"math"

	"github.com/marstr/units"
)

// Data holds the number of bits represented by a denomination of data size.
type Data uint64

// The number of bits represented by common denominations of data size.
const (
	Bit      = 1
	Byte     = 8 * Bit
	Kilobit  = 1024 * Bit
	Kilobyte = 1024 * Byte
	Megabit  = 1024 * Kilobit
	Megabyte = 1024 * Kilobyte
	Gigabit  = 1024 * Megabit
	Gigabyte = 1024 * Megabyte
	Terabit  = 1024 * Gigabit
	Terabyte = 1024 * Gigabyte
	Petabit  = 1024 * Terabit
	Petabyte = 1024 * Terabyte
	Exabit   = 1024 * Petabit
	Exabyte  = 1024 * Petabyte
	DataMax  = math.MaxUint64
)

// The suffixes associated with each level of Data capacity.
// For example a value of 12 Megabytes would would: 12MB
const (
	BitSuffix      = "b"
	ByteSuffix     = "B"
	KilobitSuffix  = units.KiloSuffix + BitSuffix
	KilobyteSuffix = units.KiloSuffix + ByteSuffix
	MegabitSuffix  = units.MegaSuffix + BitSuffix
	MegabyteSuffix = units.MegaSuffix + ByteSuffix
	GigabitSuffix  = units.GigaSuffix + BitSuffix
	GigabyteSuffix = units.GigaSuffix + ByteSuffix
	TerabitSuffix  = units.TeraSuffix + BitSuffix
	TerabyteSuffix = units.TeraSuffix + ByteSuffix
	PetabitSuffix  = units.PetaSuffix + BitSuffix
	PetabyteSuffix = units.PetaSuffix + ByteSuffix
	ExabitSuffix   = units.ExaSuffix + BitSuffix
	ExabyteSuffix  = units.ExaSuffix + ByteSuffix
)

// String formats a Data capacity as text, and returns it as encoded byte slice.
func (d Data) String() string {
	var suffix string
	var denominator Data
	switch {
	case d > Exabyte:
		suffix = ExabyteSuffix
		denominator = Exabyte
	case d > Petabyte:
		suffix = PetabyteSuffix
		denominator = Petabyte
	case d > Terabyte:
		suffix = TerabyteSuffix
		denominator = Terabyte
	case d > Gigabyte:
		suffix = GigabyteSuffix
		denominator = Gigabyte
	case d > Megabyte:
		suffix = MegabyteSuffix
		denominator = Megabyte
	case d > Kilobyte:
		suffix = KilobyteSuffix
		denominator = Kilobyte
	default:
		suffix = ByteSuffix
		denominator = Byte
	}

	return fmt.Sprintf("%0g%s", float32(d)/float32(denominator), suffix)
}
