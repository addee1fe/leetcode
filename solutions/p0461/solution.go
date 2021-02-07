package solution

func hammingDistance(x int, y int) int {
	return pop(x ^ y)
}

func pop(x int) int {
	sum := x
	for x != 0 {
		x >>= 1
		sum -= x
	}
	return sum
}

// Or you can do it in more exotic way :)
// z := int64(x ^ y)
// n := strconv.FormatInt(z, 2)
// return strings.Count(n,"1")
