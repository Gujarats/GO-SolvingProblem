package main

type Node struct {
	Next  *Node
	Value int
}

func mergeTwoLinkedList(head1 *Node, head2 *Node) *Node {
	var node3 Node
	var tempResult Node
	if head1.Value < head2.Value {
		node3.Value = head1.Value
		tempResult = Node{Value: head2.Value}
		node3.Next = &tempResult
	} else {
		node3.Value = head2.Value
		tempResult = Node{Value: head1.Value}
		node3.Next = &tempResult
	}

	if head1.Next == nil && head2.Next == nil {
		return &node3
	}

	tempResult.Next = mergeTwoLinkedList(head1.Next, head2.Next)

	return &node3
}
