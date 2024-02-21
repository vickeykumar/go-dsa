package avltree

import (
	"fmt"
)

const COUNT=4

type AVLNode struct {
	key int
	height int
	left *AVLNode
	right *AVLNode
}

func max(i,j int) int {
	if i>j {
		return i
	}
	return j
}

/* returns new node of AVLTree
*/
func AVL_newNode(keyIn int) *AVLNode{
	return &AVLNode{
		key: keyIn,
		height: 0,
		left: nil,
		right: nil,
	}
}
// with recursive call no need of the member functions

/* returns height*/
func AVL_getHeight(node *AVLNode) int {
	if node == nil {
		return 0;
	}
	return node.height
}

/* get balance factor */
func AVL_getBalance(node *AVLNode) int {
	if node == nil {
		return 0;
	}
	return (AVL_getHeight(node.left) - AVL_getHeight(node.right))
}

/**
 * Performs a left rotation on a node in an AVL tree.
 *
 * @param x A pointer to the node to be rotated.
 * @return A pointer to the root node of the subtree after the rotation.
 *
 *          x                 y
 *         / \               / \
 *        T1  y     =>      x   T3
 *           / \           / \
 *          T2 T3         T1 T2
 */

func AVL_leftRotate(x *AVLNode) *AVLNode {
    y := x.right;
    T2 := y.left;

    // Perform the rotation
    y.left = x;
    x.right = T2;

    // Update the heights of the nodes
    x.height = 1 + max(AVL_getHeight(x.left), AVL_getHeight(x.right));
    y.height = 1 + max(AVL_getHeight(y.left), AVL_getHeight(y.right));

    // Return the new root node
    return y;
}


/**
 * Performs a right rotation on a node in an AVL tree.
 *
 * @param y A pointer to the node to be rotated.
 * @return A pointer to the root node of the subtree after the rotation.
 *
 *        y                   x
 *       / \                 / \
 *      x   T3      =>      T1  y
 *     / \                     / \
 *    T1 T2                   T2 T3
 */
func AVL_rightRotate(y *AVLNode) *AVLNode {
    x := y.left
    T2 := x.right

    // perform the rotation
    x.right = y 
    y.left = T2

    // update the heights of x and y
    x.height = 1+max(AVL_getHeight(x.left), AVL_getHeight(x.right))
    y.height = 1+max(AVL_getHeight(y.left), AVL_getHeight(y.right))

    // return new root
    return x
}

/* 
 * try to search the key in avl tree,
 * search inleft subtree if key is < than the node key, else search in right subtree
 */
func AVL_search(node *AVLNode, key int) *AVLNode {
    // If the tree is empty or the key is not found, return nil
    if node == nil || node.key == key {
        return node;
    }

    // If the key is less than the key in the root node, search the left subtree
    if (key < node.key) {
        return AVL_search(node.left, key);
    } else {
    	// If the key is greater than the key in the root node, search the right subtree
        return AVL_search(node.right, key);
    }
}


/**
 * Finds the node with the minimum key in a given subtree.
 *
 * @param node A pointer to the root node of the subtree.
 * @return A pointer to the node with the minimum key.
 */
func AVL_minValueNode(node *AVLNode) *AVLNode {
    for node.left != nil {
        node = node.left
    }
    return node
}

/**
 * Inserts a key into an AVL tree.
 *
 * @param node A pointer to the root node of the AVL tree.
 * @param key The key to be inserted.
 * @return A pointer to the root node of the AVL tree after the insertion.
 */
func AVL_insert(node *AVLNode, key int) *AVLNode {
    // performs normal BST insert 
    if node == nil {
        return AVL_newNode(key);
    }

    // first insert then balance the tree
    if (key < node.key) {
    	// insert recursively in left subtree
        node.left = AVL_insert(node.left, key);
    } else if (key > node.key) {
        node.right = AVL_insert(node.right, key);
    } else {
        return node; // key should be unique, so that unique search can be availed
    }

    node.height = 1 + max(AVL_getHeight(node.left), AVL_getHeight(node.right));

    // Calculate the balance factor of the root node
    balance := AVL_getBalance(node);

    // If the balance factor is greater than 1 and the key to be inserted is less than the key in the left child,
    // perform a right rotation on the root node
    /*
        LL rotation: tree becomes unbalance after insertion in left subtree of left child of root node. so name is LL rotation.
    */
    if (balance > 1 && key < node.left.key) {
        return AVL_rightRotate(node);
    }

    // If the balance factor is less than -1 and the key to be inserted is greater than the key in the right child,
    // perform a left rotation on the root node
    /*
        RR rotation: tree becomes unbalanced by insertion in to right subtree of right child of the root node.
    */
    if (balance < -1 && key > node.right.key) {
        return AVL_leftRotate(node);
    }

    // If the balance factor is greater than 1 and the key to be inserted is greater than the key in the left child,
    // perform a left rotation on the left child first and then a right rotation on the root node
    /*
        LR Rotation : there was insertion in direction root->left->right
        steps first make LL-> by doing left rotation at child, then right rotation at root
        leftrotate(child) -> right rotate(root of subtree)
    */
    if (balance > 1 && key > node.left.key) {
        node.left = AVL_leftRotate(node.left);
        return AVL_rightRotate(node);
    }

    // If the balance factor is less than -1 and the key to be inserted is less than the key in the right child,
    // perform a right rotation on the right child first and then a left rotation on the root node
    /*
        RL rotation: sequence insertion in R(root) and left(child)
        right roatation(child) -> left rotation(root)
        i.e. first make RR and then do left rotation.
    */
    if (balance < -1 && key < node.right.key) {
        node.right = AVL_rightRotate(node.right);
        return AVL_leftRotate(node);
    }

    // Return the root node
    return node;
}

/**
 * Deletes a node from an AVL tree.
 *
 * @param root A pointer to the root node of the tree.
 * @param key The key of the node to be deleted.
 * @return A pointer to the root node of the tree after the deletion.
 */
func AVL_deleteNode(root *AVLNode, key int) *AVLNode {
    // perform bst deletion
    if (root == nil) {
        return root;
    }

    // Recursively search for the node to be deleted
    if (key < root.key) {
        root.left = AVL_deleteNode(root.left, key);
    } else if (key > root.key) {
        root.right = AVL_deleteNode(root.right, key);
    } else {
    	// found
        // Case 1: the node has no children
        if root.left == nil && root.right == nil {
            return nil;
        }
        // Case 2: the node has one child
        if root.left == nil {
            return root.right
        } else if root.right == nil {
            return root.left;
        }
        // Case 3: the node has two children
        // find the successor from right subtree,
        // update this roots key with successors key and ,
        // delete the successor from right subtree.
        successor := AVL_minValueNode(root.right);
        root.key = successor.key;
        root.right = AVL_deleteNode(root.right, successor.key);
    }

    // Update the height of the current node
    root.height = 1 + max(AVL_getHeight(root.left), AVL_getHeight(root.right))

    // Check if the tree is unbalanced and perform the necessary rotations
    balance := AVL_getBalance(root);
    if (balance > 1 && AVL_getBalance(root.left) >= 0) {
        return AVL_rightRotate(root);
    }
    if (balance > 1 && AVL_getBalance(root.left) < 0) {
        root.left = AVL_leftRotate(root.left);
        return AVL_rightRotate(root);    
    }
    if (balance < -1 && AVL_getBalance(root.right) <= 0) {
        return AVL_leftRotate(root);
    }
    if (balance < -1 && AVL_getBalance(root.right) > 0) {
        root.right = AVL_rightRotate(root.left);
        return AVL_leftRotate(root);
    }

    // Return the root node of the tree
    return root;
}

func AVL_printInOrder(node *AVLNode) {
    if node == nil {
        return;
    }

    AVL_printInOrder(node.left);
    fmt.Printf("%d ", node.key);
    AVL_printInOrder(node.right);
}

/**
 * Prints an AVL tree in a diagram format.
 *
 * @param root A pointer to the root node of the tree.
 * @param space The amount of space to print before the current node.
 */
func AVL_printAVLTree(root *AVLNode, space int) {
    // Base case: the tree is empty
    if (root == nil) {
        return;
    }

    // Increase the amount of space for the next level
    space += COUNT;

    // Recursively print the right subtree
    AVL_printAVLTree(root.right, space);

    // Print the current node
    fmt.Printf("\n");
    for i := COUNT; i < space; i++ {
        fmt.Printf(" ");
    }
    fmt.Printf("%d\n", root.key);

    // Recursively print the left subtree
    AVL_printAVLTree(root.left, space);
}
