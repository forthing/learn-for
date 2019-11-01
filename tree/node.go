package tree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func creatNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) Setvalue(value int) {
	node.value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}


//	fmt.Println(root)
//	fmt.Println(*root.left)
//	fmt.Println(*root.right)

//	nodes := []treeNode {
//		{value: 3},
//		{},
//		{6, nil, &root},
//	}
//	fmt.Println(nodes)
/*
	root.print()
	fmt.Println()

	root.right.left.setvalue(4)
	root.right.left.print()
	fmt.Println()

	root.print()
	fmt.Println()
	root.setvalue(100)
	root.print()

	pRoot := &root
*/

