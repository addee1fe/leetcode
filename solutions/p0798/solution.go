package solution

func bestRotation(A []int) int {
	n := len(A)
	bad := make([]int, n)
	for i := 0; i < n; i++ {
		l := (i - A[i] + 1 + n) % n
		r := (i + 1) % n
		bad[l]--
		bad[r]++
		if l > r {
			bad[0]--
		}
	}
	best := -n
	ans, cur := 0, 0
	for i := 0; i < n; i++ {
		cur += bad[i]
		if cur > best {
			best = cur
			ans = i
		}
	}
	return ans
}

// another solution
// func bestRotation(a []int) int {
//     n := len(a)
//     change := make([]int, n, n)
//     for i := range a {
//         change[(i - a[i] + 1 + n) % n] -= 1
//     }
//     var mi int
//     for i := 1; i < n; i++ {
//         change[i] += change[i - 1] + 1
//         if change[i] > change[mi] {
//             mi = i
//         }
//     }
//     return mi;
// }
