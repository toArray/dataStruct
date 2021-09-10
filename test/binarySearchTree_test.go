package test

import (
	"fmt"
	"luoqiangDataStruct/binarySearchTree"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := binarySearchTree.InitBinarySearchTree()

	arr := []int32{6, 3, 8, 2, 5, 1, 7, 7}
	for _, value := range arr {
		tree.Insert(value)
	}

	tree.RangeLeft(tree.Root)
	fmt.Println()

	tree.RangeMid(tree.Root)
	fmt.Println()

	tree.RangeRight(tree.Root)
	fmt.Println()

	tree.RangeLayer()
	fmt.Println()

	tree.GetHeight()
}
