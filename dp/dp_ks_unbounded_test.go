package dp

import (
	"fmt"
	"testing"
	"math"
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

func TestNoways_CoinChange(t *testing.T) {
	testCases := []struct {
		coins    []int
		sum      int
		expected int
	}{
		// Test cases where the sum can be achieved using the given coins
		{[]int{1, 2, 5}, 5, 4},   // Expected: 4 (1+1+1+1+1, 1+1+1+2, 1+2+2, 5)
		{[]int{2}, 3, 0},          // Expected: 0 (No way to make 3 with only 2)
		{[]int{1, 2, 5}, 11, 11},  // Expected: 11 (All 1s)
		{[]int{2, 3, 5}, 8, 3},   // Expected: 10 (1 way to make 8 with 2,2,2,2, 1 way  2,3,3 and 1 ways with 3,5)
		{[]int{1, 3, 4}, 7, 5},    // Expected: 5 (5 ways: 1+1+1+1+1+1+1, 1+1+1+1+3, 1+1+1+4, 1+3+3, , 3+4, 7)

		// Test cases where the sum cannot be achieved using the given coins
		{[]int{2}, 1, 0},          // Expected: 0 (No way to make 1 with only 2)
		{[]int{1, 3, 4}, 0, 1},    // Expected: 1 (Empty set)
		{[]int{2, 3, 5}, 1, 0},    // Expected: 0 (No way to make 1 with given coins)
	}

	for _, tc := range testCases {
		result := Noways_CoinChange(tc.coins, tc.sum)
		if result != tc.expected {
			t.Errorf("Test case failed: coins=%v, sum=%d, expected=%d, got=%d", tc.coins, tc.sum, tc.expected, result)
		} else {
			t.Logf("Test case passed: coins=%v, sum=%d, expected=%d, got=%d", tc.coins, tc.sum, tc.expected, result)
		}
	}

	fmt.Println("Noways_CoinChange_opt: ")
	for _, tc := range testCases {
		result := Noways_CoinChange_opt(tc.coins, tc.sum)
		if result != tc.expected {
			t.Errorf("Test case failed: coins=%v, sum=%d, expected=%d, got=%d", tc.coins, tc.sum, tc.expected, result)
		} else {
			t.Logf("Test case passed: coins=%v, sum=%d, expected=%d, got=%d", tc.coins, tc.sum, tc.expected, result)
		}
	}
}


func TestMin_CoinChange(t *testing.T) {
	// Test cases where the minimum number of coins required is unique
	testCases := []struct {
		coins      []int
		sum        int
		expected   int
		testCaseID string
	}{
		// Base cases
		{[]int{1}, 0, 0, "Base case: sum is 0"},
		{[]int{1}, 1, 1, "Base case: single coin denomination equal to sum"},
		{[]int{2}, 3, math.MaxInt32, "Base case: single coin denomination smaller than sum"},
		{[]int{1, 2, 5}, 11, 3, "Typical case: multiple coin denominations"},
		{[]int{1, 2, 5}, 13, 4, "Typical case: multiple coin denominations"},
		{[]int{1, 3, 4}, 7, 2, "Typical case: multiple coin denominations"},
		// Additional test cases
		{[]int{2, 5, 10}, 15, 2, "Additional case: multiple coin denominations with larger values"},
		{[]int{1, 3, 4}, 6, 2, "Additional case: multiple coin denominations with smaller sum"},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseID, func(t *testing.T) {
			result := Min_CoinChange(tc.coins, tc.sum)
			if result != tc.expected {
				t.Errorf("Test case %s failed: expected %d, got %d", tc.testCaseID, tc.expected, result)
			} else {
				t.Logf("Test case %s passed: expected %d, got %d", tc.testCaseID, tc.expected, result)
			}
		})
	}
}