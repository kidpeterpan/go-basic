package main

import "fmt"

type BinaryTree struct {
	value int
	left  *BinaryTree // *** ถ้าจะ ref หาตัวเองต้อง ref หา address (pointer) ***
	right *BinaryTree
}

func main() {
	root := BinaryTree{value: 2}
	left := BinaryTree{value: 1}
	right := BinaryTree{value: 3}

	root.left = &left
	root.right = &right

	fmt.Printf("%+v\n\n", root) // {value:1 left:0xc00011a000 right:0xc00011a020}

	// show depth first
	showDepthFirst(&root)
}

func showDepthFirst(node *BinaryTree) {
	if node != nil {
		showDepthFirst(node.left)
		fmt.Println(node.value)
		showDepthFirst(node.right)
	}
}
