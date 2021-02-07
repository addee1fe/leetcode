package solution

import (
	"math"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}
	h := make([][2]float64, len(points))
	res := make([][]int, K)
	for i, v := range points {
		h[i] = [2]float64{math.Hypot(float64(v[0]), float64(v[1])), float64(i)}
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i][0] < h[j][0]
	})
	for i := 0; i < K; i++ {
		res[i] = points[int(h[i][1])]
	}
	return res
}

// review 1
// func kClosest(points [][]int, K int) [][]int {
// 	l := 0
// 	r := len(points) - 1
// 	var pivot int
// 	for l < r {
// 		pivot = partition(points, l, r)
// 		if pivot == K {
// 			break
// 		}
// 		if pivot < K {
// 			l = pivot + 1
// 		} else {
// 			r = pivot - 1
// 		}
// 	}
// 	return points[:K]
// }
// func partition(points [][]int, l, r int) int {
//     pivot := points[r]
//     i := l-1
//     for j:=l; j<r; j++ {
//         if isShorter(points[j], pivot) {
//             i++
//             points[i], points[j] = points[j], points[i]
//         }
//     }
//     points[i+1], points[r] = points[r], points[i+1]
//     return i+1
// }
// func isShorter(a, b []int) bool {
//     return (a[0]*a[0] + a[1]*a[1]) < (b[0]*b[0]+ b[1]*b[1])
// }

// review 2
// func kClosest(points [][]int, K int) [][]int {
// 	quickSelect(points, 0, len(points)-1, K)
// 	res := make([][]int, 0)
// 	for i := 0; i < K; i++ {
// 		res = append(res, points[i])
// 	}
// 	return res
// }
// func quickSelect(a [][]int, l, r, k int) {
// 	pivot := l
// 	for i := l; i < r; i++ {
// 		if (a[i][0]*a[i][0] + a[i][1]*a[i][1]) <= (a[r][0]*a[r][0] + a[r][1]*a[r][1]) {
// 			a[i], a[pivot] = a[pivot], a[i]
// 			pivot += 1
// 		}
// 	}
// 	a[pivot], a[r] = a[r], a[pivot]
// 	if pivot+1 == k {
// 		return
// 	} else if pivot+1 > k {
// 		quickSelect(a, l, pivot-1, k)
// 	} else {
// 		quickSelect(a, pivot+1, r, k)
// 	}
// }

// review 3
// func kClosest(points [][]int, K int) [][]int {
//     if len(points) < K {
//         return nil
//     }
//     search(points, K)
//     return points[:K]
// }
// func search(a [][]int, k int) {
//     if len(a) <= k {
//         return
//     }
//     left, pivot := 0, norm(a[0])
//     for j := left+1; j < len(a); j++ {
//         if norm(a[j]) <= pivot {
//             left++
//             if left < j {
//                 a[left], a[j] = a[j], a[left]
//             }
//         }
//     }
//     if left > 0 {
//         a[left], a[0] = a[0], a[left]
//     }
//     if k < left {
//         search(a[:left], k)
//     } else if right := left+1; k > right {
//         search(a[right:], k-right)
//     }
// }
// func norm(p []int) int {
//     return p[0]*p[0] + p[1]*p[1]
// }
