package solution

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	var sb strings.Builder
	i, j, c := len(a)-1, len(b)-1, 0 // c - carry
	for i >= 0 || j >= 0 {
		sum := c
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		fmt.Fprintf(&sb, "%d", sum%2)
		c = sum / 2
	}
	if c != 0 {
		fmt.Fprintf(&sb, "%d", c)
	}
	return reverse(sb.String())
}

func reverse(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
