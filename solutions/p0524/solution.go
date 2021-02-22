package solution

import "strings"

func findLongestWord(s string, d []string) string {
	var max string
	for _, v := range d {
		if isSubsequence(v, s) {
			if len(v) > len(max) || len(v) == len(max) && strings.Compare(v, max) < 0 {
				max = v
			}
		}
	}
	return max
}

func isSubsequence(x, y string) bool {
	var j int
	for i := 0; i < len(y) && j < len(x); i++ {
		if y[i] == x[j] {
			j++
		}

	}
	return j == len(x)
}
