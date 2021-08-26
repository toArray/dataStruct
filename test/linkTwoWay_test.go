package test

import (
	"fmt"
	"luoqiangDataStruct/dataStruct"
	"testing"
)

func TestTwoWayLink(t *testing.T) {
	linkList := dataStruct.CreateLinkTwoWay()
	linkList.Append(0)
	linkList.Append(1)
	linkList.Append(2)
	linkList.InsertNodeByHead(2)
	linkList.InsertNodeByHead(1)
	linkList.InsertNodeByHead(0)

	node, err := linkList.Search(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(node)

	err = linkList.InsertNodeByIndex(2, 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = linkList.Delete(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	linkList.Println()
}
