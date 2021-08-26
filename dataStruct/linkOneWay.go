package dataStruct

import (
	"errors"
	"fmt"
)

//LinkMethod 单向链表实现接口
type LinkMethod interface {
	InsertNodeByHead(data interface{})                           //插入 - 头插
	InsertNodeByIndex(index int32, data interface{}) (err error) //插入 - 指定下标插入
	Append(data interface{})                                     //插入 - 尾部追加
	Delete(index int32) error                                    //删除
	Search(index int32) (node *NodeOneWay, err error)            //查找
	GetCount() int32                                             //获取个数
}

//LinkOneWayModel 单向链表数据结构
type LinkOneWayModel struct {
	Count int32       //数据个数
	Head  *NodeOneWay //表头
	Tail  *NodeOneWay //表尾
}

//NodeOneWay 链表节点
type NodeOneWay struct {
	Data interface{} //数据域
	Next *NodeOneWay //指针域
}

/*
CreatedLink
@Desc 初始化链表
@Return	*LinkModel
*/
func CreatedLink() *LinkOneWayModel {
	return &LinkOneWayModel{}
}

/*
CreateNode
@Desc: 创建节点
@Param: data int32	节点数据
@Return:	*Node	新节点
*/
func CreateNode(data interface{}) *NodeOneWay {
	return &NodeOneWay{
		Next: nil,
		Data: data,
	}
}

/*
InsertNodeByIndex
@Desc:下标插入
@Param: index 	int32	在第几个位置插入
@Param: data 	int32	插入数据
*/
func (l *LinkOneWayModel) InsertNodeByIndex(index int32, data interface{}) (err error) {
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
	newNode := CreateNode(data)
	curNode, err := l.Search(index - 1)
	if err != nil {
		err = errors.New("search by index out of size")
		return
	}

	//success
	newNode.Next = curNode.Next
	curNode.Next = newNode
	l.Count++
	return
}

/*
InsertNodeByHead
@Desc: 头插
@Param: data	int32	插入数据
*/
func (l *LinkOneWayModel) InsertNodeByHead(data interface{}) {
	newNode := CreateNode(data)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}

	l.Count++
}

/*
Append
@Desc: 尾插 - 追加节点
@Param: data	int32	插入数据
*/
func (l *LinkOneWayModel) Append(data interface{}) {
	newNode := CreateNode(data)
	if l.Tail == nil {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}

	l.Count++
	l.Tail = newNode
}

/*
Delete
@Desc: 删除节点
@Param: index	int32	删除第几个节点数据
*/
func (l *LinkOneWayModel) Delete(index int32) error {
	if index < 0 || index >= l.Count {
		return errors.New("func Delete index out of size")
	}

	//找到下个节点
	var preNode *NodeOneWay
	currentNode := l.Head
	for i := int32(0); i < index; i++ {
		preNode = currentNode
		currentNode = currentNode.Next
	}

	//如果是删除最后一个则要更新tail
	if index == l.Count-1 {
		l.Tail = preNode
	}

	//删除
	if index == 0 {
		l.Head = currentNode.Next
	} else {
		preNode.Next = currentNode.Next
	}

	l.Count--
	return nil
}

/*
GetCount
@Desc: 获得个数
@Return: int32	数据个数
*/
func (l *LinkOneWayModel) GetCount() int32 {
	return l.Count
}

/*
Search
@Desc: 搜索指定位置节点信息
@Return: int32	数据个数
*/
func (l *LinkOneWayModel) Search(index int32) (node *NodeOneWay, err error) {
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
Println
@Desc: 搜索指定位置节点信息
@Return: int32	数据个数
*/
func (l *LinkOneWayModel) Println() {
	data := l.Head
	for {
		if data == nil {
			break
		}

		fmt.Printf("data：%v pointer：%p nextPointer：%p\n", data.Data, data, data.Next)
		data = data.Next
	}

	fmt.Println()
}
