package solution

// I got the basic idea, but note
// this solution is not mine

// FreqStack ...
type FreqStack struct {
	stacks [][]int
	freq   map[int]int
}

// Constructor ...
func Constructor() FreqStack {
	return FreqStack{
		stacks: [][]int{},
		freq:   make(map[int]int),
	}
}

// Push ...
func (fs *FreqStack) Push(x int) {
	fs.freq[x]++
	n := fs.freq[x]
	if n > len(fs.stacks) {
		fs.stacks = append(fs.stacks, []int{})
	}
	fs.stacks[n-1] = append(fs.stacks[n-1], x)
}

// Pop ...
func (fs *FreqStack) Pop() int {
	m := len(fs.stacks)
	n := len(fs.stacks[m-1])
	x := fs.stacks[m-1][n-1]
	fs.stacks[m-1] = fs.stacks[m-1][:n-1]
	if len(fs.stacks[m-1]) == 0 {
		fs.stacks = fs.stacks[:m-1]
	}
	fs.freq[x]--
	return x
}
