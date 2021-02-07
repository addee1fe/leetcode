package solution

import "sort"

// For better solution take a look at Hierholzer Algorithm and Euler Path

// Please note. This is not mine solution

func findItinerary(tickets [][]string) []string {
	hash := map[string][]string{}
	hash2 := map[string][]bool{}
	for _, tkt := range tickets {
		hash[tkt[0]] = append(hash[tkt[0]], tkt[1])
		hash2[tkt[0]] = append(hash2[tkt[0]], false)
	}
	for _, v := range hash {
		sort.Strings(v)
	}
	current := "JFK"
	travel := []string{"JFK"}
	dfs(current, len(tickets), hash, hash2, &travel)
	return travel
}

func dfs(current string, left int, hash map[string][]string, hash2 map[string][]bool, travel *[]string) bool {
	if left == 0 {
		return true
	}
	if len(hash[current]) == 0 {
		return false
	}
	for i, v := range hash2[current] {
		if v == false {
			hash2[current][i] = true
			*travel = append(*travel, hash[current][i])
			if dfs(hash[current][i], left-1, hash, hash2, travel) == true {
				return true
			}
			hash2[current][i] = false
			*travel = (*travel)[:len(*travel)-1]
		}
	}
	return false
}

// DFS Solution
// func findItinerary(tickets [][]string) []string {
//     m := make(map[string][]string, len(tickets)+1)
//     var ans []string
//     for _, t := range tickets {
//         m[t[0]] = append(m[t[0]], t[1])
//     }
//     for k := range m {
//         sort.Strings(m[k])
//     }
//     DFS("JFK", m, &ans)
// 	// reverse ans array
//     i, j := 0, len(ans)-1
// 	for i < j {
// 		ans[i], ans[j] = ans[j], ans[i]
// 		i++
// 		j--
// 	}
//     return ans
// }
// func DFS(start string, m map[string][]string, ans *[]string) {
//     for len(m[start]) > 0 {
//         cur := m[start][0]
//         m[start] = m[start][1:]
//         DFS(cur, m, ans)
//     }
//     *ans = append(*ans, start)
// }

// Backtrack solution
// func findItinerary(tickets [][]string) []string {
//     sort.Slice(tickets, func(i, j int) bool {
//         return tickets[i][1] < tickets[j][1]
//     })
//     return backtrack(tickets, []string{"JFK"}, len(tickets) + 1)
// }

// func backtrack(tickets [][]string, cur []string, needed int) []string {
//     if len(cur) == needed {
//         return cur
//     }
//     for i, pair := range tickets {
//         from, to := pair[0], pair[1]
//         if from == cur[len(cur) - 1] {
//             ticketsCopy := make([][]string, len(tickets))
//             copy(ticketsCopy, tickets)
//             ticketsCopy = append(ticketsCopy[:i], ticketsCopy[i+1:]...)
//             cur = append(cur, to)
//             res := backtrack(ticketsCopy, cur, needed)
//             if len(res) > 0 {
//                 return res
//             }
//             cur = cur[:len(cur)-1]
//         }
//     }
//     return []string{}
// }

// Another interesting solution
// func findItinerary(tickets [][]string) []string {
//     ans, paths := make([]string, 0), make(map[string][]string)
//     for _, v := range tickets { paths[v[0]] = append(paths[v[0]], v[1]) }
//     for _, v := range paths { sort.Strings(v) }
//     helper("JFK", &ans, paths)
//     return reverse(ans)
// }

// func helper(src string, ans *[]string, paths map[string][]string) {
//     for len(paths[src]) > 0 {
//         d := paths[src][0]
//         paths[src] = paths[src][1:]
//         helper(d, ans, paths)
//     }
//     *ans = append(*ans, src)
// }

// func reverse(s []string) []string {
//     for i, j := 0, len(s) - 1; i < j; { s[i], s[j] = s[j], s[i]; i++; j-- }
//     return s
// }
