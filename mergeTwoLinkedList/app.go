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

	result := mergeTwoLinkedList(&node1, &n1)
	printNode(result)
	fmt.Println()
	fmt.Println("test 2 ")

	node4 = Node{Value: 15}
	node3 = Node{Next: &node4, Value: 7}
	node2 = Node{Next: &node3, Value: 2}
	node1 = Node{Next: &node2, Value: 1}

	// createing linkedlist 2
	n4 = Node{Value: 10}
	n3 = Node{Next: &n4, Value: 9}
	n2 = Node{Next: &n3, Value: 5}
	n1 = Node{Next: &n2, Value: 3}

	result = mergeTwoLinkedList(&node1, &n1)
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
	var result Node

	if head1.Next == nil {
		return head2
	}

	if head2.Next == nil {
		return head1
	}

	if head1.Value < head2.Value {
		result.Value = head1.Value
		result.Next = mergeTwoLinkedList(head1.Next, head2)
	} else {
		result.Value = head2.Value
		result.Next = mergeTwoLinkedList(head1, head2.Next)
	}

	return &result
}
