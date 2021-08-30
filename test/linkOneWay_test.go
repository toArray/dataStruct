package test

import (
	"fmt"
	"luoqiangDataStruct/dataStruct"
	"testing"
)

func TestLink(t *testing.T) {
	linkList := dataStruct.CreatedLink()
	//linkList.InsertNodeByHead(1)
	//linkList.InsertNodeByHead(2)
	//linkList.InsertNodeByHead(3)
	//linkList.Append(1)
	//linkList.Append(2)
	//linkList.Append(3)
	//linkList.InsertNodeByHead(3)
	linkList.InsertNodeByIndex(0, 1)
	linkList.InsertNodeByIndex(0, 2)
	linkList.InsertNodeByIndex(2, 3)
	node, err := linkList.Search(2)

	//err := linkList.Delete(2)
	fmt.Println(err, node)
	linkList.Println()

	////头插
	//linkList.InsertNodeByHead(0)
	//linkList.Append(1)
	//linkList.Append(2)
	//linkList.Append(3)
	//linkList.Append(4)
	//
	////指定下标位置插入
	//err := linkList.InsertNodeByIndex(0, 9999)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//linkList.Println()
	//
	////删除
	//err = linkList.Delete(4)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//linkList.Println()
	//
	//data, err := linkList.Search(1)
	//fmt.Printf("searchData:%v\n", data)
}
