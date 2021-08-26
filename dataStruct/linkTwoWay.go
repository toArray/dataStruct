package dataStruct

import (
	"errors"
	"fmt"
)

//LinkTwoWayMethod 双向链表实现接口
type LinkTwoWayMethod interface {
	InsertNodeByHead(data interface{})                           //插入 - 头插
	InsertNodeByIndex(index int32, data interface{}) (err error) //插入 - 指定下标插入
	Append(data interface{})                                     //插入 - 尾部追加
	Delete(index int32) error                                    //删除
	Search(index int32) (node *NodeTwoWay, err error)            //查找
	GetCount() int32                                             //获取个数
	RemoveHead() error                                           //移除头节点
	RemoveTail() error                                           //移除尾节点
	RemoveByIndex(index int32) error                             //通过索引移除节点
}

//LinkList 双向列表列表
type LinkList struct {
	Count int32
	Head  *NodeTwoWay //表头
	Tail  *NodeTwoWay //表尾
}

//NodeTwoWay 双向链表节点
type NodeTwoWay struct {
	Data interface{} //数据域
	Next *NodeTwoWay //下个指针域
	Prev *NodeTwoWay //上个指针域
}

//CreateLinkTwoWay 创建双向链表
func CreateLinkTwoWay() *LinkList {
	return &LinkList{}
}

/*
CreateNodeTwoWay
@Desc: data interface{}	节点数据
*/
func CreateNodeTwoWay(data interface{}) *NodeTwoWay {
	return &NodeTwoWay{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}

/*
InsertNodeByHead
@Desc: 头插
@Param: data	int32	插入数据
*/
func (l *LinkList) InsertNodeByHead(data interface{}) {
	newNode := CreateNodeTwoWay(data)
	newNode.Next = l.Head

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Head.Prev = newNode
		l.Head = newNode
	}

	l.Count++
}

/*
InsertNodeByIndex
@Desc 	指定索引位置插入
@Param 	index	int32		索引位置
@Param 	data	interface	插入数据
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
@Desc: 尾插 - 追加节点
@Param: data	int32	插入数据
*/
func (l *LinkList) Append(data interface{}) {
	newNode := CreateNodeTwoWay(data)

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}

	l.Count++
}

/*
Search
@Desc 	搜索信息
@Param	index int32	索引位置
*/
func (l *LinkList) Search(index int32) (node *NodeTwoWay, err error) {
	if index < 0 || index >= l.Count {
		err = errors.New("index out of range for search node")
		return
	}

	node = l.Head
	for i := int32(0); i < index; i++ {
		node = node.Next
	}

	return
}

/*
Delete
@Desc 	删除节点
@Param	index int32	索引位置
*/
func (l *LinkList) Delete(index int32) (err error) {
	if index < 0 || index >= l.Count {
		err = errors.New("index out of range for delete node")
		return
	}

	switch index {
	case 0:
		err = l.RemoveHead() //删除头节点
	case l.Count - 1:
		err = l.RemoveTail() //删除尾节点
	default:
		err = l.RemoveByIndex(index) //删除指定索引节点
	}

	return
}

/*
RemoveHead
@Desc	删除头节点
*/
func (l *LinkList) RemoveHead() (err error) {
	if l.Head != nil {
		nextNode := l.Head.Next
		nextNode.Prev = nil
		l.Head = nextNode
		l.Count--
	} else {
		err = errors.New("func RemoveHead is failed")
	}
	return
}

/*
RemoveTail
@Desc	删除尾节点
*/
func (l *LinkList) RemoveTail() (err error) {
	if l.Tail != nil && l.Tail.Prev != nil {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		l.Count--
	} else {
		err = errors.New("func RemoveTail is failed")
	}
	return
}

/*
RemoveByIndex
@Desc	删除指定索引节点
@Param	index int32	指定节点索引
*/
func (l *LinkList) RemoveByIndex(index int32) (err error) {
	curNode, err := l.Search(index)
	if err != nil {
		return
	}

	prevNode := curNode.Prev
	nextNode := curNode.Next
	if prevNode != nil && nextNode != nil {
		nextNode.Prev = prevNode
		prevNode.Next = nextNode
		l.Count--
	} else {
		err = errors.New("func RemoveIndex is failed")
	}

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
@Desc: 搜索指定位置节点信息
@Return: int32	数据个数
*/
func (l *LinkList) Println() {
	data := l.Head
	for {
		if data == nil {
			break
		}

		fmt.Printf("data：%v | pointer：%p | prevPointer：%p | nextPointer：%p\n", data.Data, data, data.Prev, data.Next)
		data = data.Next
	}

	if l.Head != nil {
		fmt.Printf("headData:%v | headPoint:%p | headPrevPoint:%p | headNextPoint:%p\n", l.Head.Data, l.Head, l.Head.Prev, l.Head.Next)
		fmt.Printf("tailData:%v | tailPoint:%p | tailPrevPoint:%p | tailNextPoint:%p\n", l.Tail.Data, l.Tail, l.Tail.Prev, l.Tail.Next)
	}

	fmt.Printf("count:%v | head:%p | tail:%p\n", l.Count, l.Head, l.Tail)
	fmt.Println()
}
