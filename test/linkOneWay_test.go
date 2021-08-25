package test

import (
	"fmt"
	"luoqiangDataStruct/dataStruct"
	"testing"
)

func TestLink(t *testing.T) {
	linkList := dataStruct.CreatedLink()
	//头插
	linkList.InsertNodeByHead(0)
	linkList.Append(1)
	linkList.Append(2)
	linkList.Append(3)

	//指定位置插入
	err := linkList.InsertNodeByIndex(4, 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	linkList.Println()
	linkList.Delete(5)

	linkList.Println()
	fmt.Println(linkList.Head.Data, linkList.Tail.Data)

	nodeOne := linkList.Search(4)
	fmt.Println(nodeOne)
}
