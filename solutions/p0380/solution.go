package solution

import (
	"math/rand"
	"time"
)

// Have no idea how to do it better

// RandomizedSet ...
type RandomizedSet struct {
	v      map[int]int // key is value itself, value is index in values slice
	values []int
}

// Constructor returns new RandomizedSet
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

// Insert inserts a value to the set. Return true if the set did not already contain the specified element
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.v[val]; ok {
		return false
	}
	rs.values = append(rs.values, val)
	rs.v[val] = len(rs.values) - 1
	return true
}

// Remove removes a value from the set. Return true if the set contained the specified element
func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.v[val]; !ok {
		return false
	}
	i := rs.v[val]
	rs.values[i] = rs.values[len(rs.values)-1]
	rs.v[rs.values[i]] = i
	rs.values = rs.values[:len(rs.values)-1]
	delete(rs.v, val)
	return true
}

// GetRandom return random element from the set
func (rs *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rs.values[rand.Intn(len(rs.values))]
}
