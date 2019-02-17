package space

// Notes:
// Earth: orbital period 365.25 Earth days, or 31557600 seconds
// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years

type Planet string

var earthSecondsToYear = 31557600.0
var planetSecondsMultiplier = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846, // Lol
	"Neptune": 164.79132,
}

// Note: expected answer rounds to two digits
func Age(seconds float64, planet string) float64 {
	return seconds / (earthSecondsToYear * planetSecondsMultiplier[planet])
}
