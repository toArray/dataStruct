package test

import (
	"fmt"
	"testing"
	"time"
)

func TestAvlTree(t *testing.T) {
	time.AfterFunc(time.Hour*12*15, func() {
		fmt.Println(1)
	})
	////arr := []int32{8, 5, 14, 3, 7, 12, 17, 2, 4, 9, 13, 15, 18, 16}
	//arr := []int32{8, 5, 2, 1, 9}
	//tree := avlTree.InitBinarySearchTree()
	//for _, value := range arr {
	//	tree.Insert(value)
	//}
	//
	//avlTree.GetHeight(tree.Root.Left)
	//avlTree.GetHeight(tree.Root.Right)
}
