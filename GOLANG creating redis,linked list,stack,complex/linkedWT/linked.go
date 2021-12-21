package linkedWT

import "fmt"

type element struct {
	value string
	next  *element
}
type singleList struct {
	len  int
	head *element
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(value string) {
	ele := element{
		value: value,
	}
	if s.head == nil {
		s.head = &ele
	} else {
		ele.next = s.head
		s.head = &ele
	}
	s.len++
}

func (s *singleList) AddBack(value string) {
	ele := element{
		value: value,
	}
	if s.head == nil {
		s.head = &ele
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = &ele
	}
	s.len++

}

func (s *singleList) RemoveBack() (*element,error) {
	if s.head == nil {
		return nil,fmt.Errorf("list is empty")
	}
	prev:=&element{}
	current:=s.head
	for current.next!=nil{
		prev=current
		current=current.next
	}
	if prev!=nil{
		prev.next=nil
	}else {
		s.head=nil
	}
	s.len--
	return current.next, nil

}
