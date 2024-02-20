package llist


import (
    "fmt"
)

type llnode struct {
    data interface{}
    next *llnode
} 

func new_llnode(dataIn interface{}) (*llnode) {
    return &llnode{
        data: dataIn,
        next: nil,
    }
}

type Linkedlist struct {
    head *llnode
    length int
}

func Linkedlist_new() (*Linkedlist) {
    return &Linkedlist{
        head: nil,
        length: 0,
    }
}

func (ll *Linkedlist)Linkedlist_prepend(data interface{}) {
    node := new_llnode(data)
    node.next = ll.head
    ll.head=node
    ll.length++
}

func (ll *Linkedlist)Linkedlist_append(data interface{}) {
    node := new_llnode(data)
    if ll.head==nil {
        // first node
        ll.head=node
    } else {
        // use a temp variable as head will persist in ll context
        last := ll.head
        for last.next != nil {
           last = last.next
        }
        last.next = node
    }
    ll.length++
}


func (ll *Linkedlist)Linkedlist_deleteend() interface{} {
    last := ll.head
    for last.next!=nil && last.next.next != nil {
        last=last.next
    }
    data := last.next.data
    last.next = nil
    ll.length--
    return data
}

func (ll *Linkedlist)Linkedlist_deletefront() interface{} {
    if ll.head == nil {
        return nil
    }
    data := ll.head.data
    ll.head=ll.head.next
    ll.length--
    return data
}

func (ll *Linkedlist)Size() int {
    return ll.length
}


func (ll *Linkedlist) Display(foreach func(data interface {})) {
    fmt.Printf("Linkedlist: ")
    last := ll.head 
    for i:=0; i<ll.Size(); i++ {
        foreach(last.data)
        last=last.next
    }
}

