package graph

import (
	"fmt"
	"testing"
)

func TestMST_Kruskal_Graph(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	// Create a graph instance
	g := CreateGraph_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := MST_Kruskal(g)
	fmt.Println ("MST_Kruskal: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 0, v: 3, weight: 6},
		{u: 1, v: 4, weight: 5},
		{u: 1, v: 2, weight: 3},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not, check for the reverse ordering
			revEdge := Edge{u: edge.v, v: edge.u, weight: edge.weight}
			if _, ok := expectedMap[revEdge]; ok {
				// If found, remove it from the expected map
				delete(expectedMap, revEdge)
			} else {
				// If not found in either ordering, it's an unexpected edge
				t.Errorf("Unexpected edge in MST: %v", edge)
			}
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}



func TestMST_Kruskal_GraphLL(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	// Create a graph_ll instance
	g := CreateGraph_ll_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := MST_Kruskal(g)
	fmt.Println ("MST_Kruskal: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 0, v: 3, weight: 6},
		{u: 1, v: 4, weight: 5},
		{u: 1, v: 2, weight: 3},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not, check for the reverse ordering
			revEdge := Edge{u: edge.v, v: edge.u, weight: edge.weight}
			if _, ok := expectedMap[revEdge]; ok {
				// If found, remove it from the expected map
				delete(expectedMap, revEdge)
			} else {
				// If not found in either ordering, it's an unexpected edge
				t.Errorf("Unexpected edge in MST: %v", edge)
			}
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}



func TestMST_Kruskal_Digraph(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{0, 0, 3, 8, 5},
		{0, 0, 0, 0, 7},
		{0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0},
	}

	// Create a directed graph instance
	g := CreateGraph_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := MST_Kruskal(g)
	fmt.Println ("MST_Kruskal: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 1, v: 2, weight: 3},
		{u: 1, v: 4, weight: 5},
		{u: 0, v: 3, weight: 6},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not found, it's an unexpected edge
			t.Errorf("Unexpected edge in MST: %v", edge)
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}



func TestMST_Kruskal_Digraph_ll(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{0, 0, 3, 8, 5},
		{0, 0, 0, 0, 7},
		{0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0},
	}

	// Create a directed graph instance
	g := CreateGraph_ll_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := MST_Kruskal(g)
	fmt.Println ("MST_Kruskal: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 1, v: 2, weight: 3},
		{u: 1, v: 4, weight: 5},
		{u: 0, v: 3, weight: 6},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not found, it's an unexpected edge
			t.Errorf("Unexpected edge in MST: %v", edge)
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}

func TestDijkstra(t *testing.T) {
    // Test case 1: Directed graph with positive edge weights
    arr1 := [][]int{
        {0, 4, 0, 0, 0, 0, 0, 8, 0},
        {4, 0, 8, 0, 0, 0, 0, 11, 0},
        {0, 8, 0, 7, 0, 4, 0, 0, 2},
        {0, 0, 7, 0, 9, 14, 0, 0, 0},
        {0, 0, 0, 9, 0, 10, 0, 0, 0},
        {0, 0, 4, 14, 10, 0, 2, 0, 0},
        {0, 0, 0, 0, 0, 2, 0, 1, 6},
        {8, 11, 0, 0, 0, 0, 1, 0, 7},
        {0, 0, 2, 0, 0, 0, 6, 7, 0},
    }
    expected1 := []int{0, 4, 12, 19, 21, 11, 9, 8, 14}
    start1 := 0
    CreateGraph_m(arr1, len(arr1)).PrintGraph()
    d1 := Dijkstra(arr1, start1)
    fmt.Printf("distance array after Dijkstra: %v \nExpected: %v\n", d1, expected1)
    for i := range d1 {
        if d1[i] != expected1[i] {
            t.Errorf("Test case 1 failed: Expected distance to node %d to be %d, got %d", i, expected1[i], d1[i])
        }
    }

    // Test case 2: Directed graph with negative edge weights
    arr2 := [][]int{
        {0, -1, 4, 0, 0},
        {0, 0, 3, 2, 2},
        {0, 0, 0, 0, 0},
        {0, 1, 5, 0, 0},
        {0, 0, 0, -3, 0},
    }
    expected2 := []int{0, -1, 2, -2, 1}
    start2 := 0
    CreateGraph_m(arr2, len(arr2)).PrintGraph()
    d2 := Dijkstra(arr2, start2)
    fmt.Printf("distance array after Dijkstra: %v \nExpected: %v\n", d2, expected2)

    // Test case 3: Empty graph
    arr3 := [][]int{}
    start3 := 0
    d3 := Dijkstra(arr3, start3)
    if len(d3) != 0 {
        t.Errorf("Test case 3 failed: Expected empty result for an empty graph")
    }

    // Test case 4: Single-node graph
    arr4 := [][]int{{}}
    expected4 := []int{0}
    start4 := 0
    d4 := Dijkstra(arr4, start4)
    fmt.Printf("distance array after Dijkstra: %v \nExpected: %v\n", d4, expected4)
    for i := range d4 {
        if d4[i] != expected4[i] {
            t.Errorf("Test case 4 failed: Expected distance to node %d to be %d, got %d", i, expected4[i], d4[i])
        }
    }
}


func Test_Prims_MST_Graph(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	// Create a graph instance
	g := CreateGraph_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Prims's algorithm
	resEdges := Prims_MST(g, 0)
	fmt.Println ("Prims_MST: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 0, v: 3, weight: 6},
		{u: 1, v: 4, weight: 5},
		{u: 1, v: 2, weight: 3},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not, check for the reverse ordering
			revEdge := Edge{u: edge.v, v: edge.u, weight: edge.weight}
			if _, ok := expectedMap[revEdge]; ok {
				// If found, remove it from the expected map
				delete(expectedMap, revEdge)
			} else {
				// If not found in either ordering, it's an unexpected edge
				t.Errorf("Unexpected edge in MST: %v", edge)
			}
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}



func Test_Prims_MST_GraphLL(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	// Create a graph_ll instance
	g := CreateGraph_ll_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := Prims_MST(g, 0)
	fmt.Println ("Prims_MST: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 0, v: 3, weight: 6},
		{u: 1, v: 4, weight: 5},
		{u: 1, v: 2, weight: 3},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not, check for the reverse ordering
			revEdge := Edge{u: edge.v, v: edge.u, weight: edge.weight}
			if _, ok := expectedMap[revEdge]; ok {
				// If found, remove it from the expected map
				delete(expectedMap, revEdge)
			} else {
				// If not found in either ordering, it's an unexpected edge
				t.Errorf("Unexpected edge in MST: %v", edge)
			}
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}



func Test_Prims_MST_Digraph(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{0, 0, 3, 8, 5},
		{0, 0, 0, 0, 7},
		{0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0},
	}

	// Create a directed graph instance
	g := CreateGraph_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := Prims_MST(g, 0)
	fmt.Println ("Prims_MST: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 1, v: 2, weight: 3},
		{u: 1, v: 4, weight: 5},
		{u: 0, v: 3, weight: 6},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not found, it's an unexpected edge
			t.Errorf("Unexpected edge in MST: %v", edge)
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}



func Test_Prims_MST_Digraph_ll(t *testing.T) {
	// Define the adjacency matrix
	adjMatrix := [][]int{
		{0, 2, 0, 6, 0},
		{0, 0, 3, 8, 5},
		{0, 0, 0, 0, 7},
		{0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0},
	}

	// Create a directed graph instance
	g := CreateGraph_ll_m(adjMatrix, 5)
	g.PrintGraph()

	// Calculate MST using Kruskal's algorithm
	resEdges := Prims_MST(g, 0)
	fmt.Println ("Prims_MST: ", resEdges)

	// Define expected result
	expectedEdges := []Edge{
		{u: 0, v: 1, weight: 2},
		{u: 1, v: 2, weight: 3},
		{u: 1, v: 4, weight: 5},
		{u: 0, v: 3, weight: 6},
	}

	// Create a map to track presence of expected edges
	expectedMap := make(map[Edge]bool)
	for _, edge := range expectedEdges {
		expectedMap[edge] = true
	}

	// Check each edge in the result
	for _, edge := range resEdges {
		// Check if the edge matches any expected edge
		if _, ok := expectedMap[edge]; ok {
			// If yes, remove it from the expected map
			delete(expectedMap, edge)
		} else {
			// If not found, it's an unexpected edge
			t.Errorf("Unexpected edge in MST: %v", edge)
		}
	}

	// Check if all expected edges were found
	if len(expectedMap) != 0 {
		t.Errorf("Expected edges not found in MST: %v", expectedMap)
	}
}
