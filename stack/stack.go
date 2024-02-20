package stack

import (
	"fmt"
)

type node struct {
	data interface{}
	prev *node
}

type stack struct {
	top *node
	length int
}

func newnode (dataIn interface{}) *node {
	return &node{
		data:dataIn,
		prev:nil,
	}
}

func NewStack() *stack {
	return &stack{
		top: nil,
		length: 0,
	}
}

func (st *stack) Push(dataIn interface{}) {
	node := newnode(dataIn)
	node.prev = st.top
	st.top = node
	st.length++
}

func (st *stack) IsEmpty() bool {
	return (st.length==0)
}

func (st *stack) Pop() interface{} {
	if st.IsEmpty() {
		return nil
	}
	data := st.top.data 
	st.top=st.top.prev
	st.length--
	return data
}

func (st *stack) Peek() interface{} {
	if st.IsEmpty() {
		return nil
	}
	return st.top.data 
}

func (st *stack) Display(foreach func(data interface {})) {
    fmt.Println("stack: ")
    top := st.top
    for i:=0; i<st.length; i++ {
        foreach(top.data)
        top=top.prev
    }
}
