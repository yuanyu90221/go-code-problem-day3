package sol

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFT(node *TreeNode) int {
	// step 1 建立 queue for target visit node
	queue := append([]*TreeNode{}, node)
	levelQueue := append([]int{}, 0)
	levelCount := 0
	prevLevel := 0
	resultCount := 0
	// step 2
	for len(queue) != 0 {
		visit := queue[0]
		visitLevel := levelQueue[0]

		queue = append(queue[1:])
		levelQueue = append(levelQueue[1:])
		if visitLevel == prevLevel {
			levelCount += 1
		} else {
			levelCount = 0
			levelCount += 1
		}
		if visit != nil {
			if visit.Left != nil || visit.Right != nil {
				queue = append(queue, visit.Left)
				queue = append(queue, visit.Right)
				levelQueue = append(levelQueue, visitLevel+1)
				levelQueue = append(levelQueue, visitLevel+1)
			}
		}
		if levelCount > resultCount {
			resultCount = levelCount
		}
		prevLevel = visitLevel
	}
	fmt.Println("resultCount", resultCount)
	return resultCount
}

func WidthOfBinaryTree(node *TreeNode) int {
	return BFT(node)
}
