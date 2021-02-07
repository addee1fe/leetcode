package solution

import (
	"sort"
	"strings"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0451
// BenchmarkFrequencySortMap-8       783615              1514 ns/op
// BenchmarkFrequencySort-8         1000000              1076 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0451     3.311s

func frequencySort(s string) string {
	if len(s) <= 1 {
		return s
	}
	fr := [256]int{}
	sl := make([]string, 0)
	for _, v := range s {
		fr[v]++
	}
	for i, v := range fr {
		if v != 0 {
			sl = append(sl, strings.Repeat(string(rune(i)), v))
		}
	}
	sort.SliceStable(sl, func(i, j int) bool { return len(sl[i]) > len(sl[j]) })
	return strings.Join(sl, "")
}

func frequencySortMap(s string) string {
	if len(s) <= 1 {
		return s
	}
	fr := make(map[rune]int)
	sl := make([]string, 0)
	for _, v := range s {
		fr[v]++
	}
	for i, v := range fr {
		sl = append(sl, strings.Repeat(string(i), v))
	}
	sort.SliceStable(sl, func(i, j int) bool { return len(sl[i]) > len(sl[j]) })
	return strings.Join(sl, "")
}
