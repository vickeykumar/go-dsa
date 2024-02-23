package graph 

import (
	"fmt"
	"go-dsa/stack"
	"go-dsa/queue"
)


/* push-> pop,visit,push(nighbors)
 * g is a Graph interface, as it has all the usefull functions
 * u can use any graph type implementing Graph interface */
func DFS(g Graph_t, start int) {
	fmt.Printf("DFS: ")
    st := stack.NewStack()
    st.Push(start);
    // array of all the visited nodes
    visited := make([]bool, g.Get_numNodes())
    visited[start] = true
    for !st.IsEmpty() {
        u := st.Pop().(int) // top node
        if !visited[u] {
        	visited[u] = true
        }
        fmt.Printf("%d ",u);

        for _, v := range g.Get_adjacent(u) {
        	if !visited[v] {
        		st.Push(v)
        		visited[v] = true
        	}
        }
    } 
    fmt.Printf("\n")
}



// enque, deque, enque and visit adjacent nodes as we go
func BFS(g Graph_t, start int) {
	fmt.Printf("BFS: ")
  // Queue for storing the vertices to visit
  q := queue.NewQue()
  q.Enque(start)
  // array of all the visited nodes
  visited := make([]bool, g.Get_numNodes())
  visited[start] = true

  for !q.IsEmpty() {
        u := q.Deque().(int)
        fmt.Printf("%d ",u);
 
        for _, v := range g.Get_adjacent(u) {
        	if !visited[v] {
        		q.Enque(v)
        		visited[v] = true;
        	}
        }
    }
    fmt.Printf("\n")
}

