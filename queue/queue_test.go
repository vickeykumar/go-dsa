package queue


import (
    "fmt"
    "testing"
)

func TestStack(t *testing.T) {
    q := NewQue()
    q.Enque(21)
    q.Enque(23)
    q.Enque(25)
    q.Enque(27)
    q.Display(func(data interface{}) {
         fmt.Printf("%d<- ",data.(int))
    })

    data:=q.Deque()
    t.Log("data deleted: ",data)
    t.Log("after delete:")
    q.Display(func(data interface{}) {
        fmt.Printf("%d<- ",data.(int))
    })

}

