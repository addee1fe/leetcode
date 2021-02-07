package solution

import "testing"

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	if got := ms.GetMin(); got != -3 {
		t.Errorf("GetMin() returned: %v, should return: %v", got, -3)
	}
	ms.Pop()
	if got := ms.Top(); got != 0 {
		t.Errorf("Top() returned: %v, should return: %v", got, 0)
	}
	if got := ms.GetMin(); got != -2 {
		t.Errorf("GetMin() returned: %v, should return: %v", got, -2)
	}
}
