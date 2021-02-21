package solution

func romanToInt(s string) int {
	if len(s) == 0 {
		return -1
	}
	rti := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := rti[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if rti[s[i]] >= rti[s[i+1]] {
			res += rti[s[i]]
		} else {
			res -= rti[s[i]]
		}
	}
	return res
}
