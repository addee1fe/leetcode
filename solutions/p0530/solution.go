package solution

import (
	"regexp"
	"strings"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0530
// BenchmarkDetectCapitalUse-8              9507218               117 ns/op
// BenchmarkDetectCapitalUseLoop-8         73102126                16.0 ns/op
// BenchmarkDetectCapitalUseRegex-8          209401              5414 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0530     3.757s

func detectCapitalUse(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}

func detectCapitalUseRegex(word string) bool {
	matched, _ := regexp.MatchString(`^[A-Z]*$|^[a-z]*$|^[A-Z][a-z]*$`, word)
	return matched
}
func detectCapitalUseLoop(word string) bool {
	var c int
	for _, v := range word {
		if v >= 'A' && v <= 'Z' {
			c++
		}
	}
	if c == len(word) || c == 0 {
		return true
	}
	if c == 1 && word[0] >= 'A' && word[0] < 'Z' {
		return true
	}
	return false
}
