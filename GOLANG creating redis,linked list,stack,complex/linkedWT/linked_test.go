package linkedWT

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	s := singleList{}
	expected := &s
	actual := initList()
	assert.Equal(t, expected, actual)
}

func TestInsertFrontFirst(t *testing.T) {
	s := singleList{}
	s.AddFront("first")
	expected := singleList{
		len:  1,
		head: &element{},
	}
	actual := s
	assert.Equal(t, expected, actual)

}

func TestInsertLenFirst(t *testing.T) {
	s := singleList{}
	s.AddFront("first")
	expected := 1
	actual := s.len
	assert.Equal(t, expected, actual)

}
func TestInsertHeadFirst(t *testing.T) {
	s := singleList{}
	s.AddFront("first")
	newele := element{
		value: "first",
	}
	expected := &newele
	actual := s.head
	assert.Equal(t, expected, actual)

}

func TestInsertBack(t *testing.T) {
	newele := element{
		value: "second",
	}
	oldele := element{
		value: "first",
	}
	s := singleList{
		len:  1,
		head: &oldele,
	}
	s.AddBack("second")
	expected := &newele
	actual := oldele.next
	assert.Equal(t, expected, actual)

}

func TestRemoveBack(t *testing.T) {
	thirdEl := element{
		value: "third",
	}
	secondEl := element{
		value: "second",
		next:  &thirdEl,
	}
	firstEl := element{
		value: "first",
		next:  &secondEl,
	}
	s := singleList{
		len:  3,
		head: &firstEl,
	}
	actual, err := s.RemoveBack()
	expected := secondEl.next
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)

}
