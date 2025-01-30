package pq

import (
 	"fmt"
	"reflect"
)

/*
using container/heap

// Define a struct to represent a node in the priority queue.
type Node struct {
    index int
    dist  int
}

// Define a priority queue data structure.
type PriorityQueue []*Node

// Implement the heap.Interface methods for PriorityQueue.

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // In this case, prioritize nodes with smaller distances.
    return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    // Push element x onto the heap.
    item := x.(*Node)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    // Pop and return the minimum element from the heap.
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

*/


type PriorityQueue struct {
	dataArr []interface{} // array of elements of anytype
	compare func(i,j interface{}) bool // comparater function to decide if min pq or max pq i<j for minpq
	// to maintain the heap property
}

func parent(i int) int {
	return (i-1)/2
}

func left (i int) int {
	return 2*i+1
}

func right (i int) int {
	return 2*i+2
}

func New_pq(comp func(i, j interface{}) bool) *PriorityQueue {
	arr := make([]interface{},0)
	return &PriorityQueue {
		dataArr: arr,
		compare: comp,
	}
}

func New_pq_arr(arr interface{}, comp func(i, j interface{}) bool) *PriorityQueue {
	// Check if arr is a slice
    arrValue := reflect.ValueOf(arr)
    if arrValue.Kind() != reflect.Slice {
        return nil
    }

    // Convert each element of arr to interface{} and append to varr
    var varr []interface{}
    for i := 0; i < arrValue.Len(); i++ {
        varr = append(varr, arrValue.Index(i).Interface())
    }

	pq := &PriorityQueue {
		dataArr: varr,
		compare: comp,
	}
	pq.Buildheap()
	return pq
}

func (pq *PriorityQueue)IsEmpty() bool {
	return len(pq.dataArr)==0
}

func (pq *PriorityQueue)PQ_for_each(fun func(interface{})) {
	for _, data := range pq.dataArr {
		fun(data)
	}
}

// maintains the heap property by moving the newly added element up in the heap
func (pq *PriorityQueue)heapifyUp(index int) {
	for index > 0 {
		parentindex := parent(index)
		if pq.compare(pq.dataArr[parentindex], pq.dataArr[index]) {
			// if compare is satisfied then we can break, like data[parent]< data[index] for a minpq
			// we assume rest of the up in the heap is already maintaining this property
			break
		}
		//else keep swaping and going up
		pq.dataArr[parentindex], pq.dataArr[index] = pq.dataArr[index], pq.dataArr[parentindex]
		index = parentindex
	}
}

// maintains the heap property by moving the root node down the heap,
// we solve this assuming this is a min heap
func (pq *PriorityQueue)heapifyDown(index int) {

	for index < len(pq.dataArr) {
		leftindex := left(index)
		rightindex := right(index)
		smallest := index
		if leftindex<len(pq.dataArr) && pq.compare(pq.dataArr[leftindex], pq.dataArr[smallest]) {
			smallest = leftindex
		}
		if rightindex<len(pq.dataArr) && pq.compare(pq.dataArr[rightindex], pq.dataArr[smallest]) {
			smallest = rightindex
		}

		// still smallest is this current index, no scope of going down further, BREAK
		if smallest == index {
			break
		}
		// else swap and go down in smallest direction
		pq.dataArr[index], pq.dataArr[smallest] = pq.dataArr[smallest], pq.dataArr[index]
		index = smallest
	}
}

// push inserts a new element in the heap/pq  
func (pq *PriorityQueue)Push(data interface{}) {
	pq.dataArr = append(pq.dataArr, data) // insert at last index
	// heapify up from last index till data finds its right place in the heap
	pq.heapifyUp(len(pq.dataArr)-1)
}

// pops a min/max element in the heap/pq  / based on the heap type, similar to extractmin and extract_max
func (pq *PriorityQueue)Pop() interface{} {
	if pq.IsEmpty() {
		return nil
	}
	data := pq.dataArr[0]
	pq.dataArr[0] = pq.dataArr[len(pq.dataArr)-1] 
	// gets data from last index and run heapify down from root to maintain the property
	// reduce the size of array
	pq.dataArr = pq.dataArr[:len(pq.dataArr)-1]
	pq.heapifyDown(0)
	return data
}

// runs heapifydown to build heap from down (n/2-1)to up 0
func (pq *PriorityQueue)Buildheap() {
	for i:=len(pq.dataArr)/2-1;i>=0;i-- {
		pq.heapifyDown(i)
	}
}

func HeapSort(arr []int) {
    // Build a min-heap from the input array
    heap := New_pq_arr(arr, func(i,j interface{}) bool {
    	return i.(int) < j.(int)
    })

    last:=len(arr)-1
    // Extract the minimum element from the heap and place it at the end of the array
    for i := last; i >= 0; i-- {
        d := heap.Pop().(int)
        //fmt.Println("pop: ",i, d)
        arr[last-i] = d
    }
    fmt.Println(arr)
}