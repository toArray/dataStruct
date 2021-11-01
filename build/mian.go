package main

import (
	"fmt"
	"luoqiangDataStruct/binarySearchTree"
)

var (
	FirstName   string
	SecondNames int32
)

func main() {
	arr := []int32{8, 5, 14, 3, 7, 12, 17, 2, 4, 9, 13, 15, 18, 16}
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

	binarySearchTree.Show(tree.Root)
	for {
		fmt.Printf("请输入指令[add] [del]: ")
		fmt.Scanln(&FirstName, &SecondNames) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
		switch FirstName {
		case "add":
			tree.Insert(SecondNames)
			binarySearchTree.Show(tree.Root)
		case "del":
			tree.Delete(SecondNames)
			binarySearchTree.Show(tree.Root)
		default:
			fmt.Println("指令错误。。。")
		}
	}
}
