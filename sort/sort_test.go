package sort

import (
    "fmt"
    "testing"
)

func TestStack(t *testing.T) {
    arr := []int{34,1,4,1,76,2,5,6}
    fmt.Println("Arr: ",arr)
   	duparr := make([]int,len(arr))

   	copy(duparr,arr)
   	fmt.Println("dupArr: ",duparr)

   	copy(duparr,arr)
   	fmt.Println("before selection_sort: ",duparr)
   	selection_sort(duparr)
   	fmt.Println("After selection_sort: ",duparr)

   	copy(duparr,arr)
   	fmt.Println("before bubble_sort: ",duparr)
   	bubble_sort(duparr, 0, len(duparr)-1)
   	fmt.Println("After bubble_sort: ",duparr)

   	copy(duparr,arr)
   	fmt.Println("before insertion_sort: ",duparr)
	insertion_sort(duparr, 0, len(duparr)-1)
   	fmt.Println("After insertion_sort: ",duparr)

   	copy(duparr,arr)
   	fmt.Println("before quick_sort: ",duparr)
	quick_sort(duparr, 0, len(duparr)-1)
   	fmt.Println("After quick_sort: ",duparr)

   	copy(duparr,arr)
   	fmt.Println("before randomized_quick_sort: ",duparr)
	randomized_quick_sort(duparr, 0, len(duparr)-1)
   	fmt.Println("After randomized_quick_sort: ",duparr)

   	copy(duparr,arr)
   	fmt.Println("before merge_sort: ",duparr)
	merge_sort(duparr, 0, len(duparr)-1)
   	fmt.Println("After merge_sort: ",duparr)
}

