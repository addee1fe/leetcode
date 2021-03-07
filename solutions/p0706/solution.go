package solution

// NOTE: this is the extremely easy, incorrect solution and I will replace it as soon
// as I will be not tired af

// MyHashMap ...
type MyHashMap struct {
	m map[int]int
}

// Constructor ...
func Constructor() MyHashMap {
	return MyHashMap{m: make(map[int]int, 0)}
}

// Put puts value to the hashmap. Value will always be non-negative.
func (mhm *MyHashMap) Put(key int, value int) {
	mhm.m[key] = value
}

// Get returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (mhm *MyHashMap) Get(key int) int {
	if _, ok := mhm.m[key]; !ok {
		return -1
	}
	return mhm.m[key]
}

// Remove removes the mapping of the specified value key if this map contains a mapping for the key
func (mhm *MyHashMap) Remove(key int) {
	delete(mhm.m, key)
}
