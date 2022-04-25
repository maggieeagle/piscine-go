package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	node := l

	if node == nil {
		return nil
	}
	for node != nil {
		next := node.Next
		for next != nil {
			if node.Data > next.Data {
				swap(node, next)
			}
			next = next.Next
		}
		node = node.Next
	}
	return l
}

func swap(a, b *NodeI) {
	tmp := a.Data
	a.Data = b.Data
	b.Data = tmp
}
