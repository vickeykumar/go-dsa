package sort

import (
	_ "fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
 * selection
 * sorts the array between index low and high
 */
func selection_sort(arr []int) []int {
    for i:=0; i<len(arr); i++ {
        var min int=i
        for  j:=i+1;j<len(arr);j++  {
            if arr[j]<arr[min] {
            	min=j;	
            } 
        }
        swap(arr, i, min);
    }
    return arr
}

/*
 * bubble sorts the array between index low and high
 * O(n2)
 */
func bubble_sort(arr []int, low, high int) []int {
    for i:=low; i<=high-1; i++ {
        for j:=low;j<=high-i-1;j++ {
            if arr[j]>arr[j+1] {
            	swap(arr, j, j+1);
            } 
        }
    }
    return arr
}

/*
 * insertion sorts the array between index low and high
 * O(n2) , best case O(n)
 */
func insertion_sort(arr []int, low, high int) []int {
    for i:=low; i<high; i++ {
        item := arr[i+1];
        j:=i;
        for j>=low && arr[j]>item {
             arr[j+1]=arr[j];
             j--;
        }
        arr[j+1]=item;
    }
    return arr
}

/* returns index of the element after partition,
 * such that all the elements left are less then  and all the elements right
 * are greater then the partitioned element (lomuto partition)
 * Hoare Partition: 
  - init two pointers l and h
  - pivot is smallest index
  - keep increasing l index while arr[l]=<pivot
  - keep decreasing h while arr[h]>pivot
  - if l>=h return l
  - if i<j swap(arr, l,h)
 * lomuto partition:
  - pivot will always be smallest index/ largest index
  - loop i=ipos... to high
  - alway keep ipos at last larger element.
  - as soon as i find the right fit for swaping, swap at ipos

 * Applications:
  - segregate 0 and 1a
  - segregate even and odd
  - -ve and +ve numbers
  - three way partitioning
 */
func partition(arr []int, low, high int) int {
    ipos := low;
    pivot := arr[ipos];
    for i:=ipos+1; i<=high; i++ {
        if arr[i]<=pivot { // found a right fit to place at ipos
            ipos++;   // this way ipos will always point to last largest elem visited by i
            swap(arr, i, ipos);
        }
    }
    swap(arr, ipos,low);
    return ipos;
}

/*
 * quick sorts the array between index low and high
 */
func quick_sort(arr []int, low, high int) []int {
    if low<high {
        p := partition(arr, low, high);
        arr=quick_sort(arr, low, p-1); 
        // remember still arr will have all the indices as 
        // q sort is sorting only between desired indices
        arr=quick_sort(arr, p+1, high);
    }
    return arr
}

func randomized_quick_sort(arr []int, low, high int) []int {
    if low<high {
        i:= rand.Intn(high-low+1)
        swap(arr, i, low);
        p := partition(arr, low, high);
        arr=quick_sort(arr, low, p-1);
        arr=quick_sort(arr, p+1, high);
    }
    return arr
}

/* 
 * vvi
 * merge 2 sorted array using in place sorting 
 * logic - for each ith element on right subarray if is smaller than the low-th 
 * element, use bubble sort(bubbling)/ shift from j=i-1 till low-th index till ith item finds its right place
 * in left subarray, o(n)
 * iteraten in right subarray and try to find ith items best position in left subarray j=i-1to low
 */
func merge(arr []int, low, mid, high int) []int {
    for i:=mid; i<=high; i++ {
        item := arr[i];
        j:=i-1;
        for j>=low && arr[j]>item { 
            arr[j+1]=arr[j];
            j--;
        }
        arr[j+1]=item;
        low = j+1;  // with each iteration this low par will keep increasing, 
        // as minimum of second array will always be greater than the next j
    }
    return arr

}

/*
 * sorts the array between index low and high
 */
func merge_sort(arr []int, low, high int) []int {
    if low<high {
        var mid int = (low+high)/2
        arr=merge_sort(arr, low, mid);
        arr=merge_sort(arr, mid+1, high);
        arr=merge(arr, low, mid+1, high);
    }
    return arr
}