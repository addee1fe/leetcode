package solution

func validUtf8(data []int) bool {
	var nToProc int
	m1 := 1 << 7
	m2 := 1 << 6
	for i := range data {
		if nToProc == 0 {
			m := 1 << 7
			for (m & data[i]) != 0 {
				nToProc++
				m >>= 1
			}
			if nToProc == 0 {
				continue
			}
			if nToProc > 4 || nToProc == 1 {
				return false
			}
		} else {
			if !((data[i]&m1) != 0 && (m2&data[i]) == 0) {
				return false
			}
		}
		nToProc--
	}
	return nToProc == 0
}
