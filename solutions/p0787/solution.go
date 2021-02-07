package solution

// TODO: rewrite solution to my own
// Please note, this is not my solution

import "container/heap"

type Item struct {
	Cost     int // The priority of the item in the queue.
	Location int // The value of the item; arbitrary.
	Steps    int // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, cost int, location int, steps int) {
	item.Cost = cost
	item.Location = location
	item.Steps = steps
	heap.Fix(pq, item.Index)
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {

	// put flights into a hashtable for lookup
	flightsMap := make(map[int][][]int)
	for _, flight := range flights {
		from := flight[0]
		if _, x := flightsMap[from]; x {
			flightsMap[from] = append(flightsMap[from], flight)
		} else {
			flightsMap[from] = [][]int{flight}
		}
	}
	// dijkstra's
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, src, 0, 0})
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		// this is the trickiest part 1: why first item found is the result?
		if item.Location == dst {
			return item.Cost
		}

		// the trickiest part 2: why dont we use a hastable to avoid visiting again?

		if item.Steps-1 == K {
			continue
		}

		if _, x := flightsMap[item.Location]; x {
			candidates := flightsMap[item.Location]
			for _, can := range candidates {
				i := pq.Len()
				heap.Push(pq, &Item{item.Cost + can[2], can[1], item.Steps + 1, i})
			}
		}
	}
	return -1
}

// const INFINITY = 1e9

// func min(a, b int) int {
//     if a < b {
//         return a
//     } else {
//         return b
//     }
// }

// func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
//     // D(len, j) 表示路径长度至多为 len 目的地为 j 的最短路径
//     // D(K+1, dst) 为最终答案
//     // 初始 D(0, src) = 0
//     // 转移方程：D(len, j) = min{D(len-1, j), D(len-1, u) + w}, 对于每一条边 u, j, w
//     D := make([]int, n)
//     for i := 0; i < n; i++ {
//         D[i] = INFINITY
//     }
//     D[src] = 0
//     nextD := make([]int, n)
//     for i := 0; i <= K; i++ {
//         copy(nextD, D)
//         for j := 0; j < len(flights); j++ {
//             u, v, w := flights[j][0], flights[j][1], flights[j][2]
//             if D[u] + w < D[v] {
//                 nextD[v] = min(nextD[v], D[u] + w)
//             }
//         }
//         copy(D, nextD)
//     }
//     if D[dst] == INFINITY {
//         return -1
//     } else {
//         return D[dst]
//     }
// }
