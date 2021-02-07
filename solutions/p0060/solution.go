package solution

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	ns := make([]string, n)
	for i := 0; i < n; i++ {
		ns[i] = strconv.Itoa(i + 1)
	}
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = i * f[i-1]
	}
	var sb strings.Builder
	k--
	for ; n > 0; n-- {
		pos := k / f[n-1]
		sb.WriteString(ns[pos])
		ns = append(ns[:pos], ns[pos+1:]...)
		k = k % f[n-1]
	}
	return sb.String()
}
