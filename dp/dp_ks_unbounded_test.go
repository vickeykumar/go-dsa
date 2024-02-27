package dp

import (
	"fmt"
	"testing"
)

func TestKnapsackUnbounded(t *testing.T) {
	// Test case 1: Basic test case with small values
	V1 := []int{60, 100, 120}
	W1 := []int{10, 20, 30}
	capacity1 := 50
	expectedResult1 := 300
	if result1 := Knapsack_unbounded(V1, W1, capacity1); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: V=%v, W=%v, capacity=%d, expected %d but got %d", V1, W1, capacity1, expectedResult1, result1)
	} else {
		fmt.Printf("Test case 1 passed: V=%v, W=%v, capacity=%d, result=%d\n", V1, W1, capacity1, result1)
	}

	// Test case 2: Large capacity with large values
	V2 := []int{100, 120, 300}
	W2 := []int{10, 20, 30}
	capacity2 := 100
	expectedResult2 := 1000
	if result2 := Knapsack_unbounded(V2, W2, capacity2); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: V=%v, W=%v, capacity=%d, expected %d but got %d", V2, W2, capacity2, expectedResult2, result2)
	} else {
		fmt.Printf("Test case 2 passed: V=%v, W=%v, capacity=%d, result=%d\n", V2, W2, capacity2, result2)
	}

	// Test case 3: Empty arrays
	V3 := []int{}
	W3 := []int{}
	capacity3 := 10
	expectedResult3 := 0
	if result3 := Knapsack_unbounded(V3, W3, capacity3); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: V=%v, W=%v, capacity=%d, expected %d but got %d", V3, W3, capacity3, expectedResult3, result3)
	} else {
		fmt.Printf("Test case 3 passed: V=%v, W=%v, capacity=%d, result=%d\n", V3, W3, capacity3, result3)
	}
}


func TestRodcut(t *testing.T) {
	// Test case 1: Example given in the problem description
	price1 := []int{1, 5, 8, 9, 10, 17, 17, 20}
	length1 := 8
	expectedResult1 := 22
	if result1 := Rodcut(price1, length1); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: price=%v, length=%d, expected %d but got %d", price1, length1, expectedResult1, result1)
	} else {
		t.Logf("Test case 1 passed: price=%v, length=%d, result=%d", price1, length1, result1)
	}

	// Test case 2: Another example given in the problem description
	price2 := []int{3, 5, 8, 9, 10, 17, 17, 20}
	length2 := 8
	expectedResult2 := 24
	if result2 := Rodcut(price2, length2); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: price=%v, length=%d, expected %d but got %d", price2, length2, expectedResult2, result2)
	} else {
		t.Logf("Test case 2 passed: price=%v, length=%d, result=%d", price2, length2, result2)
	}

	// Test case 3: Edge case with zero length
	price3 := []int{1, 2, 3}
	length3 := 0
	expectedResult3 := 0
	if result3 := Rodcut(price3, length3); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: price=%v, length=%d, expected %d but got %d", price3, length3, expectedResult3, result3)
	} else {
		t.Logf("Test case 3 passed: price=%v, length=%d, result=%d", price3, length3, result3)
	}

	// Add more test cases as needed...
}