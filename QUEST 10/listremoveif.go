package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	old_node := l.Head
	node := l.Head
	for node != nil {
		if CompStr(data_ref, node.Data) {
			old_node.Next = node.Next
			node.Data = nil

			node = node.Next
		} else {
			old_node = node
			node = node.Next
		}
	}
	if l.Head.Data == nil {
		l.Head = l.Head.Next
	}
}
