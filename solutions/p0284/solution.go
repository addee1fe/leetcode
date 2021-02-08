package solution

// TODO:
// Refactor entire solution
// Add tests

// // Iterator ...
// type Iterator struct {
// }

// func (i *Iterator) hasNext() bool {
// 	// Returns true if the iteration has more elements.
// 	return false
// }

// func (i *Iterator) next() int {
// 	// Returns the next element in the iteration.
// 	return 0
// }

// // PeekingIterator ...
// type PeekingIterator struct {
// 	vals []int
// 	p    int
// }

// // Constructor ...
// func Constructor(iter *Iterator) *PeekingIterator {
// 	var pi PeekingIterator
// 	for iter.hasNext() {
// 		pi.vals = append(pi.vals, iter.next())
// 	}
// 	pi.p = 0
// 	return &pi
// }

// func (pi *PeekingIterator) hasNext() bool {
// 	return pi.p < len(pi.vals)
// }

// func (pi *PeekingIterator) next() int {
// 	pi.p++
// 	return pi.vals[pi.p-1]
// }

// func (pi *PeekingIterator) peek() int {
// 	return pi.vals[pi.p]
// }
