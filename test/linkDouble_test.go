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
	linkList.Append(1)
	linkList.Append(2)
	linkList.Append(3)
	linkList.RemoveTail()
	linkList.RemoveHead()
	linkList.RemoveTail()
	//linkList.RemoveTail()
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	tail := res

	cur1 := l1
	cur2 := l2
	for {
		//退出条件
		if cur1 == nil && cur2 == nil {
			break
		}

		if cur1 != nil {
			if cur2 == nil {
				tail.Next = cur1
				cur1 = nil
			} else {
				if cur1.Val >= cur2.Val {
					tail.Next = &ListNode{Val: cur2.Val}
					tail = tail.Next
					cur2 = cur2.Next
				}
			}
		}

		if cur2 != nil {
			if cur1 == nil {
				tail.Next = cur2
				cur2 = nil
			} else {
				if cur2.Val >= cur1.Val {
					tail.Next = &ListNode{Val: cur1.Val}
					tail = tail.Next
					cur1 = cur1.Next
				}
			}
		}
	}

	return res
}
