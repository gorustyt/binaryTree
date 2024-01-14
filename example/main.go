package main

import (
	"fmt"
	bt "github.com/gorustyt/binaryTree"
)

func main() {
	arr := []int{7, 4, 9, 2, 5, 8, 11, 3, 12, 1}

	tree := bt.NewBinaryTree[int]()
	for _, v := range arr {
		tree.Add(v)
	}
	bt.PrintBinaryTree[int](tree.Root, func(node *bt.BinaryTreeNode[int]) string {
		return fmt.Sprintf("%v", node.GetElement()) //自定义显示值的方法
	})
	tree.Remove(9)
	bt.PrintBinaryTree[int](tree.Root, func(node *bt.BinaryTreeNode[int]) string {
		return fmt.Sprintf("%v", node.GetElement()) //自定义显示值的方法
	})
}
