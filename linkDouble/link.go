package linkDouble

import (
	"errors"
	"fmt"
)

//LinkMethod 双向链表实现接口
type LinkMethod interface {
	InsertNodeByHead(data interface{})                           //插入 - 头插
	InsertNodeByIndex(index int32, data interface{}) (err error) //插入 - 指定下标插入
	Append(data interface{})                                     //插入 - 尾部追加
	Search(index int32) (node *Node, err error)                  //查找
	GetCount() int32                                             //获取个数
	RemoveHead() error                                           //删除 - 移除头节点
	RemoveTail() error                                           //删除 - 移除尾节点
	RemoveByIndex(index int32) error                             //删除 - 通过索引移除节点
}

//LinkList 双向链表列表
type LinkList struct {
	Count int32 //数据个数
	Head  *Node //表头
	Tail  *Node //表尾
}

//Node 双向链表节点
type Node struct {
	Data interface{} //数据域
	Next *Node       //下个指针域
	Prev *Node       //上个指针域
}

//CreateLinkList 创建双向链表
func CreateLinkList() *LinkList {
	node := CreateNodeTwoWay(0) //创建哨兵节点
	return &LinkList{
		Count: 0,
		Head:  node,
		Tail:  node,
	}
}

/*
CreateNodeTwoWay
@Desc	创建新节点
@Param	data	interface{}	节点数据
*/
func CreateNodeTwoWay(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}

/*
InsertNodeByHead
@Desc	头插
@Param	data	interface{}	插入数据
*/
func (l *LinkList) InsertNodeByHead(data interface{}) {
	//新节点
	firstNode := l.Head.Next
	newNode := CreateNodeTwoWay(data)
	newNode.Next = firstNode
	newNode.Prev = l.Head
	l.Head.Next = newNode

	//更新旧节点指向和尾巴节点
	if firstNode != nil {
		firstNode.Prev = newNode
	} else {
		l.Tail = newNode
	}

	l.Count++
}

/*
InsertNodeByIndex
@Desc 	指定索引位置插入
@Param 	index	int32		索引位置
@Param 	data	interface{}	插入数据
*/
func (l *LinkList) InsertNodeByIndex(index int32, data interface{}) (err error) {
	if index < 0 || index > l.Count {
		err = errors.New("index out of range for insert node")
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

	//找当前索引节点数据
	curNode, err := l.Search(index)
	if err != nil {
		return
	}

	//新节点
	newNode := CreateNodeTwoWay(data)
	newNode.Next = curNode
	newNode.Prev = curNode.Prev

	//赋值
	curNode.Prev.Next = newNode
	curNode.Prev = newNode
	l.Count++
	return
}

/*
Append
@Desc 	尾插 - 追加节点
@Param 	data	interface{}	插入数据
*/
func (l *LinkList) Append(data interface{}) {
	newNode := CreateNodeTwoWay(data)
	newNode.Prev = l.Tail
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Count++
}

/*
Search
@Desc 	搜索信息
@Param	index int32	索引位置
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
RemoveHead
@Desc	删除头节点
*/
func (l *LinkList) RemoveHead() (err error) {
	//验证是否空链表
	firstNode := l.Head.Next
	if firstNode == nil {
		err = errors.New("is an empty list for remove head")
		return
	}

	//更新第二节点指向
	secondNode := firstNode.Next
	l.Head.Next = secondNode
	if secondNode != nil {
		secondNode.Prev = l.Head
	}

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
	//验证是否空链表
	prevNode := l.Tail.Prev
	if prevNode == nil {
		err = errors.New("is an empty list for remove tail")
		return
	}

	//更新尾巴
	prevNode.Next = nil
	l.Tail = prevNode
	l.Count--
	return
}

/*
RemoveByIndex
@Desc	删除指定索引节点
@Param	index	int32	指定节点索引
*/
func (l *LinkList) RemoveByIndex(index int32) (err error) {
	//获得当前索引数据
	curNode, err := l.Search(index)
	if err != nil {
		err = errors.New("index out of range for remove index")
		return
	}

	//更新节点指向
	nextNode := curNode.Next
	prevNode := curNode.Prev
	prevNode.Next = nextNode
	if nextNode != nil {
		nextNode.Prev = prevNode
	}

	//更新尾节
	if index == l.Count-1 {
		l.Tail = l.Head
	}

	l.Count--
	return
}

/*
GetCount
@Desc 获得个数
*/
func (l *LinkList) GetCount() int32 {
	return l.Count
}

/*
Println
@Desc	打印数据
*/
func (l *LinkList) Println() {
	data := l.Head
	for {
		if data == nil {
			break
		}
		fmt.Printf("value：%v | pointer：%p | prev：%p | next：%p\n", data.Data, data, data.Prev, data.Next)
		data = data.Next
	}
	fmt.Println()
}
