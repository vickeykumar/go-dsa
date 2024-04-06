package tree 

import (
	"reflect"
	"testing"
)

func TestConstructTreeFromLevelOrder(t *testing.T) {
    tests := []struct {
        vals       []int
        inorderSeq []int
    }{
        {[]int{1, 2, 3, 4, 5}, []int{4, 2, 5, 1, 3}},
        {[]int{1, 2, 3, -1, 5}, []int{-1, 2, 5, 1, 3}},
        {[]int{1, -1, 3, -1, 5}, []int{-1, -1, 5, 1, 3}},
    }

    for _, test := range tests {
        tree := ConstructTreeFromLevelOrder(test.vals)
        PrintBinaryTree(tree, 0)
        inorder := BT_InOrder(tree)
        if !reflect.DeepEqual(inorder, test.inorderSeq) {
            t.Errorf("For vals=%v, expected inorder=%v, got=%v (Failed)", test.vals, test.inorderSeq, inorder)
        } else {
            t.Logf("For vals=%v, expected inorder=%v, got=%v (Success)", test.vals, test.inorderSeq, inorder)
        }
    }
}


func TestDiameterOfBinaryTree(t *testing.T) {
    // Test cases with different tree structures
    tests := []struct {
        levelOrder []int
        expected   int
    }{
        {
            levelOrder: []int{1, 2, 3, 4, 5},
            expected:   3,
        },
        {
            levelOrder: []int{1, 2, 3, 4, 5, 6, 7},
            expected:   4,
        },
        {
            levelOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
            expected:   5,
        },
    }

    for _, test := range tests {
        root := ConstructTreeFromLevelOrder(test.levelOrder)
        actual := DiameterOfBinaryTree(root)
        if actual != test.expected {
            t.Errorf("For levelOrder=%v, expected=%d, but got=%d", test.levelOrder, test.expected, actual)
        } else {
            t.Logf("For levelOrder=%v, expected=%d, got=%d (Success)", test.levelOrder, test.expected, actual)
        }
    }
}

func TestMaxPathSum(t *testing.T) {
    // Test cases with level order traversal sequences and expected maximum path sums
    tests := []struct {
        levelOrder []int
        expected   int
    }{
        // Test case 1: Single node tree
        {
            levelOrder: []int{1},
            expected:   1,
        },
        // Test case 2: Binary tree with one level
        {
            levelOrder: []int{1, 2, 3},
            expected:   6,
        },
        // Test case 3: Binary tree with multiple levels
        {
            levelOrder: []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1},
            expected:   29,
        },
        // Add more test cases as needed
    }

    // Iterate over the test cases
    for i, test := range tests {
        // Construct the tree from the level order traversal sequence
        tree := ConstructTreeFromLevelOrder(test.levelOrder)
        PrintBinaryTree(tree, 0)
        // Calculate the maximum path sum of the tree
        maxPathSum := MaxPathSum(tree)

        // Compare the calculated maximum path sum with the expected value
        if maxPathSum != test.expected {
            t.Errorf("Test case %d failed: expected maximum path sum %d, but got %d", i+1, test.expected, maxPathSum)
        } else {
            t.Logf("Test case %d passed: maximum path sum of the tree is %d", i+1, maxPathSum)
        }
    }
}
