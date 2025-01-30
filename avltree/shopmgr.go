package avltree


type UserNode struct {
	start int      // start time when user enters
    end   int      // end time when user leaves
	height int     // height of the tree
    max   int      // maximum end time rooted by this user
	left *UserNode
	right *UserNode
    data  interface{} // other data of user like userid, emailid, 
                      //mobile number, we can maintain another map or avl tree based on other key
}

/* returns new node of AVLTree
*/
func newUser(start, end int) *UserNode{
	return &UserNode{
		start: start,
        end: end,
		height: 0,
        max: end,
		left: nil,
		right: nil,
        data: nil,
	}
}
// with recursive call no need of the member functions

/* returns height*/
func getHeight(node *UserNode) int {
	if node == nil {
		return 0;
	}
	return node.height
}

/* get balance factor */
func getBalance(node *UserNode) int {
	if node == nil {
		return 0;
	}
	return (getHeight(node.left) - getHeight(node.right))
}

/* returns max interval of users rooted by this guy*/
func getMax(node *UserNode) int {
    if node == nil {
        return 0;
    }
    return node.max
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

func leftRotate(x *UserNode) *UserNode {
    y := x.right;
    if y==nil {
        return x; // not possible to left rotate
    }
    T2 := y.left;

    // Perform the rotation
    y.left = x;
    x.right = T2;

    // Update the heights of the nodes
    x.height = 1 + max(getHeight(x.left), getHeight(x.right));
    y.height = 1 + max(getHeight(y.left), getHeight(y.right));
    // Update the max interval of users rooted by these guys
    x.max = max(getMax(x.left), getMax(x.right));
    y.max = max(getMax(y.left), getMax(y.right));
    

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
func rightRotate(y *UserNode) *UserNode {
    x := y.left
    if x==nil {
        return y; // not possible
    }
    T2 := x.right

    // perform the rotation
    x.right = y 
    y.left = T2

    // update the heights of x and y
    y.height = 1+max(getHeight(y.left), getHeight(y.right))
    x.height = 1+max(getHeight(x.left), getHeight(x.right))
    // Update the max interval of users rooted by these guys
    y.max = max(getMax(y.left), getMax(y.right));
    x.max = max(getMax(x.left), getMax(x.right));

    
    

    // return new root
    return x
}

/* 
 * try to search the key in avl tree,
 * search inleft subtree if key is < than the node key, else search in right subtree
 */
func CountUsers(node *UserNode, timestamp int) int {
    // If the tree is empty or the key is not found, return nil
    if node == nil {
        return 0;
    }

    // first find the first node which is greater than this timestamp
    if (timestamp < node.start) {
        return CountUsers(node.left, timestamp);
    } else {
    	// now count all the nodes in this subtree which satisfies, start<=timestamp<=end
        if node.start<=timestamp && node.end>=timestamp {
            return 1 + CountUsers(node.left, timestamp) + CountUsers(node.right, timestamp)
        } else {
            return CountUsers(node.left, timestamp) + CountUsers(node.right, timestamp)
        }
    }
}


/**
 * Finds the node with the minimum key in a given subtree.
 *
 * @param node A pointer to the root node of the subtree.
 * @return A pointer to the node with the minimum key.
 */
func minValueNode(node *UserNode) *UserNode {
    for node.left != nil {
        node = node.left
    }
    return node
}

/**
 * Inserts a USer into an AVL tree keyed on start timestamp
 *
 * @param node A pointer to the root node of the AVL tree.
 * @param key The key to be inserted.
 * @return A pointer to the root node of the AVL tree after the insertion.
 */
func InsertUser(node *UserNode, start, end int) *UserNode {
    // performs normal BST insert 
    if node == nil {
        return newUser(start, end);
    }

    // first insert then balance the tree, left == 
    if (start <= node.start) {
    	// insert recursively in left subtree
        node.left = InsertUser(node.left, start, end);
    } else if (start > node.start) {
        node.right = InsertUser(node.right, start, end);
    }

    // update balance parameters, height, count and max while backtracking
    node.height = 1 + max(getHeight(node.left), getHeight(node.right));
    node.max = max(getMax(node.left), getMax(node.right));
    

    // Calculate the balance factor of the root node
    balance := getBalance(node);

    // If the balance factor is greater than 1 and the key to be inserted is less than the key in the left child,
    // perform a right rotation on the root node
    /*
        LL rotation: tree becomes unbalance after insertion in left subtree of left child of root node. so name is LL rotation.
    */
    if (balance > 1 && start < node.left.start) {
        return rightRotate(node);
    }

    // If the balance factor is less than -1 and the key to be inserted is greater than the key in the right child,
    // perform a left rotation on the root node
    /*
        RR rotation: tree becomes unbalanced by insertion in to right subtree of right child of the root node.
    */
    if (balance < -1 && start > node.right.start) {
        return leftRotate(node);
    }

    // If the balance factor is greater than 1 and the key to be inserted is greater than the key in the left child,
    // perform a left rotation on the left child first and then a right rotation on the root node
    /*
        LR Rotation : there was insertion in direction root->left->right
        steps first make LL-> by doing left rotation at child, then right rotation at root
        leftrotate(child) -> right rotate(root of subtree)
    */
    if (balance > 1 && start > node.left.start) {
        node.left = leftRotate(node.left);
        return rightRotate(node);
    }

    // If the balance factor is less than -1 and the key to be inserted is less than the key in the right child,
    // perform a right rotation on the right child first and then a left rotation on the root node
    /*
        RL rotation: sequence insertion in R(root) and left(child)
        right roatation(child) -> left rotation(root)
        i.e. first make RR and then do left rotation.
    */
    if (balance < -1 && start < node.right.start) {
        node.right = rightRotate(node.right);
        return leftRotate(node);
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
func DeleteUser(root *UserNode, start, end int) *UserNode {
    // perform bst deletion
    if (root == nil) {
        return root;
    }

    // Recursively search for the node to be deleted
    if (start < root.start) {
        root.left = DeleteUser(root.left, start, end);
    } else if (start > root.start) {
        root.right = DeleteUser(root.right, start, end);
    } else {
        if (start == root.start && end != root.end) {
            // not found yet , search in left subtree only
            root.left = DeleteUser(root.left, start, end);
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
            successor := minValueNode(root.right);
            root.start = successor.start;
            root.end = successor.end;
            root.right = DeleteUser(root.right, successor.start, successor.end);
        }
    }

    // Update the height of the current node
    root.height = 1 + max(getHeight(root.left), getHeight(root.right))
    root.max = max(getMax(root.left), getMax(root.right));
    

    // Check if the tree is unbalanced and perform the necessary rotations
    balance := getBalance(root);
    if (balance > 1 && getBalance(root.left) >= 0) {
        return rightRotate(root);
    }
    if (balance > 1 && getBalance(root.left) < 0) {
        root.left = leftRotate(root.left);
        return rightRotate(root);    
    }
    if (balance < -1 && getBalance(root.right) <= 0) {
        return leftRotate(root);
    }
    if (balance < -1 && getBalance(root.right) > 0) {
        root.right = rightRotate(root.left);
        return leftRotate(root);
    }

    // Return the root node of the tree
    return root;
}