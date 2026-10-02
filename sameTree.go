package main

import (
	"fmt"
	"math"
	"slices"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func markTree(node *TreeNode, marker *[]int) {
	if node == nil {
		*marker = append(*marker, math.MinInt)
	} else {
		*marker = append(*marker, node.Val)
		markTree(node.Left, marker)
		markTree(node.Right, marker)
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var markerP, markerQ []int
	markTree(p, &markerP)
	markTree(q, &markerQ)

	return slices.Equal(markerP, markerQ)
}

func main() {
	leftNodeP := TreeNode{
		Val: 2,
	}

	rootNodeP := TreeNode{
		Val:  1,
		Left: &leftNodeP,
	}

	rightNodeQ := TreeNode{
		Val: 2,
	}

	rootNodeQ := TreeNode{
		Val:   1,
		Right: &rightNodeQ,
	}

	fmt.Println("main result", isSameTree(&rootNodeP, &rootNodeQ))
}
