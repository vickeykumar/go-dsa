package dsf

import (
    "reflect"
    "sort"
    "testing"
)

func TestDisjointSet(t *testing.T) {
    // Create some nodes
    node1 := MakeSet(1)
    node2 := MakeSet(2)
    node3 := MakeSet(3)

    // Test MakeSet function
    if node1.parent != node1 || node1.rank != 0 || node1.data != 1 {
        t.Errorf("MakeSet(1) failed: got %v, want parent=%v, rank=%d, data=%v", node1, node1, 0, 1)
    }

    // Test FindSet function
    set1 := FindSet(node1)
    set2 := FindSet(node2)
    if set1 != node1 || set2 != node2 {
        t.Errorf("FindSet failed: expected node's parent to be itself")
    }

    // Test unionSet function
    UnionSet(node1, node2)
    if FindSet(node1) != FindSet(node2) {
        t.Errorf("unionSet failed: expected node1's parent to be same as node2's")
    }

    // Test FindSet after unionSet
    set3 := FindSet(node3)
    if set3 != node3 {
        t.Errorf("FindSet failed: expected node's parent to be itself")
    }

    UnionSet(node1, node3)
    if FindSet(node1) != FindSet(node3) {
        t.Errorf("unionSet failed: expected node3's parent to be same as node2's")
    }
}

func TestAccountsMerge(t *testing.T) {
    tests := []struct {
        accounts [][]string
        expected [][]string
    }{
        // Test case with single account
        {
            accounts: [][]string{
                {"John", "john@example.com"},
            },
            expected: [][]string{
                {"John", "john@example.com"},
            },
        },
        // Test case with multiple accounts
        {
            accounts: [][]string{
                {"Jane", "jane@example.com"},
                {"John", "john@example.com", "john.doe@example.com"},
            },
            expected: [][]string{
                {"Jane", "jane@example.com"},
                {"John", "john@example.com", "john.doe@example.com"},
            },
        },
        // Test case with accounts having overlapping emails
        {
            accounts: [][]string{
                {"John", "john@example.com"},
                {"Jane", "john@example.com", "jane@example.com"},
            },
            expected: [][]string{
                {"John", "jane@example.com", "john@example.com"},
            },
        },
        // Test case with empty accounts
        {
            accounts: [][]string{},
            expected: [][]string{},
        },
        // Add more test cases as needed
    }

    for i, test := range tests {
        merged := AccountsMerge(test.accounts)
        for i, _ := range test.expected {
            if len(test.expected[i])>1 {
                sort.Strings(test.expected[i][1:]) // sort inplace so that we can compare
            }
        }
        
        if !reflect.DeepEqual(merged, test.expected) {
            t.Errorf("Test: %d For accounts %#v, expected %#v, but got %#v", i, test.accounts, test.expected, merged)
        }
    }
}
