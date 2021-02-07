package solution

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges, indegree := make([][]int, numCourses), make([]int, numCourses)
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indegree[v[0]]++
	}
	queue, res := []int{}, []int{}
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		numCourses--
		res = append(res, cur)
		for _, v := range edges[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if numCourses != 0 {
		return []int{}
	}
	return res
}

// func findOrder(numCourses int, prerequisites [][]int) []int {
//     graph := make([][]int, numCourses)
//     indeg := make([]int, numCourses)
//     // Initialize the graph and indeg.
//     for _, prereq := range prerequisites {
//         from, to := prereq[1], prereq[0]
//         graph[from] = append(graph[from], to)
//         indeg[to]++
//     }
//     var res []int
//     // Top sort.
//     for len(res) < numCourses {
//         i := getZeroIdx(indeg)
//         if i == -1 {
//             return []int{}
//         }
//         res = append(res, i) // 拓扑顺序
//         for _, val := range graph[i] {
//             indeg[val]-- // 所有 i 开始的路径，终点入度减 1
//         }
//     }
//     return res
// }
// func getZeroIdx(nums []int) int {
//     for i, val := range nums {
//         if val == 0 {
//             nums[i] = -1
//             return i
//         }
//     }
//     return -1
// }
