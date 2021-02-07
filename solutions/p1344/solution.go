package solution

import "math"

func angleClock(hour int, minutes int) float64 {
	a := math.Abs(30*float64(hour) - 5.5*float64(minutes))
	if a > 180 {
		return 360 - a
	}
	return a
}
