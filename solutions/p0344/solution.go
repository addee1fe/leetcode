package solution

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0344
// BenchmarkReverseString-8                148144622                8.04 ns/op
// BenchmarkReverseStringOneIndex-8        93635186                11.6 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0344     3.219s

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
func reverseStringOneIndex(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
