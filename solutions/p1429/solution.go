package solution

import (
	"container/list"
)

// FirstUnique ...
type FirstUnique struct {
	uniq  *list.List
	items map[int]*list.Element
}

// Constructor ...
func Constructor(nums []int) FirstUnique {
	uniq := list.New()
	items := make(map[int]*list.Element)
	for _, v := range nums {
		if el, ok := items[v]; ok {
			uniq.Remove(el)
		} else {
			items[v] = uniq.PushBack(v)
		}
	}
	return FirstUnique{uniq, items}
}

// ShowFirstUnique ...
func (fu *FirstUnique) ShowFirstUnique() int {
	if fu.uniq.Len() == 0 {
		return -1
	}
	return fu.uniq.Front().Value.(int)
}

// Add ...
func (fu *FirstUnique) Add(value int) {
	if el, ok := fu.items[value]; ok {
		if el != nil {
			fu.uniq.Remove(el)
		}
	} else {
		fu.items[value] = fu.uniq.PushBack(value)
	}
}
