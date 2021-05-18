package main

import "github.com/yuanyu90221/go-code-problem-day3/sol"

func main() {
	tree := &sol.TreeNode{
		Val: 9,
		Left: &sol.TreeNode{
			Val: 8,
			Left: &sol.TreeNode{
				Val: 6,
			},
		},
		Right: &sol.TreeNode{
			Val: 7,
			Left: &sol.TreeNode{
				Val: 5,
				Right: &sol.TreeNode{
					Val: 3,
				},
			},
			Right: &sol.TreeNode{
				Val: 4,
			},
		},
	}
	sol.WidthOfBinaryTree(tree)
}
