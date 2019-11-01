package main

type TreeNode 

func main() {
	var root TreeNode

	root := TreeNode{value: 3}
	root.Left = &treeNode{}
	root.Right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = creatNode(2)

	root.traverse()
}
