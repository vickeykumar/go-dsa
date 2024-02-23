package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	// Create a new graph with 4 nodes
	g := CreateGraph(4)

	// Add some edges
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	// Check the number of nodes
	if g.numNodes != 4 {
		t.Errorf("Expected number of nodes: 4, got: %d", g.numNodes)
	}
	g.PrintGraph()
	fmt.Println("EdgeList: ", g.Get_EdgeList())
	
	// Check the adjacency matrix
	expectedAdjMatrix := [][]int{
		{0, 1, 1, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
		{0, 0, 1, 0},
	}
	for i := 0; i < len(expectedAdjMatrix); i++ {
		for j := 0; j < len(expectedAdjMatrix[i]); j++ {
			if g.adjMatrix[i][j] != expectedAdjMatrix[i][j] {
				t.Errorf("Expected adjacency matrix[%d][%d]: %d, got: %d", i, j, expectedAdjMatrix[i][j], g.adjMatrix[i][j])
			}
		}
	}

	DFS(g,0)
	BFS(g,0)

	g = CreateDigraph(4)
	// Test adding directed edges
	g.AddEdge_digraph(1, 3)
	g.AddEdge_digraph(3, 0)
	g.AddEdge_digraph(0, 1)
	g.AddEdge_digraph(0, 2)
	g.AddEdge_digraph(1, 2)
	g.AddEdge_digraph(2, 3)
	g.PrintGraph()
	fmt.Println("EdgeList: ", g.Get_EdgeList())
	// Check the modified adjacency matrix
	expectedAdjMatrixDirected := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 0, 1},
		{1, 0, 0, 0},
	}
	for i := 0; i < len(expectedAdjMatrixDirected); i++ {
		for j := 0; j < len(expectedAdjMatrixDirected[i]); j++ {
			if g.adjMatrix[i][j] != expectedAdjMatrixDirected[i][j] {
				t.Errorf("Expected modified adjacency matrix[%d][%d]: %d, got: %d", i, j, expectedAdjMatrixDirected[i][j], g.adjMatrix[i][j])
			}
		}
	}

	DFS(g,0)
	BFS(g,0)
}
