package solution

// also this can be solved with Trie
func minimumLengthEncoding(words []string) int {
	good := make(map[string]struct{}, 0)
	for _, v := range words {
		good[v] = struct{}{}
	}
	for _, w := range words {
		for k := 1; k < len(w); k++ {
			delete(good, w[k:])
		}
	}
	var res int
	for k := range good {
		res += len(k) + 1
	}
	return res
}
