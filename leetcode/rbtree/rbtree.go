package rbtree

const (
	RED   = true
	BLACK = false
)

type RBTree struct {
	NIL   *RBTNode
	root  *RBTNode
	count uint64
}

type RBTNode struct {
	parent *RBTNode
	left   *RBTNode
	right  *RBTNode
	color  bool
	item   int
}

func New() *RBTree {
	node := RBTNode{nil, nil, nil, BLACK, 0}
	return &RBTree{
		NIL:   &node,
		root:  &node,
		count: 0,
	}
}

func (rbt *RBTree) Insert(item int) {
	n := &RBTNode{rbt.NIL, rbt.NIL, rbt.NIL, RED, item}

	x := rbt.root
	y := rbt.NIL
	for x != rbt.NIL {
		y = x
		if item < x.item {
			x = x.left
		} else if item > x.item {
			x = x.right
		} else {
			return
		}
	}
	n.parent = y
	if y == rbt.NIL {
		rbt.root = n
	} else if n.item < y.item {
		y.left = n
	} else {
		y.right = n
	}
	rbt.count++
	rbt.insertFixup(n)
}

func (rbt *RBTree) insertFixup(n *RBTNode) {
	for n.parent.color == RED {
		if n.parent == n.parent.parent.left { // n's parent is left child
			y := n.parent.parent.right
			if y.color == RED {
				// Case 4
				n.parent.color = BLACK
				y.color = BLACK
				n.parent.parent.color = RED
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					// Case 5
					if n == n.parent.right {
						n = n.parent
						rbt.leftRotate(n)
					}
				}
				n.parent.color = BLACK
				n.parent.parent.color = RED
				rbt.rightRotate(n)
			}
		} else { // n's parent is right child
			y := n.parent.parent.left
			if y.color == RED {
				// Case 4
				n.parent.color = BLACK
				y.color = BLACK
				n.parent.parent.color = RED
			} else {

			}
		}
	}
	rbt.root.color = BLACK
}

// Left Rotate
//
//	  |                                  |
//	  X                                  Y
//	 / \         left rotate            / \
//	α   Y       ------------->         X   γ
//	   / \                            / \
//	  β   γ                          α   β
func (rbt *RBTree) leftRotate(n *RBTNode) {
	if n.right == nil {
		return
	}

	rchild := n.right
	n.right = rchild.left
	if rchild.left != nil {
		rchild.left.parent = n
	}
	rchild.parent = n.parent
	if n.parent != nil {
		rbt.root = rchild
	} else if n == n.parent.left {
		n.parent.left = rchild
	} else {
		n.parent.right = rchild
	}
	rchild.left = n
	n.parent = rchild
}

// Right Rotate
//
//	    |                                  |
//	    X                                  Y
//	   / \         right rotate           / \
//	  Y   γ       -------------->        α   X
//	 / \                                    / \
//	α   β                                  β   γ
func (rbt *RBTree) rightRotate(n *RBTNode) {
	if n.left == nil {
		return
	}

	lchild := n.left
	n.left = lchild.left
	if lchild.right != nil {
		lchild.right.parent = n
	}
	lchild.parent = n.parent
	if n.parent != nil {
		rbt.root = lchild
	} else if n == n.parent.left {
		n.parent.left = lchild
	} else {
		n.parent.right = lchild
	}
	lchild.right = n
	n.parent = lchild
}
