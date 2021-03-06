// Code generated by "stringer -type=CelestialBody"; DO NOT EDIT.

package moons

import "strconv"

const _CelestialBody_name = "SunMercuryVenusEarthMarsJupiterSaturnUranusNeptuneCeresUnknown"

var _CelestialBody_index = [...]uint8{0, 3, 10, 15, 20, 24, 31, 37, 43, 50, 55, 62}

func (i CelestialBody) String() string {
	if i < 0 || i >= CelestialBody(len(_CelestialBody_index)-1) {
		return "CelestialBody(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CelestialBody_name[_CelestialBody_index[i]:_CelestialBody_index[i+1]]
}
