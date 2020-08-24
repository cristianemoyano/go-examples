package main

import (
	"errors"
	"fmt"
)

//TreeNode data structure represents a typical binary tree
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//PrintInorder prints the elements in order
func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}

	t.left.PrintInorder()
	fmt.Printf("%v ", t.val)
	t.right.PrintInorder()
}

//AppendListInOrder prints the elements in order
func (t *TreeNode) AppendListInOrder(s *[]int) {
	if t == nil {
		return
	}

	t.left.AppendListInOrder(s)
	*s = append(*s, t.val)
	t.right.AppendListInOrder(s)
}

//Print prints the elements without sorting
func (t *TreeNode) Print() {
	if t == nil {
		return
	}
	fmt.Printf("%v \n", t.val)
	t.left.Print()
	t.right.Print()

}

//Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) Insert(value int) error {

	if t == nil {

		return errors.New("Tree is nil")
	}

	if t.val == value {

		return errors.New("This node value already exists")
	}

	if t.val > value {

		if t.left == nil {

			t.left = &TreeNode{val: value}
			return nil
		}

		return t.left.Insert(value)
	}

	if t.val < value {

		if t.right == nil {

			t.right = &TreeNode{val: value}
			return nil
		}

		return t.right.Insert(value)
	}

	return nil
}

//Insert inserts a new node into the binary tree in order.
func (t *TreeNode) InsertInOrder(value int) error {

	if t == nil {
		return errors.New("Tree is nil")
	} else if value <= t.val {
		if t.left == nil {
			t.left = &TreeNode{val: value, left: nil, right: nil}
		} else {
			t.left.InsertInOrder(value)
		}
	} else {
		if t.right == nil {
			t.right = &TreeNode{val: value, left: nil, right: nil}
		} else {
			t.right.InsertInOrder(value)
		}
	}
	return nil
}

//Find finds the treenode for the given node val
func (t *TreeNode) Find(value int) bool {

	if t == nil {
		return false
	}

	switch {
	case value == t.val:
		return true
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

//Delete removes the Item with value from the tree
func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (t *TreeNode) remove(value int) *TreeNode {

	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}
	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

//FindMax finds the max element in the given BST
func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

//FindMin finds the min element in the given BST
func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//AppendListInOrder prints the elements in order
func InsertByList(list []int) *TreeNode {
	if list == nil {
		return &TreeNode{}
	}
	tree := &TreeNode{}

	for _, value := range list {
		tree.InsertInOrder(value)
	}
	return tree
}

func main() {
	var slice []int

	tree := &TreeNode{val: 454}
	tree.InsertInOrder(100)
	tree.InsertInOrder(20)
	tree.InsertInOrder(200)
	tree.InsertInOrder(70)
	tree.InsertInOrder(550)
	tree.InsertInOrder(80)

	// if t.Find(80) {
	// 	fmt.Println("Se encontro valor")
	// }

	// t.Delete(70)

	// t.PrintInorder()

	// fmt.Println("")

	tree.AppendListInOrder(&slice)
	fmt.Println("")
	printSlice(slice)
	fmt.Println("")

	newtree := InsertByList(slice)
	newtree.PrintInorder()
	fmt.Println("")

	fmt.Println("min is", tree.FindMin())
	fmt.Println("max is", tree.FindMax())
}
