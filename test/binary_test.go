package test

import (
	"fmt"
	"luoqiangDataStruct/binary"
	"testing"
)

func TestA(t *testing.T) {
	tree := new(binary.Tree)
	arr := []int32{6, 4, 56, 7, 81, 1, 2, 3}
	for _, value := range arr {
		tree.InsertNode(value)
	}

	binary.Show(tree.Root)
	fmt.Println(binary.RightMin(tree.Root))
}
