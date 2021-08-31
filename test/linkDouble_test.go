package test

import (
	"luoqiangDataStruct/linkDouble"
	"testing"
)

func TestLinkDouble(t *testing.T) {
	linkList := linkDouble.CreateLinkList()

	//linkList.InsertNodeByHead(3)
	//linkList.InsertNodeByHead(2)
	//linkList.InsertNodeByHead(1)
	linkList.Append(0)
	//linkList.Append(-1)
	//linkList.Append(-2)
	linkList.Println()
	//linkList.Append(2)
	//linkList.Append(3)
	//linkList.InsertNodeByHead(4)
	//linkList.RemoveByIndex(0)
	//linkList.RemoveByIndex(0)
	//linkList.RemoveByIndex(0)
	//linkList.RemoveTail()
	//linkList.RemoveTail()
	//linkList.RemoveHead()
	//linkList.Append(3)
	//linkList.Append(4)
	//linkList.Delete(0)
	//linkList.Delete(0)
	//linkList.Delete(0)
	//linkList.Delete(0)
	//fmt.Println(linkList)

	//linkList.Append(0)
	//linkList.Append(1)
	//linkList.Append(2)
	//linkList.InsertNodeByHead(2)
	//linkList.InsertNodeByHead(1)
	//linkList.InsertNodeByHead(0)
	//
	//node, err := linkList.Search(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(node)
	//
	//err = linkList.InsertNodeByIndex(2, 4)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//err = linkList.Delete(2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//linkList.Println()
}
