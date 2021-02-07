package solution

func possibleBipartition(N int, dislikes [][]int) bool {
	if N <= 0 {
		return false
	}
	if len(dislikes) == 0 {
		return true
	}
	graph := make([][]int, N)
	for i := range graph {
		graph[i] = make([]int, N)
	}
	for _, v := range dislikes {
		graph[v[0]-1][v[1]-1] = 1
		graph[v[1]-1][v[0]-1] = 1
	}
	group := make([]int, N)
	for i := 0; i < N; i++ {
		if group[i] == 0 && !dfs(graph, group, i, 1) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, group []int, index, g int) bool {
	group[index] = g
	for i := range graph {
		if graph[i][index] == 1 {
			if group[i] == g {
				return false
			}
			if group[i] == 0 && !dfs(graph, group, i, -g) {
				return false
			}
		}
	}
	return true
}

// TODO: review this solution

// func possibleBipartition(N int, dislikes [][]int) bool {
//     g := make([][]int, N + 1)
//     c := make([]int, N + 1)

//     for _, edge := range dislikes {
//         g[edge[0]] = append(g[edge[0]], edge[1])
//         g[edge[1]] = append(g[edge[1]], edge[0])
//     }

//     for i := 1; i <= N; i++ {
//         if c[i] == 0 && !bipartite(g, c, i, 1) {
//             return false
//         }
//     }

//     return true
// }

// func bipartite(g [][]int, c []int, v, color int) bool {
//     if c[v] != 0 {
//         return c[v] == color
//     }

//     c[v] = color
//     nextColor := color ^ 3
//     for _, w := range g[v] {
//         if !bipartite(g, c, w, nextColor) {
//             return false
//         }
//     }

//     return true
// }
