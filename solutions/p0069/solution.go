package solution

import "math"

const delta = 1e-2

// Newton's method
func mySqrt(x int) int {
	z := float64(x)
	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-float64(x))/(2*z)
	}
	return int(z)
}
