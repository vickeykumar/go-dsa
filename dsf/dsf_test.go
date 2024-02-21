package dsf

import "testing"

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