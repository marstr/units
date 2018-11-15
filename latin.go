package units

// Latin prefixes
const (
	Yocto = Zepto / 1000
	Zepto = Atto / 1000
	Atto  = Femto / 1000
	Femto = Pico / 1000
	Pico  = Nano / 1000
	Nano  = Micro / 1000
	Micro = Milli / 1000
	Milli = 1 / 1000
	Centi = 1 / 100
	Deci  = 1 / 10
	Deca  = 10
	Hecto = 100
	Kilo  = 1000
	Mega  = 1000 * Kilo
	Giga  = 1000 * Mega
	Tera  = 1000 * Giga
	Peta  = 1000 * Tera
	Exa   = 1000 * Peta
	Zetta = 1000 * Exa
	Yotta = 1000 * Zetta
)

//
const (
	KiloSuffix  = "K"
	MegaSuffix  = "M"
	GigaSuffix  = "G"
	TeraSuffix  = "T"
	PetaSuffix  = "P"
	ExaSuffix   = "E"
	ZettaSuffix = "Z"
	YottaSuffix = "Y"
)
