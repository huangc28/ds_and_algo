package main

import "log"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	firstNode := Node{
		Value: 3,
	}

	middleNode := Node{
		Value: 5,
	}

	firstNode.Next = &middleNode

	lastNode := Node{
		Value: 7,
	}

	middleNode.Next = &lastNode

	PrintList(&firstNode)
}

func PrintList(node *Node) {
	for node != nil {
		log.Printf("node value: %v", node.Value)
		node = node.Next
	}
}
