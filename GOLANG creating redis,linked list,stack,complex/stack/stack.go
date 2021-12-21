package stack

import (
	"container/list"
	"errors"
)

type Stack struct {
	stack *list.List
}

var ErrEmptyPush = errors.New("empty push is not allowed")

func (c *Stack) CreateStack() *list.List {
	firstStack := &Stack{
		stack: list.New(),
	}
	return firstStack.stack
}

/*func (c *Stack) Length() int {
	l := c.stack.Len()
	return l
}*/
func (c *Stack) Push(value string) (bool,error) {
	c.stack.PushFront(value)
	if value == ""{
		return true,ErrEmptyPush
	}
	return true,nil
}
