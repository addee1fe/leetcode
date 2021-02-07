package solution

import "sort"

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 1 {
		return people
	}
	res := make([][]int, len(people))
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	indexes := make([]int, len(people))
	for i := range indexes {
		indexes[i] = i
	}
	for _, v := range people {
		res[indexes[v[1]]] = v
		indexes = append(indexes[:v[1]], indexes[v[1]+1:]...)
	}
	return res
}
