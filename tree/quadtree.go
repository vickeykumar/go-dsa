package tree

import (
    "fmt"
)
/*

**--
**--
**-*
*--* 

* B
- W

i internal node

root:                           i
children:              B   W       i       i
2nd chilren:                    B B B W   W B W B

find Black% in the quad tree

pow(4,h-1-l) num of nodes at each leaf

6/13
9/16

*/

type point struct {
    i int
    j int
}

type QuadTreeNode struct {
    color bool // 1 for black / 0 for W
    isleaf bool
    children [4]*QuadTreeNode
    points  []point     // num of points this node has
} 

func buildTree(matrix [][]byte, row, col, size int) *QuadTreeNode {
    // Base case: if the size is 1, create a leaf node
    if size == 1 {
        return &QuadTreeNode{color: matrix[row][col] == 'B',
                            isleaf: true,
                            points: []point{point{row, col}},
                            }
    }

    // Check if all elements in the current region have the same color
    sameColor := true
    points := []point{}
    for i := row; i < row+size; i++ {
        for j := col; j < col+size; j++ {
            if matrix[i][j] != matrix[row][col] {
                sameColor = false
                break
            }
            points = append(points, point{row, col})
        }
        if !sameColor {
            break
        }
    }

    // If all elements have the same color, create a leaf node
    if sameColor {
        return &QuadTreeNode{color: matrix[row][col] == 'B',
                             isleaf: true,
                             points: points,
                         }
    }

    // Otherwise, divide the current region into four quadrants and recursively build the tree
    halfSize := size / 2
    root := &QuadTreeNode{}
    root.children[0] = buildTree(matrix, row, col, halfSize)
    root.children[1] = buildTree(matrix, row, col+halfSize, halfSize)
    root.children[2] = buildTree(matrix, row+halfSize, col, halfSize)
    root.children[3] = buildTree(matrix, row+halfSize, col+halfSize, halfSize)

    return root
}

func buildQuadTree(matrix [][]byte) *QuadTreeNode {
    return buildTree(matrix, 0, 0, len(matrix))
}

func percentdfs(root *QuadTreeNode, total *int) (blacknum int) {
    if root==nil {
        return 0
    }
    
    if root.isleaf {
        *total += len(root.points) // only leaves will have num of points this node consists of
        if root.color==true {
            return len(root.points)
        }
    }
    
    for _, child := range root.children {
        blacknum += percentdfs(child, total)
    }
    return blacknum
}

// number of edges from root to farthest node
func calculateHeight(root *QuadTreeNode) int {
    if root == nil {
        return -1
    }
    maxChildHeight := 0
    for _, child := range root.children {
        childHeight := calculateHeight(child)
        if childHeight > maxChildHeight {
            maxChildHeight = childHeight
        }
    }
    return maxChildHeight + 1
}

func Getblackpercent(root *QuadTreeNode) float64 {
    total := 0
    blacknum := percentdfs(root, &total)
    fmt.Println("blacknum total", blacknum, total)
    return (float64(blacknum)/float64(total)) *100
}
