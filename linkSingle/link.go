package linkSingle

import (
	"errors"
	"fmt"
)

//LinkMethod 单向链表实现接口
type LinkMethod interface {
	InsertNodeByHead(data interface{})                           //插入 - 头插
	InsertNodeByIndex(index int32, data interface{}) (err error) //插入 - 指定下标插入
	Append(data interface{})                                     //插入 - 尾部追加
	Search(index int32) (node *Node, err error)                  //查找
	GetCount() int32                                             //获取个数
	DeleteByIndex(index int32) (err error)                       //删除 - 指定索引删除
	RemoveHead() error                                           //删除 - 删除头
	RemoveTail() error                                           //删除 - 删除尾

	Reverse() //反转链表
}

//LinkList 单向链表列表
type LinkList struct {
	Count int32 //数据个数
	Head  *Node //表头
	Tail  *Node //表尾
}

//Node 链表节点
type Node struct {
	Data interface{} //数据域
	Next *Node       //指针域
}

/*
CreatedLinkList
@Desc 	初始化链表
@Return	*LinkModel
*/
func CreatedLinkList() *LinkList {
	node := CreateNode(0) //创建哨兵节点
	return &LinkList{
		Count: 0,
		Head:  node,
		Tail:  node,
	}
}

/*
CreateNode
@Desc 	创建新节点
@Param 	data int32	节点数据
*/
func CreateNode(data interface{}) *Node {
	return &Node{
		Next: nil,
		Data: data,
	}
}

/*
InsertNodeByIndex
@Desc	下标插入
@Param 	index 	int32	在第几个位置插入
@Param 	data 	int32	插入数据
*/
func (l *LinkList) InsertNodeByIndex(index int32, data interface{}) (err error) {
	if index < 0 || index > l.Count {
		err = errors.New("func InsertNodeByIndex index out of size")
		return
	}

	//头插
	if index == 0 {
		l.InsertNodeByHead(data)
		return
	}

	//尾插
	if index == l.Count {
		l.Append(data)
		return
	}

	//指定索引插入
	curNode, err := l.Search(index - 1)
	if err != nil {
		err = errors.New("search by index out of size")
		return
	}

	//success
	newNode := CreateNode(data)
	newNode.Next = curNode.Next
	curNode.Next = newNode
	l.Count++
	return
}

/*
InsertNodeByHead
@Desc	头插
@Param	data	int32	插入数据
*/
func (l *LinkList) InsertNodeByHead(data interface{}) {
	//更新头节点
	newNode := CreateNode(data)
	newNode.Next = l.Head.Next
	l.Head.Next = newNode
	l.Count++

	//更新尾巴节点
	if newNode.Next == nil {
		l.Tail = newNode
	}
}

/*
Append
@Desc	尾插 - 追加节点
@Param	data	int32	插入数据
*/
func (l *LinkList) Append(data interface{}) {
	newNode := CreateNode(data)
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Count++
}

/*
DeleteByIndex
@Desc	删除指定索引节点
@Param	index	int32	删除第几个节点数据
*/
func (l *LinkList) DeleteByIndex(index int32) (err error) {
	if index < 0 || index >= l.Count {
		return errors.New("func Delete index out of size")
	}

	//找到下个节点
	var preNode *Node
	currentNode := l.Head.Next
	for i := int32(0); i < index; i++ {
		preNode = currentNode
		currentNode = currentNode.Next
	}

	//如果是删除最后一个则要更新tail
	if index == l.Count-1 {
		l.Tail = l.Head
	}

	//删除
	if index == 0 {
		l.Head.Next = currentNode.Next
	} else {
		preNode.Next = currentNode.Next
	}

	l.Count--
	return nil
}

/*
RemoveHead
@Desc	删除头节点
*/
func (l *LinkList) RemoveHead() (err error) {
	//是否空链表
	firstNode := l.Head.Next
	if firstNode == nil {
		err = errors.New("is an empty list for remove head")
		return
	}

	//删除头节点
	secondNode := firstNode.Next
	l.Head.Next = secondNode

	//更新尾节点
	l.Count--
	if l.Count == 0 {
		l.Tail = l.Head
	}

	return
}

/*
RemoveTail
@Desc	删除尾节点
*/
func (l *LinkList) RemoveTail() (err error) {
	//是否空链表
	currentNode := l.Head.Next
	if currentNode == nil {
		err = errors.New("is an empty list for remove tail")
		return
	}

	//找到倒数第二个节点
	var preNode *Node
	for currentNode.Next != nil {
		preNode = currentNode
		currentNode = currentNode.Next
	}

	//更新尾节点
	if preNode != nil {
		preNode.Next = nil
		l.Tail = preNode
	} else {
		l.Head.Next = nil
		l.Tail = l.Head
	}

	return
}

/*
GetCount
@Desc	获得个数
*/
func (l *LinkList) GetCount() int32 {
	return l.Count
}

/*
Search
@Desc	搜索指定位置节点信息
@Param	index int32		指定索引位置
*/
func (l *LinkList) Search(index int32) (node *Node, err error) {
	if index < 0 || index >= l.Count {
		err = errors.New("index out of range for search node")
		return
	}

	node = l.Head.Next
	for i := int32(0); i < index; i++ {
		node = node.Next
	}

	return
}

/*
Println
@Desc: 打印数据
*/
func (l *LinkList) Println() {
	data := l.Head
	for {
		if data == nil {
			break
		}

		fmt.Printf("value：%v | pointer：%p | nextPointer：%p\n", data.Data, data, data.Next)
		data = data.Next
	}
	fmt.Printf("head: %v | tail: %v", l.Head, l.Tail)
	fmt.Println()
}

/*
Reverse
@Desc: 反转链表
*/
func (l *LinkList) Reverse() {
	cur := l.Head.Next
	head := l.Head.Next
	var prev *Node
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	l.Head.Next = prev
	l.Tail = head
}
