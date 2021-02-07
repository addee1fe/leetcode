package solution

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	group := make([][]string, 0)
	groupMap := make(map[string][]string)
	for _, v := range strs {
		k := encodeString(v)
		groupMap[k] = append(groupMap[k], v)
	}
	for _, v := range groupMap {
		group = append(group, v)
	}
	return group
}

func encodeString(s string) string {
	var chars [26]int
	for i := range s {
		chars[s[i]-'a']++
	}
	return strings.Trim(strings.Replace(fmt.Sprint(chars), " ", "", -1), "[]")
}
