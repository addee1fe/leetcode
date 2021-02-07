package solution

// Node ...
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	p := root
	for p != nil {
		if p.Child != nil {
			right := p.Next
			p.Next = flatten(p.Child)
			p.Next.Prev = p
			p.Child = nil

			for p.Next != nil {
				p = p.Next
			}

			if right != nil {
				p.Next = right
				p.Next.Prev = p
			}
		}
		p = p.Next
	}
	return root
}
