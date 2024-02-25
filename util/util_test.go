package util

import (
    "testing"
)

func TestMax(t *testing.T) {
    // Test case 1: Max of positive integers
    i1, j1 := 10, 20
    expectedResult1 := 20
    if result1 := Max(i1, j1); result1 != expectedResult1 {
        t.Errorf("Test case 1 failed: expected %d but got %d", expectedResult1, result1)
    } else {
        t.Logf("Test case 1 passed: Max(%d, %d) = %d", i1, j1, result1)
    }

    // Test case 2: Max of negative integers
    i2, j2 := -10, -5
    expectedResult2 := -5
    if result2 := Max(i2, j2); result2 != expectedResult2 {
        t.Errorf("Test case 2 failed: expected %d but got %d", expectedResult2, result2)
    } else {
        t.Logf("Test case 2 passed: Max(%d, %d) = %d", i2, j2, result2)
    }

    // Test case 3: Max of positive and negative integers
    i3, j3 := -15, 10
    expectedResult3 := 10
    if result3 := Max(i3, j3); result3 != expectedResult3 {
        t.Errorf("Test case 3 failed: expected %d but got %d", expectedResult3, result3)
    } else {
        t.Logf("Test case 3 passed: Max(%d, %d) = %d", i3, j3, result3)
    }
}

func TestMin(t *testing.T) {
    // Test case 1: Min of positive integers
    i1, j1 := 10, 20
    expectedResult1 := 10
    if result1 := Min(i1, j1); result1 != expectedResult1 {
        t.Errorf("Test case 1 failed: expected %d but got %d", expectedResult1, result1)
    } else {
        t.Logf("Test case 1 passed: Min(%d, %d) = %d", i1, j1, result1)
    }

    // Test case 2: Min of negative integers
    i2, j2 := -10, -5
    expectedResult2 := -10
    if result2 := Min(i2, j2); result2 != expectedResult2 {
        t.Errorf("Test case 2 failed: expected %d but got %d", expectedResult2, result2)
    } else {
        t.Logf("Test case 2 passed: Min(%d, %d) = %d", i2, j2, result2)
    }

    // Test case 3: Min of positive and negative integers
    i3, j3 := -15, 10
    expectedResult3 := -15
    if result3 := Min(i3, j3); result3 != expectedResult3 {
        t.Errorf("Test case 3 failed: expected %d but got %d", expectedResult3, result3)
    } else {
        t.Logf("Test case 3 passed: Min(%d, %d) = %d", i3, j3, result3)
    }
}
