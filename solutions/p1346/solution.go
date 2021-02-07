package solution

func checkIfExist(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	m := make(map[int]bool)
	for _, v := range arr {
		if m[v*2] != false {
			return true
		}
		if v%2 == 0 && m[v/2] != false {
			return true
		}
		m[v] = true
	}
	return false
}
