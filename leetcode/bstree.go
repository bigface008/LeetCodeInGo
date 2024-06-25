package leetcode

import (
	"errors"
	"fmt"
)

type BinarySearchTree struct {
	root *BSTNode
	size int
}

type BSTNode struct {
	key    int
	val    interface{}
	left   *BSTNode
	right  *BSTNode
	parent *BSTNode
}

func NewBSTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) Set(key int, val interface{}) {
	node := &BSTNode{key, val, nil, nil, nil}
	if bst.root == nil {
		bst.root = node
		bst.size++
		return
	}

	ptr := bst.root
	for {
		temp := ptr
		if key == ptr.key {
			ptr.val = val
			return
		} else if key > ptr.key {
			ptr = ptr.right
			if ptr == nil {
				temp.right = node
				node.parent = temp
				bst.size++
				return
			}
		} else {
			ptr = ptr.left
			if ptr == nil {
				temp.left = node
				node.parent = temp
				bst.size++
				return
			}
		}
	}
}

func (bst *BinarySearchTree) Find(key int) interface{} {
	ptr := bst.root
	for ptr != nil {
		if key == ptr.key {
			return ptr.val
		} else if key > ptr.key {
			ptr = ptr.right
		} else {
			ptr = ptr.left
		}
	}
	return nil
}

func (bst *BinarySearchTree) Remove(key int) error {
	if bst.size == 0 {
		return errors.New("tree is empty")
	}
	ptr := bst.root
	for ptr != nil {
		if key == ptr.key {
			break
		} else if key > ptr.key {
			ptr = ptr.right
		} else {
			ptr = ptr.left
		}
	}
	if ptr == nil {
		return errors.New(fmt.Sprintf("%v not found", key))
	}

	newChild := func() *BSTNode {
		if ptr.right == nil {
			if ptr.left == nil {
				return nil
			} else {
				return ptr.left
			}
		} else {
			if ptr.left == nil {
				return ptr.right
			} else {
				temp := ptr.right
				for temp.right != nil {
					temp = temp.right
				}
				temp.right = ptr.left
				return ptr.right
			}
		}
	}()
	if ptr.parent == nil {
		bst.root = newChild
		if newChild != nil {
			newChild.parent = nil
		}
	} else {
		if ptr.parent.left == ptr {
			ptr.parent.left = newChild
		} else {
			ptr.parent.right = newChild
		}
		if newChild != nil {
			newChild.parent = ptr.parent
		}
	}
	bst.size--
	return nil
}
