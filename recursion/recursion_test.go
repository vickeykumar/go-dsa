package recursion

import (
    "fmt"
    "reflect"
    "testing"
    "strings"
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
    fmt.Println("TestPermutehelper: ", result)
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

func TestGenSubsets(t *testing.T) {
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
    result := GenSubsets(nums, []int{},0)
    fmt.Println("powerset: ", result)
    if !slicesEqual(result, expected) {
        t.Errorf("Powersets(%v) = %v, want %v", nums, result, expected)
    }
}


func TestCombinations(t *testing.T) {
    // Test case 1: n=4, k=2
    n1, k1 := 4, 2
    expected1 := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
    result1 := Combinations(n1, 1, k1, []int{})
    if !slicesEqual(result1, expected1) {
        t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
    } else {
        t.Logf("Test case 1 passed. Result: %v", result1)
    }

    // Test case 2: n=5, k=3
    n2, k2 := 5, 3
    expected2 := [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}}
    result2 := Combinations(n2, 1, k2, []int{})
    if !slicesEqual(result2, expected2) {
        t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
    } else {
        t.Logf("Test case 2 passed. Result: %v", result2)
    }

    // Additional test case: n=3, k=1
    n3, k3 := 3, 1
    expected3 := [][]int{{1}, {2}, {3}}
    result3 := Combinations(n3, 1, k3, []int{})
    if !slicesEqual(result3, expected3) {
        t.Errorf("Additional test case failed. Expected %v, got %v", expected3, result3)
    } else {
        t.Logf("Additional test case passed. Result: %v", result3)
    }
}

func TestCombinations2(t *testing.T) {
    // Test case 1: items=[1,2,3], k=2
    items1 := []int{1, 2, 3}
    k1 := 2
    expected1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
    result1 := Combinations2(items1, 0, k1, []int{})
    if !slicesEqual(result1, expected1) {
        t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
    } else {
        t.Logf("Test case 1 passed. Result: %v", result1)
    }
}

func TestCombinationsSum(t *testing.T) {
    // Test case 1: items=[2,3,6,7], target=7
    items1 := []int{2, 3, 6, 7}
    target1 := 7
    expected1 := [][]int{{2, 2, 3}, {7}}
    result1 := CombinationsSum(items1, 0, target1, []int{})
    if !slicesEqual(result1, expected1) {
        t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
    } else {
        t.Logf("Test case 1 passed. Result: %v", result1)
    }
}


func TestSolveNQueens(t *testing.T) {
    // Test case 1: n=4
    n1 := 4
    expected1 := [][]string{
        {".Q..", "...Q", "Q...", "..Q."},
        {"..Q.", "Q...", "...Q", ".Q.."},
    }
    result1 := SolveNQueens(n1)
    if !equal2DStringSlices(result1, expected1) {
        t.Errorf("Test case 1 failed. Expected:") 
        printChessboard(expected1)
        t.Errorf("got : ")
        printChessboard(result1)
    } else {
        fmt.Println("Test case 1 passed.")
        printChessboard(result1)
    }

    // Test case 2: n=5
    n2 := 5
    expected2 := [][]string{
        {"Q....","..Q..","....Q",".Q...","...Q."},
        {"Q....","...Q.",".Q...","....Q","..Q.."},
        {".Q...","...Q.","Q....","..Q..","....Q"},
        {".Q...","....Q","..Q..","Q....","...Q."},
        {"..Q..","Q....","...Q.",".Q...","....Q"},
        {"..Q..","....Q",".Q...","...Q.","Q...."},
        {"...Q.","Q....","..Q..","....Q",".Q..."},
        {"...Q.",".Q...","....Q","..Q..","Q...."},
        {"....Q",".Q...","...Q.","Q....","..Q.."},
        {"....Q","..Q..","Q....","...Q.",".Q..."},
    }
    result2 := SolveNQueens(n2)
    if !equal2DStringSlices(result2, expected2) {
        t.Errorf("Test case 2 failed. Expected:") 
        printChessboard(expected2)
        t.Errorf("got : ")
        printChessboard(result2)
    } else {
        fmt.Println("Test case 2 passed.")
        printChessboard(result2)
    }
}

// Helper function to compare two [][]string slices for equality
func equal2DStringSlices(slice1, slice2 [][]string) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    for i := 0; i < len(slice1); i++ {
        if !equalStringSlices(slice1[i], slice2[i]) {
            return false
        }
    }
    return true
}

// Helper function to compare two []string slices for equality
func equalStringSlices(slice1, slice2 []string) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    for i := 0; i < len(slice1); i++ {
        if slice1[i] != slice2[i] {
            return false
        }
    }
    return true
}

// Helper function to print chessboard in a well-formatted format
func printChessboard(chessboard [][]string) {
    fmt.Println("Chessboards:")
    for i, row := range chessboard {
        fmt.Printf("\nChessboard %d:\n",i+1)
        fmt.Println(strings.Join(row, "\n"))
    }
}
