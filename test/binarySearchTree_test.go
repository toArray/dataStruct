package test

import (
	"fmt"
	"luoqiangDataStruct/binarySearchTree"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	//arr := []int32{8, 5, 14, 3, 7, 12, 17, 2, 4, 9, 13, 15, 18, 16}
	arr := []int32{8, 5}
	tree := binarySearchTree.InitBinarySearchTree()
	for _, value := range arr {
		tree.Insert(value)
	}

	//左
	fmt.Println("左序遍历")
	tree.RangeLeft(tree.Root)
	fmt.Println()

	//中
	fmt.Println("中序遍历")
	tree.RangeMid(tree.Root)
	fmt.Println()

	//右
	fmt.Println("右序遍历")
	tree.RangeRight(tree.Root)
	fmt.Println()

	//层级
	fmt.Println("层级遍历")
	tree.RangeLayer()
	fmt.Println()
}
