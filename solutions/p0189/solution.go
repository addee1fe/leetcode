package solution

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0189
// BenchmarkRotate-8               51830130                22.0 ns/op
// BenchmarkRotateSecond-8         12453039                92.8 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0189     2.894s

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, s, e int) {
	for s < e {
		tmp := nums[s]
		nums[s] = nums[e]
		nums[e] = tmp
		s++
		e--
	}
}

func rotateSecond(nums []int, k int) {
	n := len(nums)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[(i+k)%n] = nums[i]
	}
	copy(nums, r)
}
