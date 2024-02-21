package avltree

import (
    "fmt"
    "testing"
)

func TestAVLTree(t *testing.T) {
    // Create a new AVL tree
    root := AVL_newNode(10)

    // Insert some keys into the AVL tree
    keys := []int{5, 15, 3, 7, 12, 20}
    for _, key := range keys {
        root = AVL_insert(root, key)
    }

    // Print the AVL tree in order
    fmt.Println("In-order traversal of AVL tree:")
    AVL_printInOrder(root)

    // Test search function
    searchKey := 7
    node := AVL_search(root, searchKey)
    if node == nil || node.key != searchKey {
        t.Errorf("Search for key %d failed", searchKey)
    } else {
        fmt.Printf("\nFound key %d in AVL tree\n", searchKey)
    }

    // Test delete function
    deleteKey := 12
    root = AVL_deleteNode(root, deleteKey)
    fmt.Printf("AVL tree after deleting key %d:\n", deleteKey)
    fmt.Println("In-order AVL: ")
    AVL_printInOrder(root)

    // Print AVL tree in a tree-like format
    fmt.Printf("\nTree-like representation of AVL tree:\n")
    AVL_printAVLTree(root, 0)
}
