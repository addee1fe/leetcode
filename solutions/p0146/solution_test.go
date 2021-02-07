package solution

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if r := cache.Get(1); r != 1 {
		t.Errorf("cache.Get(1) == %v, want = 1", r)
	}
	cache.Put(3, 3) // should evict 2, 2
	if r := cache.Get(2); r != -1 {
		t.Errorf("cache.Get(2) == %v, want = -1", r)
	}
	cache.Put(4, 4) // should evict 1, 1
	if r := cache.Get(1); r != -1 {
		t.Errorf("cache.Get(1) == %v, want = -1", r)
	}
	if r := cache.Get(3); r != 3 {
		t.Errorf("cache.Get(3) == %v, want = 3", r)
	}
	if r := cache.Get(4); r != 4 {
		t.Errorf("cache.Get(4) == %v, want = 4", r)
	}
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	if r := cache.Get(2); r != 2 {
		t.Errorf("cache.Get(2) == %v, want = 2", r)
	}
	cache.Put(1, 1)
	cache.Put(4, 1)
	if r := cache.Get(2); r != -1 {
		t.Errorf("cache.Get(2) == %v, want = -1", r)
	}
}
