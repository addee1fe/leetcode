package solution

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 || len(prerequisites) == 0 {
		return true
	}
	g := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}
	s := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if dfs(g, i, s) == false {
			return false
		}
	}
	return true
}

func dfs(graph map[int][]int, c int, state []int) bool {
	if state[c] == 1 {
		return true
	}
	if state[c] == 2 {
		return false
	}
	state[c] = 2
	for _, v := range graph[c] {
		if dfs(graph, v, state) == false {
			return false
		}
	}
	state[c] = 1
	return true
}

// func canFinish(numCourses int, prerequisites [][]int) bool {

//     ans := 0
//     incomingEdges := make([]int, numCourses)
//     m := make(map[int][]int, numCourses)

//     for _, p := range prerequisites {
//         a,b := p[0], p[1]
//         m[a] = append(m[a], b)
//         incomingEdges[b]++
//     }

//     var q []int

//     for course := range incomingEdges {
//         if incomingEdges[course] == 0 {
//             q = append(q, course)
//         }
//     }

//     for len(q) > 0 {

//         num := q[0]
//         q = q[1:]
//         ans++

//         for _, x := range m[num] {
//             incomingEdges[x]--
//             if incomingEdges[x] == 0 {
//                 q = append(q, x)
//             }
//         }

//     }

//     return ans == numCourses

// }

// Kahn’s algorithm for Topological Sorting (queue)
// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	adj := map[int][]int{}
// 	indegree := make([]int, numCourses)
// 	for _, p := range prerequisites {
// 		adj[p[1]] = append(adj[p[1]], p[0])
// 		indegree[p[0]]++
// 	}
// 	queue := []int{}
// 	for n, in := range indegree {
// 		if in == 0 {
// 			queue = append(queue, n)
// 		}
// 	}
// 	for len(queue) > 0 {
// 		n := queue[0]
// 		numCourses--
// 		for _, nei := range adj[n] {
// 			indegree[nei]--
// 			if indegree[nei] == 0 {
// 				queue = append(queue, nei)
// 			}
// 		}
// 		queue = queue[1:]
// 	}
// 	return numCourses == 0
// }
