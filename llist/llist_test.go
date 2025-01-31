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
        ll := CreateList(test.input)

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


// Test LinkedListIntersection function
func TestLinkedListIntersection(t *testing.T) {
    // Helper function to create a linked list with given nodes
    createList := func(values []int) *Linkedlist {
        ll := Linkedlist_new()
        for _, v := range values {
            ll.Linkedlist_append(v)
        }
        return ll
    }

    // Case 1: No intersection
    listA := createList([]int{1, 2, 3, 4})
    listB := createList([]int{5, 6, 7, 8})
    if LinkedListIntersection(listA, listB) != nil {
        t.Errorf("Expected no intersection, but found one")
    }

    // Case 2: Intersection at the head
    listC := createList([]int{10, 20, 30})
    listD := &Linkedlist{head: listC.head} // Pointing to the same head node
    if LinkedListIntersection(listC, listD) != listC.head {
        t.Errorf("Expected intersection at head, but got different node")
    }

    // Case 3: Intersection at a middle node
    listE := createList([]int{1, 2})
    listF := createList([]int{9, 8})
    commonNode := new_llnode(3)
    commonNode.next = new_llnode(4)
    listE.head.next.next = commonNode // 1 -> 2 -> 3 -> 4
    listF.head.next.next = commonNode // 9 -> 8 -> 3 -> 4
    if LinkedListIntersection(listE, listF) != commonNode {
        t.Errorf("Expected intersection at node 3, but got a different node")
    }

    // Case 4: Intersection at the last node
    listG := createList([]int{100, 200})
    listH := createList([]int{300})
    lastNode := new_llnode(400)
    listG.head.next.next = lastNode // 100 -> 200 -> 400
    listH.head.next = lastNode      // 300 -> 400
    if LinkedListIntersection(listG, listH) != lastNode {
        t.Errorf("Expected intersection at last node 400, but got different node")
    }

    fmt.Println("✅ All intersection test cases passed!")
}

func TestLRUCache(t *testing.T) {
    cache := NewLRUCache(3)

    // Insert elements
    cache.put([]byte("a"), []byte("1"))
    cache.put([]byte("b"), []byte("2"))
    cache.put([]byte("c"), []byte("3"))

    // Test retrieving existing elements
    if string(cache.get([]byte("a"))) != "1" {
        t.Errorf("Expected 1, got %s", cache.get([]byte("a")))
    }

    // Insert another element to trigger eviction
    cache.put([]byte("d"), []byte("4"))

    // Ensure LRU element ('b') was evicted
    if cache.get([]byte("b")) != nil {
        t.Errorf("Expected 'b' to be evicted")
    }

    // Ensure 'c' and 'a' are still in cache
    if string(cache.get([]byte("c"))) != "3" {
        t.Errorf("Expected 3, got %s", cache.get([]byte("c")))
    }
    if string(cache.get([]byte("a"))) != "1" {
        t.Errorf("Expected 1, got %s", cache.get([]byte("a")))
    }

    // Ensure 'd' is most recently used
    if string(cache.get([]byte("d"))) != "4" {
        t.Errorf("Expected 4, got %s", cache.get([]byte("d")))
    }

    fmt.Println("✅ LRU Cache test passed!")
}