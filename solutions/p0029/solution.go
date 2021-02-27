package solution

import "math"

// This is truly hardcoded solution and probably not one what would expect on interview
// But I do it just for fun, so i don't care
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return int(int32(dividend) - 1)
	}
	return dividend / divisor
}
