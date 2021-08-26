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
	linkList.Append(4)

	//指定下标位置插入
	err := linkList.InsertNodeByIndex(0, 9999)
	if err != nil {
		fmt.Println(err)
		return
	}

	linkList.Println()

	//删除
	err = linkList.Delete(4)
	if err != nil {
		fmt.Println(err)
		return
	}

	linkList.Println()

	data := linkList.Search(1)
	fmt.Printf("searchData:%v\n", data)
}
