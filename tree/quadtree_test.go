package tree

import (
    "testing"
)

func TestGetBlackPercent(t *testing.T) {
    // Test case 1: Square grid with all black pixels
    grid1 := [][]byte{
        {'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B'},
    }
    root1 := buildQuadTree(grid1)
    expected1 := 100.0 // (16 black pixels / 16 total pixels) * 100
    result1 := Getblackpercent(root1)
    if result1 != expected1 {
        t.Errorf("Test case 1: Expected black percentage %f, but got %f", expected1, result1)
    } else {
        t.Logf("Test case 1: Expected black percentage %f, got %f", expected1, result1)
    }

    // Test case 2: Square grid with all white pixels
    grid2 := [][]byte{
        {'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W'},
    }
    root2 := buildQuadTree(grid2)
    expected2 := 0.0 // (0 black pixels / 16 total pixels) * 100
    result2 := Getblackpercent(root2)
    if result2 != expected2 {
        t.Errorf("Test case 2: Expected black percentage %f, but got %f", expected2, result2)
    } else {
        t.Logf("Test case 2: Expected black percentage %f, got %f", expected2, result2)
    }

    // Test case 3: Rectangular grid with mixed black and white pixels
    grid3 := [][]byte{
        {'B', 'W', 'B', 'W'},
        {'W', 'W', 'B', 'B'},
        {'B', 'W', 'W', 'W'},
        {'B', 'B', 'W', 'B'},
    }
    root3 := buildQuadTree(grid3)
    expected3 := 50.0// (7 black pixels / 16 total pixels) * 100
    result3 := Getblackpercent(root3)
    if result3 != expected3 {
        t.Errorf("Test case 3: Expected black percentage %f, but got %f", expected3, result3)
    } else {
        t.Logf("Test case 3: Expected black percentage %f, got %f", expected3, result3)
    }
}

func TestGetBlackPercent2(t *testing.T) {
    // Test case 4: Larger square grid with all black pixels
    grid4 := [][]byte{
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
    }
    root4 := buildQuadTree(grid4)
    expected4 := 100.0 // (64 black pixels / 64 total pixels) * 100
    result4 := Getblackpercent(root4)
    if result4 != expected4 {
        t.Errorf("Test case 4: Expected black percentage %f, but got %f", expected4, result4)
    } else {
        t.Logf("Test case 4: Expected black percentage %f, got %f", expected4, result4)
    }

    // Test case 5: Larger square grid with all white pixels
    grid5 := [][]byte{
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
    }
    root5 := buildQuadTree(grid5)
    expected5 := 0.0 // (0 black pixels / 64 total pixels) * 100
    result5 := Getblackpercent(root5)
    if result5 != expected5 {
        t.Errorf("Test case 5: Expected black percentage %f, but got %f", expected5, result5)
    } else {
        t.Logf("Test case 5: Expected black percentage %f, got %f", expected5, result5)
    }

    // Test case 6: Larger rectangular grid with mixed black and white pixels
    grid6 := [][]byte{
        {'B', 'W', 'B', 'W', 'B', 'W', 'B', 'W'},
        {'W', 'W', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'B', 'B', 'W', 'B', 'B', 'B', 'B', 'B'},
        {'B', 'B', 'W', 'B', 'B', 'B', 'B', 'B'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
        {'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
        {'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
    }
    root6 := buildQuadTree(grid6)
    expected6 := 51.562500 // (33 black pixels / 64 total pixels) * 100
    result6 := Getblackpercent(root6)
    if result6 != expected6 {
        t.Errorf("Test case 6: Expected black percentage %f, but got %f", expected6, result6)
    } else {
        t.Logf("Test case 6: Expected black percentage %f, got %f", expected6, result6)
    }
}
