# binaryTree
go语言实现的二叉树，avl树和红黑树学习例子

example:

```go
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
```

结果：  

    ===================================================================
                     7
               ┌─────┴─────┐
               4           9
         ┌─────┴──┐     ┌──┴───┐
         2        5     8     11
      ┌──┴──┐                  └──┐     
      1     3                    12    
    ===================================================================
                     7
               ┌─────┴──────┐
               4           11
         ┌─────┴──┐     ┌───┴───┐
         2        5     8      12          
      ┌──┴──┐                            
      1     3                           
    
                      
我相信最难的可能不只是如何轻松实现红黑树，还有如何实现上面树状打印结果，后面会出一篇文章专门讲解

你可以用任何语言轻松实现树状打印。方法为自己研究，之前找了很久网上的文章，想打印一颗二叉树，观察代码是否正确，找了很久都没找到，最后还是自己去研究，最后研究出来这个方法，纯套公式，可以控制打印结果，兼容汉字。有问题的朋友评论区留言，抽空解答。

关于红黑树的研究，网上有很多最全教程都有问题，结果就是困扰自己很久，后面通过断点调试，慢慢还原了整个场景，总结出了整套方法
