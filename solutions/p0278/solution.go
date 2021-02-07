package solution

// Dummy stub
func isBadVersion(version int) bool {
	if version == 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	s, e := 0, n
	for s < e {
		m := s + (e-s)/2
		if !isBadVersion(m) {
			s = m + 1
		} else {
			e = m
		}
	}
	return s
}
