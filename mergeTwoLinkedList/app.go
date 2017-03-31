package main

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

func main() {
	// createing linkedList 1
	node4 := Node{Value: 5}
	node3 := Node{Next: &node4, Value: 3}
	node2 := Node{Next: &node3, Value: 2}
	node1 := Node{Next: &node2, Value: 1}

	// createing linkedlist 2
	n4 := Node{Value: 8}
	n3 := Node{Next: &n4, Value: 8}
	n2 := Node{Next: &n3, Value: 7}
	n1 := Node{Next: &n2, Value: 6}

	result := mergeTwoLinkedList2(&node1, &n1)
	printNode(result)

}

func printNode(head *Node) {
	fmt.Print(head.Value, " ")

	if head.Next == nil {
		return
	}

	printNode(head.Next)

	return
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

func mergeTwoLinkedList2(head1 *Node, head2 *Node) *Node {
	var result Node

	if head1.Next == nil {
		return head2
	}

	if head2.Next == nil {
		return head1
	}

	if head1.Value < head2.Value {
		result.Value = head1.Value
		result.Next = mergeTwoLinkedList2(head1.Next, head2)
	} else {
		result.Value = head2.Value
		result.Next = mergeTwoLinkedList2(head1, head2.Next)
	}

	return &result
}
