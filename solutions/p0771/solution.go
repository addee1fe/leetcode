package solution

import "strings"

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0771
// BenchmarkNumJewelsInStones-8            10450988               104 ns/op
// BenchmarkRuneNumJewelsInStones-8         2985520               380 ns/op
// BenchmarkDummyNumJewelsInStones-8        2428962               494 ns/op
// BenchmarkSliceNumJewelsInStones-8       27384177                41.9 ns/op
// BenchmarkBitNumJewelsInStones-8         15349761                77.1 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0771     7.013s

// Final solution coz it easy to read and it understandable from first look
func numJewelsInStones(J string, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}
	res := 0
	for _, v := range strings.Split(J, "") {
		res += strings.Count(S, v)
	}
	return res
}

func runeNumJewelsInStones(J string, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}
	res := 0
	for _, v := range S {
		if strings.IndexRune(J, v) != -1 {
			res++
		}
	}
	return res
}

func dummyNumJewelsInStones(J string, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}
	s := 0
	m := make(map[rune]struct{})
	for _, v := range J {
		m[v] = struct{}{}
	}
	for _, v := range S {
		if _, ok := m[v]; ok {
			s++
		}
	}
	return s
}

func sliceNumJewelsInStones(J, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}
	res := 0
	sl := make([]int, 58) // for upper and lowercase
	for i := len(J) - 1; i >= 0; i-- {
		sl[J[i]-'A'] = 1
	}
	for i := len(S) - 1; i >= 0; i-- {
		res += sl[S[i]-'A']
	}
	return res
}

func bitNumJewelsInStones(J, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}
	res := 0
	bm := uint64(0)
	for _, v := range J {
		bm |= 1 << uint64(v-'A')
	}
	for _, v := range S {
		if bm&(1<<uint64(v-'A')) > 0 {
			res++
		}
	}
	return res
}
