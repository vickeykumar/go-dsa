package dsf

/*
 * disjoint set forest Data setructure
 * Applications: 
 * 1. merge accounts with the same email id - email id can be the key for this data structure
 *          - https://leetcode.com/problems/accounts-merge/description/
 * 2.
 *
 */

// a node in the disjoint-set forest
type node struct {
  data interface{} //payload
  rank int 			// height of the node in dsf
  parent *node      // parent
}


// create a new set with element x, x as parent
func MakeSet(data interface{}) *node {
	var n node
	n.data = data
	n.rank = 0
	n.parent = &n
	return &n
}

// finds the root of the set that element x belongs to
func FindSet(x *node) *node {
	for x.parent != x {
		x=x.parent
	}
	return x
}

/* merges the set containing element x with the set containing element y
 * node with more height becomes parent to balance the height/rank
 */
func UnionSet(x, y *node) (root *node) {
	root1 := FindSet(x)
	root2 := FindSet(y)
	if root1!=root2 {
		if root1.rank < root2.rank {
			root1.parent=root2
			root = root2
		} else if root1.rank > root2.rank {
			root2.parent=root1
			root = root1
		} else {
			// both are eq, make parent to any and increase height of the other guy
			root2.parent = root1
			root1.rank++
			root = root1
		}
	}
	return
}
