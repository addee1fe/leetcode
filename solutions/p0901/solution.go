package solution

// Ugly and slow solution

// StockSpanner ...
type StockSpanner struct {
	prices  []int
	weights []int
}

// Constructor ...
func Constructor() StockSpanner {
	return StockSpanner{}
}

// Next ...
func (sp *StockSpanner) Next(price int) int {
	w := 1
	for len(sp.prices) != 0 && sp.prices[len(sp.prices)-1] <= price {
		sp.prices = sp.prices[:len(sp.prices)-1]
		w += sp.weights[len(sp.weights)-1]
		sp.weights = sp.weights[:len(sp.weights)-1]
	}
	sp.prices = append(sp.prices, price)
	sp.weights = append(sp.weights, w)
	return w
}


//type StockSpanner struct {
//   stack [][2]int
//}
//func Constructor() StockSpanner {
//    return StockSpanner{make([][2]int, 0)}
//}
//func (this *StockSpanner) Next(price int) int {
//  count := 1
//    for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
//        count += this.stack[len(this.stack)-1][1]
//        this.stack = this.stack[:len(this.stack)-1]
//    }
//    this.stack = append(this.stack, [2]int{price, count})
//    return count
//}
