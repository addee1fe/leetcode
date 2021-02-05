package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	var max int
	for _, v := range candies {
		if max < v {
			max = v
		}
	}
	for i := range res {
		res[i] = candies[i]+extraCandies >= max
	}
	return res
}
