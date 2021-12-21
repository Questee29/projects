package stack

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfCreated(t *testing.T) {
	s := Stack{
		stack: list.New(),
	}
	expected := s.stack
	actual := s.CreateStack()
	assert.Equal(t, expected, actual)

}
func TestPushOne(t *testing.T) {
	s := Stack{
		stack: list.New(),
	}
	actual,err := s.Push("artem")
	expected :=true
	assert.Equal(t, expected, actual)
	assert.Nil(t,err)
	

}
func TestPushEmpty(t *testing.T) {
	s := Stack{
		stack: list.New(),
	}
	actual,err := s.Push("")
	expected :=true
	assert.Equal(t, expected, actual)
	assert.NotNil(t,err)
	

}
/*func TestNotEmpty(t *testing.T){
	s := Stack{
		stack: list.New(),
	}
	err := s.Push()
	expected := ErrEmptyPush
	assert.Equal(t, expected, err)

}*/
