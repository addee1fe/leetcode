package solution

import "strings"

func intToRoman(num int) string {
	v := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	m := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb strings.Builder
	for i := 0; num != 0; {
		for v[i] > num {
			i++
		}
		num -= v[i]
		sb.WriteString(m[i])
	}
	return sb.String()
}

// Another solution
// func intToRoman(num int) string {
//     M := [...]string{"", "M", "MM", "MMM"}
//     C := [...]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
//     X := [...]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
//     I := [...]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
//     return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
// }
