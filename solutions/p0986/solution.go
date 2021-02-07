package solution

func intervalIntersection(A [][]int, B [][]int) [][]int {
	m, n := len(A), len(B)
	if m == 0 || n == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	for i, j := 0, 0; i < m && j < n; {
		a := A[i]
		b := B[j]
		sMax := a[0]
		if sMax < b[0] {
			sMax = b[0]
		}
		eMin := a[1]
		if eMin > b[1] {
			eMin = b[1]
		}
		if eMin >= sMax {
			res = append(res, []int{sMax, eMin})
		}
		if a[1] == eMin {
			i++
		}
		if b[1] == eMin {
			j++
		}
	}
	return res
}
