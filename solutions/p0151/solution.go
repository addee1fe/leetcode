package solution

import (
	"strings"
)

func reverseWords(s string) string {
	w := strings.Fields(s)
	for i := len(w)/2 - 1; i >= 0; i-- {
		opp := len(w) - 1 - i
		w[i], w[opp] = w[opp], w[i]
	}
	return strings.Join(w, " ")
}
