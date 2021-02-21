package solution

func brokenCalc(x int, y int) int {
	var ans int
	for y > x {
		ans++
		if y%2 == 1 {
			y++
		} else {
			y /= 2
		}
	}
	return ans + x - y
}
