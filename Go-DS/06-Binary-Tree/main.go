package main

import (
	"fmt"
	"math"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func NewTree(key int) *Node {
	return &Node{Key: key}
}

//Insert
func (n *Node) Insert(key int) {
	if key > n.Key {
		//go right
		if n.Right == nil {
			n.Right = NewTree(key)
		} else {
			n.Right.Insert(key)
		}
	} else if key < n.Key {
		//go left
		if n.Left == nil {
			n.Left = NewTree(key)
		} else {
			n.Left.Insert(key)
		}
	}

}

//Search
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if key > n.Key {
		return n.Right.Search(key)
	} else if key < n.Key {
		return n.Left.Search(key)
	}

	return true
}

//Print InOrder
func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Println(n.Key)
	n.Right.InOrder()
}

//Mirror Tree
func (n *Node) Mirror() {
	if n == nil {
		return
	}
	n.Left.Mirror()
	n.Right.Mirror()

	n.Left, n.Right = n.Right, n.Left
}

//validate search
var prev *Node

func (n *Node) isValid() bool {
	if n != nil {
		if !n.Left.isValid() {
			return false
		}

		if prev != nil && n.Key <= prev.Key {
			return false
		}
		prev = n
		if !n.Right.isValid() {
			return false
		}

	}
	return true
}

func (n *Node) checkValid(x, y int) bool {
	if n == nil {
		return true
	}

	if n.Key < x || n.Key > y {
		return false
	}
	return (n.Left.checkValid(x, n.Key-1) && n.Right.checkValid(n.Key+1, y))
}

func main() {
	tree := NewTree(4)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)
	tree.InOrder()
	fmt.Println(tree.isValid())
	fmt.Println(tree.checkValid(math.MinInt, math.MaxInt))
	tree.Mirror()
	fmt.Println("After Mirror")
	fmt.Println(tree.checkValid(math.MinInt, math.MaxInt))
	tree.InOrder()
	fmt.Println(tree.isValid())
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(1))
}
