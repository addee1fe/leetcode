package solution

func distributeCandies(candies []int) int {
	lc := len(candies) / 2
	set := make(map[int]struct{}, lc)
	for i := 0; i < len(candies) && len(set) < lc; i++ {
		set[candies[i]] = struct{}{}
	}
	return len(set)
}
