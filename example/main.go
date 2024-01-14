package main

func main() {
	arr := []int{7, 4, 9, 2, 5, 8, 11, 3, 12, 1}
	tree := binaryTree.NewBinaryTree()
	for _, v := range arr {
		tree.Add(v)
	}
	tree.TreePrint()
	tree.Remove(9)
	//tree.TreePrint()
}
