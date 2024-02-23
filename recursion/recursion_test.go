package recursion

import (
    "fmt"
    "reflect"
    "testing"
)

func slicesEqual(s1, s2 [][]int) bool {
    if len(s1) != len(s2) {
        return false
    }
    for _, x := range s1 {
        if !sliceContains(s2, x) {
            return false
        }
    }
    for _, x := range s2 {
        if !sliceContains(s1, x) {
            return false
        }
    }
    return true
}

func sliceContains(s [][]int, x []int) bool {
    for _, y := range s {
        if reflect.DeepEqual(x, y) {
            return true
        }
    }
    return false
}

func TestPowersets(t *testing.T) {
    nums := []int{1, 2, 3}
    expected := [][]int{
        {},      // Empty set
        {1},     // Sets with single elements
        {2},
        {3},
        {1, 2},  // Sets with two elements
        {1, 3},
        {2, 3},
        {1, 2, 3},  // Set with all elements
    }
    result := Powersets(nums, 0, len(nums)-1)
    fmt.Println("powerset: ", result)
    if !slicesEqual(result, expected) {
        t.Errorf("Powersets(%v) = %v, want %v", nums, result, expected)
    }
}

func TestPermutehelper(t *testing.T) {
    nums := []int{1, 2, 3}
    expected := [][]int{
        {1, 2, 3},
        {1, 3, 2},
        {2, 1, 3},
        {2, 3, 1},
        {3, 2, 1},
        {3, 1, 2},
    }
    result := Permutehelper(nums, 0)
    fmt.Println("TestPermutehelper2: ", result)
    if !slicesEqual(result, expected) {
        t.Errorf("Permutehelper(%v) = %v, want %v", nums, result, expected)
    }
}

func TestLookupAndInsert(t *testing.T) {
    set := []int{1, 2, 3}
    c := 4
    expectedSet := []int{1, 2, 3, 4}
    expectedFound := false
    resultSet, resultFound := lookup_and_insert(set, c)
    if !reflect.DeepEqual(resultSet, expectedSet) || resultFound != expectedFound {
        t.Errorf("lookup_and_insert(%v, %v) = (%v, %v), want (%v, %v)", set, c, resultSet, resultFound, expectedSet, expectedFound)
    }
}

func TestPermutehelper2(t *testing.T) {
    nums := []int{1, 2, 3}
    expected := [][]int{
        {1, 2, 3},
        {1, 3, 2},
        {2, 1, 3},
        {2, 3, 1},
        {3, 2, 1},
        {3, 1, 2},
    }
    result := Permutehelper2(nums, 0)
    fmt.Println("TestPermutehelper2: ", result)
    if !slicesEqual(result, expected) {
        t.Errorf("Permutehelper2(%v) = %v, want %v", nums, result, expected)
    }
}
