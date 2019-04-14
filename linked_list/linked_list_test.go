package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LinkedListTestSuite struct {
	suite.Suite
}

func TestLinkedListSuite(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}

func (s *LinkedListTestSuite) AddNodesToFirstOfTheList() {
	ll := &LinkedList{}

	n1 := Node{
		Value: 9,
	}

	ll.AddToFirst(&n1)
}
