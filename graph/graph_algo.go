package graph

import (
    _ "fmt"
    "sort"
    dsf "go-dsa/dsf"
    "go-dsa/pq"
    "math"
)


/*
* data structure used -  graph, disjoint set forest
* graph apis used - Get_numNodes, Get_EdgeList
* dsf apis used - MakeSet, FindSet, UnionSet
*/
func MST_Kruskal(g Graph_t) (res_edges []Edge) {  
    set := make([]*dsf.Dset, g.Get_numNodes()) // disjoint set datastructures for each node
    for i:=0; i<g.Get_numNodes(); i++ {
        set[i]=dsf.MakeSet(i);
    }
    edges := g.Get_EdgeList()
    sort.Slice(edges, func(i,j int) bool {
        return edges[i].weight < edges[j].weight
    })

    for _, edge := range edges {
        if dsf.FindSet(set[edge.u]) != dsf.FindSet(set[edge.v]) {
            // we can consider this edge 
            dsf.UnionSet(set[edge.u], set[edge.v])
            res_edges = append(res_edges, edge)
        }
    }
    return res_edges
}

/*
 * @param arr: adjacency matrix of the graph
 * @param V: number of nodes in graph
 * @param start: start node
 * @returns an array of distance from start node to that node
 * data structure used -  graph 
 * graph apis used - Get_numNodes, Get_adjacent, Get_edge_w
 */
func Dijkstra(arr [][]int, start int) (d []int) {
    g := CreateGraph_m(arr, len(arr));
    if g.Get_numNodes()==0 {
        return
    }
    d = make([]int, g.Get_numNodes()) // distance array from start node
    // initialize single source
    for i:=0; i<g.Get_numNodes(); i++ {
        d[i] = math.MaxInt32
    }
    d[start]=0;
    
    //priority queue
    pq := pq.New_pq(func(i, j interface{}) bool {
                    return d[i.(int)] < d[j.(int)] // Min priority queue on distance array
                })
    pq.Push(start)

    for !pq.IsEmpty() {
        u:=pq.Pop().(int);
        for _, v := range g.Get_adjacent(u) {
            e_w := g.Get_edge_w(u,v);
            if e_w>0 && d[v] > d[u] + e_w { // relax the edge as v is more reachable via u than any other path
                d[v] = d[u] + e_w; // also perform decrease key if already present
                pq.Push(v); /* decrease-key operation(based on the priority key u have set), 
                check less_than_weight, payload is priority here, it will insert in to the pq, if key not found 
                Alternatively pq_insert can also be used */
            }
        }       
    }
    return d;
}


/*
 * @param arr: adjacency matrix of the graph
 * @param V: number of nodes in graph
 * @param start: start node
 * @returns a list of edges in MST
 * data structure used -  graph 
 * graph apis used - Get_numNodes, Get_adjacent, Get_edge_w
 */
func Prims_MST(g Graph_t, start int) (res_edges []Edge) {
    if g.Get_numNodes()==0 {
        return
    }
    parent := make([]int, g.Get_numNodes()) // list of parents per node
    key := make([]int, g.Get_numNodes()) // priority key array from start node, 
    // key[v] means v is more reachable via u so in MST, parent of v should be u, ew(u,v) is less 
    // initialize single source
    for i:=0; i<g.Get_numNodes(); i++ {
        key[i] = math.MaxInt32
    }
    key[start]=0;
    parent[start]=start
    
    //priority queue
    pq := pq.New_pq(func(i, j interface{}) bool {
                    return key[i.(int)] < key[j.(int)] // Min priority queue on key array
                })
    pq.Push(start)

    for !pq.IsEmpty() {
        u:=pq.Pop().(int);
        for _, v := range g.Get_adjacent(u) {
            e_w := g.Get_edge_w(u,v);
            if e_w>0 && key[v] >  e_w { // relax the edge as v is more reachable via u than any other path
                key[v] = e_w; // also perform decrease key if already present
                pq.Push(v); /* decrease-key operation(based on the priority key u have set), */
                parent[v]=u

            }
        }       
    }
    // all vertex keys are relaxed to minimum, now traverse parents to get MST
    for i:=0; i<g.Get_numNodes();i++ {
        // node should not be parent of itself
        if i != parent[i] {
            res_edges = append(res_edges, Edge{
                u: parent[i], // edges will be always from parent to this node
                v: i,
                weight: g.Get_edge_w(parent[i], i),
            })
        }
    }
    return res_edges
}


func IsBipartite(graph [][]int) bool {
    colors := make(map[int]int) // 0 and 1 represent two colors
    for i := range graph {
        if _, exists := colors[i]; !exists && !dfs_color(graph, i, 0, colors) {
            // if coloring is not possible
            return false
        }
    }
    return true
}

func dfs_color(graph [][]int, node, color int, colors map[int]int) bool {
    if c, exists := colors[node]; exists {
        return c == color
    }
    colors[node] = color
    for _, neighbor := range graph[node] {
        if !dfs_color(graph, neighbor, 1-color, colors) {
            return false
        }
    }
    return true
}
