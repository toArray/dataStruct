package test

import (
	"luoqiangDataStruct/linkSingle"
	"testing"
)

func TestLink(t *testing.T) {
	linkList := linkSingle.CreatedLinkList()
	linkList.Append(1)
	linkList.Append(2)
	linkList.Append(3)
	linkList.Println()

	linkList.Reverse()
	linkList.Println()
	//
	//linkList.InsertNodeByIndex(0, 1)
	//linkList.InsertNodeByIndex(1, 2)
	//linkList.Println()

	//linkList.InsertNodeByHead(1)
	//linkList.InsertNodeByHead(2)
	//linkList.InsertNodeByHead(3)
	//linkList.Append(1)
	//linkList.Append(2)
	//linkList.Append(3)
	//linkList.InsertNodeByHead(3)
	//linkList.InsertNodeByIndex(0, 1)
	//linkList.InsertNodeByIndex(0, 2)
	//linkList.InsertNodeByIndex(2, 3)
	//node, err := linkList.Search(2)
	//linkList.RemoveHead()
	//linkList.RemoveHead()
	//linkList.RemoveHead()
	//err := linkList.RemoveHead()
	//err = linkList.RemoveHead()
	//err = linkList.RemoveHead()
	//linkList.Delete(0)
	//linkList.Delete(0)
	//linkList.Delete(0)
	//linkList.RemoveTail()
	//linkList.RemoveTail()
	//linkList.RemoveTail()
	//linkList.RemoveTail()
	//linkList.Append(1)
	//linkList.Append(1)
	//err = linkList.RemoveHead()
	//fmt.Println(err)
	//err := linkList.Delete(2)
	//fmt.Println(err, node)
	//linkList.Println()

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
