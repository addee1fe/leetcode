package solution

import (
	"math"
	"math/rand"
)

// http://www.anderswallin.net/2009/05/uniform-random-points-in-a-circle-using-polar-coordinates/

type Solution struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, x_center: x_center, y_center: y_center}
}

func (s *Solution) RandPoint() []float64 {
	l := math.Sqrt(rand.Float64()) * s.radius
	d := math.Pi * 2 * rand.Float64()
	x := s.x_center + l*math.Cos(d)
	y := s.y_center + l*math.Sin(d)
	return []float64{x, y}
}
