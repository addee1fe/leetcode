package solution

func hasAllCodes(s string, k int) bool {
	need := 1 << k
	seen := make([]bool, need)
	mask := need - 1
	var curr int
	for i := 0; i < len(s); i++ {
		curr = ((curr << 1) | int(s[i]-'0')) & mask
		if i >= k-1 && !seen[curr] {
			seen[curr] = true
			need--
			if need == 0 {
				return true
			}
		}
	}
	return false
}
