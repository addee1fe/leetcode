package solution

const (
	// RED ...
	RED = 1
	// BLUE ...
	BLUE = 2
)

// I'm suck in graph algorithms so this solution was copy-pasted
// TODO: review this solution
func isBipartite(graph [][]int) bool {
	color := make(map[int]int)
	for i := range graph {
		if _, ok := color[i]; ok {
			continue
		}
		color[i] = RED // we found an unconnected one

		// walk everything from here
		q := []int{i}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			thisColor := color[n]
			nextColor := BLUE
			if thisColor == BLUE {
				nextColor = RED
			}
			for _, neighbor := range graph[n] {
				if cc, ok := color[neighbor]; ok {
					if cc == thisColor {
						return false
					}
					continue
				}
				color[neighbor] = nextColor
				q = append(q, neighbor)
			}
		}
	}
	return true
}
