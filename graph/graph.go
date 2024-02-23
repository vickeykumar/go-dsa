package graph 

// Define the Graph interface implementing these usefull functions
type Graph_t interface {
	PrintGraph()
	AddEdge(u, v int)
	AddEdge_digraph(u, v int)
	AddEdge_weight(u, v, w int)
	AddEdge_digraph_weight(u, v, w int)
	Get_adjacent(u int) []int
	Get_edge_w(u, v int) int
	Get_numNodes() int
	Get_EdgeList() []Edge // for kruskal dsf
}

type Edge struct {
    u int
    v int
    weight int  // weight of the edge
}