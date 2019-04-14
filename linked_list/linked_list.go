package linkedlist

import "log"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

func (ll *LinkedList) AddToFirst(n *Node) {
	currentHead := ll.Head
	ll.Head = n

	ll.Head.Next = currentHead
	ll.Count++

	if ll.Count == 0 {
		ll.Tail = n
	}
}

func (ll *LinkedList) AddToLast(n *Node) {
	if ll.Count == 0 {
		ll.Head = n
	} else {
		ll.Tail.Next = n
	}

	ll.Tail = n
	ll.Count++
}

func (ll *LinkedList) RemoveLast(n *Node) {
	var current *Node

	if ll.Count == 1 {
		ll.Head = nil
		ll.Tail = nil
	} else {
		for ll.Head.Next != nil {
			current = ll.Head.Next
		}

		current.Next = nil
		ll.Tail = current
	}

	ll.Count--
}

func (ll *LinkedList) RemoveFirst(n *Node) {
	if ll.Count != 0 {
		ll.Head = ll.Head.Next
		ll.Count--

		if ll.Count == 0 {
			ll.Tail = nil
		}
	}
}

func (ll *LinkedList) Enumerate(n *Node) {
	var current *Node = ll.Head

	for current != nil {
		log.Printf("current node value %v", current.Value)
		current = ll.Head.Next
	}
}
