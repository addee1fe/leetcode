package solution

func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)
	if l == 0 || l == 1 {
		return false
	}
	if l == 2 {
		return true
	}
	p, q := coordinates[0], coordinates[1]
	for i := 2; i < l; i++ {
		c := coordinates[i]
		if (c[0]-p[0])*(q[1]-p[1]) != (c[1]-p[1])*(q[0]-p[0]) {
			return false
		}
	}
	return true
}
