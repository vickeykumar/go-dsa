package pq

import (
    "fmt"
    "testing"
)

// Define a struct to represent operations on the priority queue
type operation struct {
    op   string
    data interface{}
}


// Helper function to check if the array is sorted
func isSorted(arr []int) bool {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] > arr[i+1] {
            return false
        }
    }
    return true
}

// Helper function to check if two arrays are equal
func isEqual(arr1, arr2 []int) bool {
    if len(arr1) != len(arr2) {
        return false
    }
    for i := 0; i < len(arr1); i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}


func TestPriorityQueue(t *testing.T) {
    // Test cases
    testCases := []struct {
        name          string
        pqInitializer func() *PriorityQueue
        operations    []operation
        expected      []interface{}
    }{
        {
            name: "Test Case 1",
            pqInitializer: func() *PriorityQueue {
                return New_pq(func(i, j interface{}) bool {
                    return i.(int) < j.(int) // Min priority queue
                })
            },
            operations: []operation{
                {op: "Push", data: 3},
                {op: "Push", data: 1},
                {op: "Push", data: 2},
                {op: "Pop", data: 1},
                {op: "Push", data: 4},
            },
            expected: []interface{}{2, 3, 4},
        },
        {
            name: "Test Case 2",
            pqInitializer: func() *PriorityQueue {
                return New_pq(func(i, j interface{}) bool {
                    return i.(int) > j.(int) // Max priority queue
                })
            },
            operations: []operation{
                {op: "Push", data: 5},
                {op: "Push", data: 7},
                {op: "Push", data: 3},
                {op: "Push", data: 9},
                {op: "Pop", data: 9},
                {op: "Push", data: 6},
            },
            expected: []interface{}{7, 6, 5, 3},
        },
        // Add more test cases as needed
    }

    // Run test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            pq := tc.pqInitializer()
            for _, op := range tc.operations {
                switch op.op {
                case "Push":
                    pq.Push(op.data)
                case "Pop":
                    popped := pq.Pop()
                    if popped != op.data {
                        t.Errorf("Pop() returned %v, expected %v", popped, op.data)
                    }
                }
                fmt.Printf("%s %v -> ",op.op, op.data)
                pq.PQ_for_each(func(data interface{}) {
                    fmt.Printf("%v ", data)
                })
                fmt.Println("")
            }

        })
    }
}

func TestHeapSort(t *testing.T) {
    // Test cases
    testCases := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "Test Case 1",
            input:    []int{5, 3, 8, 1, 2, 9, 4, 7, 6},
            expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
        },
        // Add more test cases as needed
    }

    // Run test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            fmt.Println("input: ", tc.input)
            HeapSort(tc.input)
            fmt.Println("output: ", tc.input)
            if !isSorted(tc.input) || !isEqual(tc.input, tc.expected) {
                t.Errorf("Test case failed: %s", tc.name)
            }
        })
    }
}
