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

// Test function
func TestReverseLinkedList(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
    }{
        {[]int{}, []int{}},                   // Empty list
        {[]int{1}, []int{1}},                 // Single node list
        {[]int{1, 2}, []int{2, 1}},           // Two elements
        {[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}}, // Multiple elements
    }

    for _, test := range tests {
        fmt.Println("Running test: ", test.input)
        ll := CreateLinkedList(test.input)

        fmt.Println("before: ")
        // Display before reversal
        ll.Display(func(data interface{}) {
            fmt.Printf("%d -> ", data.(int))
        })

        // Reverse and test
        reversedLL := ReverseLinkedList(ll)

        fmt.Println("after reversal: ")
        // Display after reversal
        reversedLL.Display(func(data interface{}) {
            fmt.Printf("%d -> ", data.(int))
        })
        fmt.Printf("\n")
    }
}

// Helper function to create a linked list from a slice
func CreateLinkedList(values []int) *Linkedlist {
    if len(values) == 0 {
        return &Linkedlist{}
    }
    ll := Linkedlist_new()
    for _, v := range values {
        ll.Linkedlist_append(v)
    }
    return ll
}