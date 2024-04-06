package tree

import (
	"fmt"
	"math"
	"go-dsa/util"
	"go-dsa/queue"
)

const COUNT=4

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func ConstructTreeFromLevelOrder(vals []int) *TreeNode {
    if len(vals) == 0 {
        return nil
    }

    root := &TreeNode{Val: vals[0]}
    q := queue.NewQue()
    q.Enque(root)
    index := 1

    for !q.IsEmpty() {
        node := q.Deque().(*TreeNode) // get first node from queue and add l,r child

        if index < len(vals) {
            node.Left = &TreeNode{Val: vals[index]}
            q.Enque(node.Left)
        }
        index++

        if index < len(vals) {
            node.Right = &TreeNode{Val: vals[index]}
            q.Enque(node.Right)
        }
        index++
    }

    return root
}


func BT_InOrder(node *TreeNode) (inorder []int) {
    if node == nil {
        return;
    }

    leftorder := BT_InOrder(node.Left);
    inorder = append(inorder, leftorder...);
    inorder = append(inorder, node.Val);
    rightorder := BT_InOrder(node.Right);
    inorder = append(inorder, rightorder...);
    return
}

/**
 * Prints an binary tree in a diagram format.
 *
 * @param root A pointer to the root node of the tree.
 * @param space The amount of space to print before the current node.
 */
func PrintBinaryTree(root *TreeNode, space int) {
    // Base case: the tree is empty
    if (root == nil) {
        return;
    }

    // Increase the amount of space for the next level
    space += COUNT;

    // Recursively print the right subtree
    PrintBinaryTree(root.Right, space);

    // Print the current node
    fmt.Printf("\n");
    for i := COUNT; i < space; i++ {
        fmt.Printf(" ");
    }
    fmt.Printf("%d\n", root.Val);

    // Recursively print the left subtree
    PrintBinaryTree(root.Left, space);
}


/*
tree algo
*/
func Findheight(root *TreeNode, dim *int) int {
    if root==nil {
        return 0
    }
    // left sides height and right sides height
    lh := Findheight(root.Left, dim)
    rh := Findheight(root.Right, dim)
    *dim = util.Max(lh+rh, *dim)
    return util.Max(lh,rh)+1
}

func DiameterOfBinaryTree(root *TreeNode) int {
    dim:=0
    Findheight(root, &dim)
    //fmt.Println(h, dim)
    return dim
}


func MaxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    maxPathSumRecursive(root, &maxSum)
    return maxSum
}

func maxPathSumRecursive(node *TreeNode, maxSum *int) int {
    if node == nil {
        return 0
    }

    leftSum := maxPathSumRecursive(node.Left, maxSum)
    rightSum := maxPathSumRecursive(node.Right, maxSum)

    // Calculate the maximum path sum including the current node, along the height verticaly
    maxSingle := util.Max(node.Val, util.Max(node.Val+leftSum, node.Val+rightSum))

    // Calculate the maximum path sum that goes through the current node, or new max starts from this node itself
    maxWithNode := util.Max(maxSingle, node.Val+leftSum+rightSum)

    // Update the overall maximum path sum found so far
    *maxSum = util.Max(*maxSum, maxWithNode)

    // Return the maximum path sum including the current node
    return maxSingle
}
