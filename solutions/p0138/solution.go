package solution

// Node ...
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for v := head; v != nil; v = v.Next.Next {
		t := &Node{
			Val:    v.Val,
			Next:   v.Next,
			Random: v.Random,
		}
		v.Next = t
	}
	for v := head; v != nil; v = v.Next.Next {
		if v.Next.Random != nil {
			v.Next.Random = v.Next.Random.Next
		}
	}
	dummy := head.Next
	for ; head.Next.Next != nil; head = head.Next {
		head.Next, head.Next.Next = head.Next.Next, head.Next.Next.Next
	}
	head.Next = nil
	return dummy
}
