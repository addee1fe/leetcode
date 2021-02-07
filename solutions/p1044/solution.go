package solution

// Please note
// This is not mine solution

func pow(a, x int) int {
	switch x {
	case 0:
		return 1
	case 1:
		return a
	}
	res := a
	for i := 0; i < x-1; i++ {
		res *= a
	}
	return res
}

func find(S string, l int) string {
	var (
		hash     int = 0
		prime    int = 1001
		primePow int = pow(prime, l-1)
		hashes       = make(map[int]int, len(S))
	)
	for i := 0; i < l; i++ {
		hash = hash*prime + int(S[i])
	}
	hashes[hash] = 0

	for i := 1; i <= len(S)-l; i++ {
		hash -= int(S[i-1]) * primePow
		hash *= prime
		hash += int(S[i+l-1])
		if pos, found := hashes[hash]; found {
			return S[pos : pos+l]
		}
		hashes[hash] = i
	}
	return ""
}

func longestDupSubstring(S string) string {
	l := 0
	r := len(S) - 1
	res := ""
	for l <= r {
		mid := l + (r-l)/2
		if curr := find(S, mid); len(curr) > len(res) {
			res = curr
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
