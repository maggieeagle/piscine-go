package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	listPushBackNode(l, data_ref)
	ListSort(l)
	return l
}

func listPushBackNode(l *NodeI, data int) {
	n := &NodeI{Data: data, Next: nil}

	node := l
	for node.Next != nil {
		node = node.Next
	}
	node.Next = n
}
