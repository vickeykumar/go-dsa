package graph

import (
	"fmt"
)

// undirected graph type
type Graph struct {
	numNodes int
	adjMatrix [][]int
}

type Digraph = Graph

// Define a method for the original type
func (g *Graph) PrintGraph() {
	fmt.Println("Number of nodes:", g.numNodes)
	fmt.Println("Adjacency matrix:")
	for _, row := range g.adjMatrix {
		fmt.Println(row)
	}
}


// Create a new graph with n nodes
func CreateGraph(V int) *Graph {
	adjMatrix := make([][]int, V)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, V)
	}
	return &Graph {
		numNodes: V,
		adjMatrix: adjMatrix,
	}
}

// Create a new graph with n nodes
func CreateDigraph(V int) *Digraph {
	return CreateGraph(V)
}

func CreateGraph_m(adjm [][]int, V int) *Graph {
	return &Graph {
		numNodes: V,
		adjMatrix: adjm,
	}
}

// Add an edge from node u to node v, and v to u as it is undirected
func (g *Graph) AddEdge(u, v int) {
	g.adjMatrix[u][v] = 1
	g.adjMatrix[v][u] = 1
}

// Add an edge from node u to node v
func (g *Graph) AddEdge_digraph(u, v int) {
	g.adjMatrix[u][v] = 1
}

// Add an edge from node u to node v, and v to u as it is undirected
func (g *Graph) AddEdge_weight(u, v, w int) {
	g.adjMatrix[u][v] = w
	g.adjMatrix[v][u] = w
}

// Add an edge from node u to node v, and v to u as it is undirected
func (g *Graph) AddEdge_digraph_weight(u, v, w int) {
	g.adjMatrix[u][v] = w
}

func (g *Graph) Get_adjacent(u int) (result []int) {
	for v, w := range g.adjMatrix[u] {
		if w!=0 {
			result = append(result, v)
		}
	}
	return
}

func (g *Graph) Get_edge_w(u, v int) int {
	return g.adjMatrix[u][v]
}

func (g *Graph) Get_numNodes() int {
	return g.numNodes
}

func (g *Graph) Get_EdgeList() (edges []Edge) {
	for u, adjm_u := range g.adjMatrix {
		for v, w := range adjm_u {
			if w!=0 {
				edges = append(edges, Edge{
					u: u,
					v: v,
					weight: w,
				})
			}
		}
	}
	return
}
