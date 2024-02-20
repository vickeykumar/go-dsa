package queue

import (
	"go-dsa/llist"
)

type queue llist.Linkedlist

func NewQue() *queue {
	return (*queue)(llist.Linkedlist_new())
}

func (q *queue)Enque(data interface{}) {
	(*llist.Linkedlist)(q).Linkedlist_append(data)
}

func (q *queue)Deque() (data interface{}) {
	data=(*llist.Linkedlist)(q).Linkedlist_deletefront()
	return data
}

func (q *queue)Size() int {
	return (*llist.Linkedlist)(q).Size()
}

func (q *queue)IsEmpty() bool {
	return (*llist.Linkedlist)(q).Size()==0
}

func (q *queue)Display(foreach func(data interface {})) {
	(*llist.Linkedlist)(q).Display(foreach)
}