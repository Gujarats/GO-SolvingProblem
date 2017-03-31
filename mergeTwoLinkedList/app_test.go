package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLinkedList(t *testing.T) {
	type testObject struct {
		NodeData1 Node
		NodeData2 Node
		Expected  Node
	}

	// createing linkedList 1
	node4 := Node{Value: 9}
	node3 := Node{Next: &node4, Value: 7}
	node2 := Node{Next: &node3, Value: 3}
	node1 := Node{Next: &node2, Value: 0}

	// createing linkedlist 2
	n4 := Node{Value: 8}
	n3 := Node{Next: &n4, Value: 5}
	n2 := Node{Next: &n3, Value: 2}
	n1 := Node{Next: &n2, Value: 1}

	// create node expedted result
	result8 := Node{Value: 9}
	result7 := Node{Next: &result8, Value: 8}
	result6 := Node{Next: &result7, Value: 7}
	result5 := Node{Next: &result6, Value: 5}
	result4 := Node{Next: &result5, Value: 3}
	result3 := Node{Next: &result4, Value: 2}
	result2 := Node{Next: &result3, Value: 1}
	result1 := Node{Next: &result2, Value: 0}

	testObjects := []testObject{
		{NodeData1: node1, NodeData2: node2, Expected: result1},
	}

	for _, test := range testObjects {
		actual := mergeTwoLinkedList(&node1, &n1)

		fmt.Println("actual")
		printNode(actual)
		fmt.Println()

		fmt.Println("expected")
		printNode(&test.Expected)
		fmt.Println()

		if !isNodeSimilar(actual, &test.Expected) {
			t.Errorf("Error Actual = %+v, Expected = %+v", actual, test.Expected)
		}
	}
}

func isNodeSimilar(head1 *Node, head2 *Node) bool {
	var result bool

	if head1.Value != head2.Value {
		result = false
	} else {
		result = true
	}

	if head1.Next == nil && head2.Next == nil {
		return result
	}

	//checking next value nill or not
	if (head1.Next == nil && head2.Next != nil) || (head1.Next != nil && head2.Next == nil) {
		return false
	}

	result = isNodeSimilar(head1.Next, head2.Next)

	return result

}
