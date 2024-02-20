package llist


import (
    "fmt"
    "testing"
)

func TestLinkedList(t *testing.T) {
    ll := Linkedlist_new()
    ll.Linkedlist_append(21)
    ll.Linkedlist_append(23)
    ll.Linkedlist_append(21)
    ll.Linkedlist_append(23)
    result := ""
    ll.Display(func(data interface{}) {
        result += fmt.Sprintf("%d -> ",data.(int))
    })
    t.Log(result)

    data:=ll.Linkedlist_deleteend()
    t.Log("data deleted: ",data)
    t.Log("after delete:")
    result = ""
    ll.Display(func(data interface{}) {
        result += fmt.Sprintf("%d -> ",data.(int))
    })
    t.Log(result)

}

