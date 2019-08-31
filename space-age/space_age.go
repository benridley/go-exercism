package space

// Planet is a string represenitng the name of a planet in Earth's solar system.
type Planet string

var planetYears = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const secondsInEarthYear float64 = 31557600

// Age returns how old someone is on another planet given their age in seconds on Earth.
func Age(ageInSeconds float64, planetName Planet) float64 {
	return (ageInSeconds / secondsInEarthYear) / planetYears[planetName]
}
