package dp

import (
	"fmt"
	"testing"
)



func TestKnapsack_01(t *testing.T) {
	tests := []struct {
		V        []int
		W        []int
		Capacity int
		Expected int
	}{
		// Test case 1: Regular case
		{[]int{60, 100, 120}, []int{10, 20, 30}, 50, 220},
		// Test case 2: Empty arrays
		{[]int{}, []int{}, 50, 0},
		// Test case 3: Capacity is zero
		{[]int{60, 100, 120}, []int{10, 20, 30}, 0, 0},
		// Add more test cases as needed
	}

	for idx, test := range tests {
		fmt.Println ("Knapsack_01", test.V, test.W, test.Capacity)
		result := Knapsack_01(test.V, test.W, test.Capacity)
		if result != test.Expected {
			t.Errorf("Test case %d failed: Expected %d, got %d", idx+1, test.Expected, result)
		}
	}
}


func TestIsSubsetSumPresent(t *testing.T) {
	// Test case 1: Subset sum is present
	nums1 := []int{3, 34, 4, 12, 5, 2}
	sum1 := 9
	expectedResult1 := true
	if result1 := IsSubsetSumPresent(nums1, sum1); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: expected %v but got %v | Parameters: nums=%v, sum=%v", expectedResult1, result1, nums1, sum1)
	} else {
		fmt.Printf("Function: TestIsSubsetSumPresent | Parameters: nums=%v, sum=%v Result: %t\n", nums1, sum1, result1)
	}

	// Test case 2: Subset sum is not present
	nums2 := []int{3, 34, 4, 12, 5, 2}
	sum2 := 30
	expectedResult2 := false
	if result2 := IsSubsetSumPresent(nums2, sum2); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: expected %v but got %v | Parameters: nums=%v, sum=%v", expectedResult2, result2, nums2, sum2)
	} else {
		fmt.Printf("Function: TestIsSubsetSumPresent | Parameters: nums=%v, sum=%v Result: %t\n", nums2, sum2, result2)
	}

	// Test case 3: Empty array
	nums3 := []int{}
	sum3 := 0
	expectedResult3 := true
	if result3 := IsSubsetSumPresent(nums3, sum3); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: expected %v but got %v | Parameters: nums=%v, sum=%v", expectedResult3, result3, nums3, sum3)
	} else {
		fmt.Printf("Function: TestIsSubsetSumPresent | Parameters: nums=%v, sum=%v Result: %t\n", nums3, sum3, result3)
	}

	// Test case 4: Single element array where sum is present
	nums4 := []int{5}
	sum4 := 5
	expectedResult4 := true
	if result4 := IsSubsetSumPresent(nums4, sum4); result4 != expectedResult4 {
		t.Errorf("Test case 4 failed: expected %v but got %v | Parameters: nums=%v, sum=%v", expectedResult4, result4, nums4, sum4)
	} else {
		fmt.Printf("Function: TestIsSubsetSumPresent | Parameters: nums=%v, sum=%v Result: %t\n", nums4, sum4, result4)
	}

	// Test case 5: Single element array where sum is not present
	nums5 := []int{5}
	sum5 := 10
	expectedResult5 := false
	if result5 := IsSubsetSumPresent(nums5, sum5); result5 != expectedResult5 {
		t.Errorf("Test case 5 failed: expected %v but got %v | Parameters: nums=%v, sum=%v", expectedResult5, result5, nums5, sum5)
	} else {
		fmt.Printf("Function: TestIsSubsetSumPresent | Parameters: nums=%v, sum=%v Result: %t\n", nums5, sum5, result5)
	}
}


func TestIsEqualSumPartPresent(t *testing.T) {
	// Test case 1: Even length array with equal sum partition present
	nums1 := []int{1, 5, 11, 5}
	expectedResult1 := true
	if result1 := IsEqualSumPartPresent(nums1); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: expected %v but got %v", expectedResult1, result1)
	} else {
		fmt.Printf("Test case 1 passed: nums=%v, result=%v\n", nums1, result1)
	}

	// Test case 2: Even length array with no equal sum partition present
	nums2 := []int{1, 2, 3, 4}
	expectedResult2 := true
	if result2 := IsEqualSumPartPresent(nums2); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: expected %v but got %v", expectedResult2, result2)
	} else {
		fmt.Printf("Test case 2 passed: nums=%v, result=%v\n", nums2, result2)
	}

	// Test case 3: Empty array
	nums3 := []int{}
	expectedResult3 := true // An empty array has an equal sum partition of 0
	if result3 := IsEqualSumPartPresent(nums3); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: expected %v but got %v", expectedResult3, result3)
	} else {
		fmt.Printf("Test case 3 passed: nums=%v, result=%v\n", nums3, result3)
	}

	// Test case 4: Single element array
	nums4 := []int{5}
	expectedResult4 := false // Single element array cannot have equal sum partition
	if result4 := IsEqualSumPartPresent(nums4); result4 != expectedResult4 {
		t.Errorf("Test case 4 failed: expected %v but got %v", expectedResult4, result4)
	} else {
		fmt.Printf("Test case 4 passed: nums=%v, result=%v\n", nums4, result4)
	}

	// Test case 5: Odd length array
	nums5 := []int{1, 2, 3}
	expectedResult5 := false // Odd length array cannot have equal sum partition
	if result5 := IsEqualSumPartPresent(nums5); result5 != expectedResult5 {
		t.Errorf("Test case 5 failed: expected %v but got %v", expectedResult5, result5)
	} else {
		fmt.Printf("Test case 5 passed: nums=%v, result=%v\n", nums5, result5)
	}
}

func TestCountSubsetSum(t *testing.T) {
	// Test case 1: Subset sum is present
	nums1 := []int{3, 34, 4, 12, 5, 2}
	sum1 := 9
	expectedResult1 := 2
	if result1 := CountSubsetSum(nums1, sum1); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: nums=%v, sum=%d, expected %d but got %d", nums1, sum1, expectedResult1, result1)
	} else {
		fmt.Printf("Test case 1 passed: nums=%v, sum=%d, result=%d\n", nums1, sum1, result1)
	}

	// Test case 2: Subset sum is not present
	nums2 := []int{3, 34, 4, 12, 5, 2}
	sum2 := 30
	expectedResult2 := 0
	if result2 := CountSubsetSum(nums2, sum2); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: nums=%v, sum=%d, expected %d but got %d", nums2, sum2, expectedResult2, result2)
	} else {
		fmt.Printf("Test case 2 passed: nums=%v, sum=%d, result=%d\n", nums2, sum2, result2)
	}

	// Test case 3: Empty array
	nums3 := []int{}
	sum3 := 0
	expectedResult3 := 1 // Empty array has a subset sum of 0
	if result3 := CountSubsetSum(nums3, sum3); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: nums=%v, sum=%d, expected %d but got %d", nums3, sum3, expectedResult3, result3)
	} else {
		fmt.Printf("Test case 3 passed: nums=%v, sum=%d, result=%d\n", nums3, sum3, result3)
	}

	// Test case 4: Single element array with subset sum present
	nums4 := []int{5}
	sum4 := 5
	expectedResult4 := 1 // Single element itself is the subset sum
	if result4 := CountSubsetSum(nums4, sum4); result4 != expectedResult4 {
		t.Errorf("Test case 4 failed: nums=%v, sum=%d, expected %d but got %d", nums4, sum4, expectedResult4, result4)
	} else {
		fmt.Printf("Test case 4 passed: nums=%v, sum=%d, result=%d\n", nums4, sum4, result4)
	}

	// Test case 5: Single element array with subset sum not present
	nums5 := []int{5}
	sum5 := 10
	expectedResult5 := 0 // Single element cannot have subset sum of 10
	if result5 := CountSubsetSum(nums5, sum5); result5 != expectedResult5 {
		t.Errorf("Test case 5 failed: nums=%v, sum=%d, expected %d but got %d", nums5, sum5, expectedResult5, result5)
	} else {
		fmt.Printf("Test case 5 passed: nums=%v, sum=%d, result=%d\n", nums5, sum5, result5)
	}
}

func TestMinSubsetSumDiff(t *testing.T) {
    // Test case 1: Subset sum difference is present
    nums1 := []int{3, 1, 4, 2, 2}
    expectedResult1 := 0 // The optimal subset sums are {1, 3, 4} and {2, 2}, with a difference of 0
    if result1 := MinSubsetSumDiff(nums1); result1 != expectedResult1 {
        t.Errorf("Test case 1 failed: nums=%v, expected %d but got %d", nums1, expectedResult1, result1)
    } else {
        t.Logf("Test case 1 passed: nums=%v, result=%d", nums1, result1)
    }

    // Test case 2: Subset sum difference is present
    nums2 := []int{5, 6, 7, 8, 9}
    expectedResult2 := 1 // The optimal subset sums are {5, 6, 8} and {7, 9}, with a difference of 1
    if result2 := MinSubsetSumDiff(nums2); result2 != expectedResult2 {
        t.Errorf("Test case 2 failed: nums=%v, expected %d but got %d", nums2, expectedResult2, result2)
    } else {
        t.Logf("Test case 2 passed: nums=%v, result=%d", nums2, result2)
    }

    // Test case 3: Subset sum difference is present
    nums3 := []int{10, 20, 30, 40, 50}
    expectedResult3 := 10 // The optimal subset sums are {10, 30, 40} and {20, 50}, with a difference of 0
    if result3 := MinSubsetSumDiff(nums3); result3 != expectedResult3 {
        t.Errorf("Test case 3 failed: nums=%v, expected %d but got %d", nums3, expectedResult3, result3)
    } else {
        t.Logf("Test case 3 passed: nums=%v, result=%d", nums3, result3)
    }

    // Test case 4: Subset sum difference is present
    nums4 := []int{1, 3, 5, 7, 9}
    expectedResult4 := 1 // The optimal subset sums are {1, 5, 9} and {3, 7}, with a difference of 1
    if result4 := MinSubsetSumDiff(nums4); result4 != expectedResult4 {
        t.Errorf("Test case 4 failed: nums=%v, expected %d but got %d", nums4, expectedResult4, result4)
    } else {
        t.Logf("Test case 4 passed: nums=%v, result=%d", nums4, result4)
    }

    // Test case 5: Subset sum difference is present
    nums5 := []int{2, 4, 6, 8, 10}
    expectedResult5 := 2 // The optimal subset sums are {2, 6, 8} and {4, 10}, with a difference of 2
    if result5 := MinSubsetSumDiff(nums5); result5 != expectedResult5 {
        t.Errorf("Test case 5 failed: nums=%v, expected %d but got %d", nums5, expectedResult5, result5)
    } else {
        t.Logf("Test case 5 passed: nums=%v, result=%d", nums5, result5)
    }

    // Test case 6: Subset sum difference is present
    nums6 := []int{5, 15, 25, 35, 45}
    expectedResult6 := 5 // The optimal subset sums are {5, 25, 35} and {15, 45}, with a difference of 0
    if result6 := MinSubsetSumDiff(nums6); result6 != expectedResult6 {
        t.Errorf("Test case 6 failed: nums=%v, expected %d but got %d", nums6, expectedResult6, result6)
    } else {
        t.Logf("Test case 6 passed: nums=%v, result=%d", nums6, result6)
    }

    // Test case 7: Empty array
    nums7 := []int{}
    expected7 := 0 // No subsets possible, so difference is 0
    if result7 := MinSubsetSumDiff(nums7); result7 != expected7 {
        t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result7)
    } else {
        t.Logf("Test case 7 passed: nums=%v, result=%d", nums7, result7)
    }

    // Test case 8: Single element
    nums8 := []int{5}
    expected8 := 5 // Only one subset possible, so difference is the value itself
    if result8 := MinSubsetSumDiff(nums8); result8 != expected8 {
        t.Errorf("Test case 8 failed: expected %d, got %d", expected8, result8)
    } else {
        t.Logf("Test case 8 passed: nums=%v, result=%d", nums8, result8)
    }
}



func TestCountSubsetSumdiff(t *testing.T) {
	// Test case 1: Subset sum difference is present
	nums1 := []int{1, 1, 2, 3}
	diff1 := 1
	expectedResult1 := 3 // {1,2}, {1,2}, {2,3}
	if result1 := CountSubsetSumdiff(nums1, diff1); result1 != expectedResult1 {
		t.Errorf("Test case 1 failed: nums=%v, diff=%d, expected %d but got %d", nums1, diff1, expectedResult1, result1)
	} else {
		fmt.Printf("Test case 1 passed: nums=%v, diff=%d, result=%d\n", nums1, diff1, result1)
	}

	// Test case 2: Subset sum difference is not present
	nums2 := []int{1, 2, 3, 4}
	diff2 := 5
	expectedResult2 := 0 // No subset sum difference possible
	if result2 := CountSubsetSumdiff(nums2, diff2); result2 != expectedResult2 {
		t.Errorf("Test case 2 failed: nums=%v, diff=%d, expected %d but got %d", nums2, diff2, expectedResult2, result2)
	} else {
		fmt.Printf("Test case 2 passed: nums=%v, diff=%d, result=%d\n", nums2, diff2, result2)
	}

	// Test case 3: Empty array
	nums3 := []int{}
	diff3 := 0
	expectedResult3 := 1 // Empty array has a subset sum difference of 0
	if result3 := CountSubsetSumdiff(nums3, diff3); result3 != expectedResult3 {
		t.Errorf("Test case 3 failed: nums=%v, diff=%d, expected %d but got %d", nums3, diff3, expectedResult3, result3)
	} else {
		fmt.Printf("Test case 3 passed: nums=%v, diff=%d, result=%d\n", nums3, diff3, result3)
	}
}
