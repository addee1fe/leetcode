package solution

import "math"

// MinStack ...
type MinStack struct {
	stack []int
	min   int
}

// Constructor ...
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		math.MaxInt32,
	}
}

// Push ...
func (ms *MinStack) Push(x int) {
	if x < ms.min {
		ms.min = x
	}
	ms.stack = append(ms.stack, x)
}

// Pop ...
func (ms *MinStack) Pop() {
	if ms.stack[len(ms.stack)-1] == ms.min && len(ms.stack) > 0 {
		ms.min = math.MaxInt32
		for i := 0; i < len(ms.stack)-1; i++ {
			if ms.stack[i] < ms.min {
				ms.min = ms.stack[i]
			}
		}
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
}

// Top ...
func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[len(ms.stack)-1]
}

// GetMin ...
func (ms *MinStack) GetMin() int {
	return ms.min
}
