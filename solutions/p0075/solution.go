package solution

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0075
// BenchmarkSortColors-8           112505938               10.7 ns/op
// BenchmarkSortColorsDummy-8      66231770                17.7 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0075     3.240s

// Look to Dutch national flag problem for more info

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	min, max := 0, len(nums)-1
	for i := 0; i <= max; {
		n := nums[i]
		switch n {
		case 0:
			nums[i], nums[min] = nums[min], 0
			i++
			min++
		case 1:
			i++
		case 2:
			nums[i], nums[max] = nums[max], 2
			max--
		}
	}
}

func sortColorsDummy(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
