package list

type Node[T any] struct {
	Prev *Node[T]
	Next *Node[T]
	Val  T
}

// Insert adds a node after this one with a value of val.
func (n *Node[T]) Insert(val T) {
	new := &Node[T]{
		Next: n.Next,
		Prev: n,
		Val:  val,
	}
	if n.Next != nil {
		n.Next.Prev = new
	}
	n.Next = new
}

// Delete removes the current node, attaching the ones before and after it together.
func (n *Node[T]) Delete() T {
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	return n.Val
}
