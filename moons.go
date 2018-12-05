package moons

import (
	"strings"

	"fmt"
)

// CelestialBody represents a natural body in our solar system.
//go:generate stringer -type=CelestialBody
type CelestialBody int

// Constants with the most prominent celestial bodies which this
// library understands.
const (
	Sun CelestialBody = iota
	Mercury
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune

	// Unknown is used when converting from string to CelestialBody
	// and shows that the returned body is not recognized.
	Unknown
)

// Count returns the number of moons a CelestialBody planet has.
func Count(planet CelestialBody) (int, error) {
	if planet == Unknown {
		return 0, fmt.Errorf("unknown CelestialBody: %s", planet)
	}

	count, ok := moonsCount[planet]
	if !ok {
		return 0, fmt.Errorf("%s is not a planet", planet)
	}
	return count, nil
}

// BodyFromString returns a CelestialBody which corresponds to
// the string `name`.
func BodyFromString(name string) CelestialBody {
	body, ok := nameToBody[strings.ToLower(name)]
	if !ok {
		return Unknown
	} else {
		return body
	}

	return Unknown
}

var moonsCount = map[CelestialBody]int{
	Mercury: 0,
	Venus:   0,
	Earth:   1,
	Jupiter: 67,
	Saturn:  62,
	Uranus:  27,
	Neptune: 14,
}

var nameToBody = map[string]CelestialBody{
	"sun":     Sun,
	"mercury": Mercury,
	"venus":   Venus,
	"earth":   Earth,
	"jupiter": Jupiter,
	"saturn":  Saturn,
	"uranus":  Uranus,
	"neptune": Neptune,
}
