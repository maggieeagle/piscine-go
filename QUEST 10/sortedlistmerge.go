package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	node := n2
	if n1 == nil {
		return n2
	}
	for node != nil {
		SortListInsert(n1, node.Data)
		node = node.Next
	}
	return n1
}
