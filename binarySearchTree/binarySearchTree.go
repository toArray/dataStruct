package binarySearchTree

import "fmt"

//BinarySearchTree 二叉搜索树实现接口
type BinarySearchTree interface {
	Insert(data int32) //插入节点
	//Delete()           //删除
	//Search()           //搜索
	GetHeight() //获得树高度

	RangeLeft(node *Node)  //左序遍历
	RangeMid(node *Node)   //中序遍历
	RangeRight(node *Node) //右序遍历
	RangeLayer()           //层序遍历（广度优先遍历）
}

//Node node
type Node struct {
	Data  int32 //数据域
	Left  *Node //左子树
	Right *Node //右子树
}

type Tree struct {
	Root *Node
}

//InitBinarySearchTree init tree
func InitBinarySearchTree() *Tree {
	return &Tree{}
}

/*
Insert
@Desc	插入节点
@Param	data int32
*/
func (t *Tree) Insert(data int32) {
	//tree is nil
	if t.Root == nil {
		t.Root = &Node{
			Data: data,
		}
		return
	}

	temp := t.Root
	for {
		//insert left
		if data < temp.Data {
			if temp.Left != nil {
				temp = temp.Left
				continue
			}

			temp.Left = &Node{
				Data: data,
			}
			break
		}

		//insert right
		if temp.Right != nil {
			temp = temp.Right
			continue
		}

		temp.Right = &Node{
			Data: data,
		}
		break
	}
}

/*
RangeLeft
@Desc 左序遍历
*/
func (t *Tree) RangeLeft(node *Node) {
	if node != nil {
		fmt.Printf("%v ", node.Data)
		t.RangeLeft(node.Left)
		t.RangeLeft(node.Right)
	}
}

/*
RangeMid
@Desc 中序遍历
*/
func (t *Tree) RangeMid(node *Node) {
	if node != nil {
		t.RangeMid(node.Left)
		fmt.Printf("%v ", node.Data)
		t.RangeMid(node.Right)
	}
}

/*
RangeRight
@Desc 右序遍历
*/
func (t *Tree) RangeRight(node *Node) {
	if node != nil {
		t.RangeRight(node.Left)
		t.RangeRight(node.Right)
		fmt.Printf("%v ", node.Data)
	}
}

/*
RangeLayer
@Desc 层序遍历
*/
func (t *Tree) RangeLayer() {
	if t.Root == nil {
		return
	}

	//use queue
	queue := []*Node{t.Root}

	for len(queue) > 0 {
		node := queue[0]
		fmt.Printf("%v ", node.Data)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		queue = queue[1:]
	}

}

/*
GetHeight
@Desc 获得树的高度
*/
func (t *Tree) GetHeight() {
	height := getHeight(t.Root)
	fmt.Println(height)
}

//getHeight
func getHeight(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)
	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}
	return max + 1
}
