package avlTree

import (
	"fmt"
)

//AvlTree 平衡二叉树实现接口
type AvlTree interface {
	Insert(data int32)       //插入节点
	Delete(data int32)       //删除节点
	Search(data int32) *Node //搜索节点
	GetHeight() int          //获得高度

	RangeLeft(node *Node)  //左序遍历（深度优先）
	RangeMid(node *Node)   //中序遍历（深度优先）
	RangeRight(node *Node) //右序遍历（深度优先）
	RangeLayer()           //层序遍历（广度优先）

	//左旋
	//右旋
}

type Tree struct {
	Root *Node
}

//Node node
type Node struct {
	Data   int32 //数据域
	Left   *Node //左子树
	Right  *Node //右子树
	Parent *Node //父节点
}

//InitBinarySearchTree init tree
func InitBinarySearchTree() *Tree {
	return &Tree{}
}

func NewNode(data int32, parent *Node) *Node {
	return &Node{
		Data:   data,
		Left:   nil,
		Right:  nil,
		Parent: parent,
	}
}

/*
Insert
@Desc	插入节点
@Param	data int32
*/
func (t *Tree) Insert(data int32) {
	//tree is nil
	if t.Root == nil {
		t.Root = NewNode(data, nil)
		return
	}

	temp := t.Root
	var parent *Node
	for {
		parent = temp

		if data < temp.Data {
			//insert left
			if temp.Left != nil {
				temp = temp.Left
			} else {
				temp.Left = NewNode(data, parent)
				break
			}
		} else {
			//insert right
			if temp.Right != nil {
				temp = temp.Right
			} else {
				temp.Right = NewNode(data, parent)
				break
			}
		}

		//调整树

	}
}

/*
RangeLeft
@Desc 左序遍历-中左右
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
@Desc 中序遍历-左中右
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
@Desc 右序遍历-左右中
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
func (t *Tree) GetHeight() int {
	height := getHeight(t.Root)
	fmt.Println(height)
	return height
}

/*
GetHeight
@Desc 获得树的高度
*/
func GetHeight(node *Node) int {
	height := getHeight(node)
	fmt.Println(height)
	return height
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

/*
Search
@Desc	搜索
@Param	data int32
*/
func (t *Tree) Search(data int32) *Node {
	if t.Root == nil {
		return nil
	}

	temp := t.Root
	for temp != nil {
		if temp.Data == data {
			return temp
		}

		if data > temp.Data {
			temp = temp.Right
			continue
		}

		temp = temp.Left
	}

	return nil
}

/*
Delete
@Desc 删除节点
*/
func (t *Tree) Delete(data int32) {
	//空树
	if t.Root == nil {
		return
	}

	//找到需要删除的节点
	node := t.Search(data)
	if node == nil {
		return
	}

	//该节点的父节点
	parentNode := node.Parent

	//情况1:属于叶子节点
	if node.Left == nil && node.Right == nil {
		//刚好是根节点
		if parentNode == nil {
			t.Root = nil
		} else {
			if parentNode.Left == node {
				parentNode.Left = nil //该节点属于父节点的左节点
			} else {
				parentNode.Right = nil //该节点属于父节点的右节点
			}
		}
		return
	}

	//情况2:该节点只有一个左子节点或者只有一个右子节点
	if (node.Left != nil && node.Right == nil) || (node.Left == nil && node.Right != nil) {
		//替代节点
		replaceNode := node.Left
		if node.Right != nil {
			replaceNode = node.Right
		}

		//刚好是父节点
		if parentNode == nil {
			replaceNode.Parent = nil
			t.Root = replaceNode
		} else {
			replaceNode.Parent = parentNode //更新替代节点的父节点
			if parentNode.Left == node {
				parentNode.Left = replaceNode //该节点属于父节点的左节点
			} else {
				parentNode.Right = replaceNode //该节点属于父节点的右节点
			}
		}
		return
	}

	//情况3:该节点有2个子节点
	if node.Left != nil && node.Right != nil {
		replaceNode := node.rightMinNode() //后继节点

		//情况3-1
		//删除的后继节点刚好是右子树
		if node.Right == replaceNode {
			if parentNode == nil {
				replaceNode.Parent = nil
				t.Root = replaceNode
			} else {
				//维护后继节点的左节点和父节点
				replaceNode.Parent = parentNode
				if parentNode.Left == node {
					parentNode.Left = replaceNode
				} else {
					parentNode.Right = replaceNode
				}
			}

			//更新左子树
			if node.Left != nil {
				replaceNode.Left = node.Left
				node.Left.Parent = replaceNode
			}
		} else {
			//情况3-2
			replaceNodeParent := replaceNode.Parent //代替节点的父节点

			//后继节点没有右子树,维护后继节点的父节点
			if replaceNode.Right == nil {
				if replaceNodeParent.Left == replaceNode {
					replaceNodeParent.Left = nil
				} else {
					replaceNodeParent.Right = nil
				}
			} else {
				//后继节点有右子树,维护后继节点的右节点和父节点
				if replaceNodeParent.Left == replaceNode {
					replaceNodeParent.Left = replaceNode.Right
				} else {
					replaceNodeParent.Right = replaceNode.Right
				}
				replaceNode.Right.Parent = replaceNodeParent
			}

			//维护替代节点的左右子树,左右子树的父节点变更为替代节点
			if node.Left != nil {
				replaceNode.Left = node.Left
				node.Left.Parent = replaceNode
			}
			if node.Right != nil {
				replaceNode.Right = node.Right
				node.Right.Parent = replaceNode
			}

			//维护父节点的左右子树
			replaceNode.Parent = parentNode
			if parentNode == nil {
				t.Root = replaceNode //根节点
			} else {
				if parentNode.Left == node {
					parentNode.Left = replaceNode
				} else {
					parentNode.Right = replaceNode
				}
			}
		}
	}
}

/*
rightMinNode
@Desc 找后继节点
*/
func (n *Node) rightMinNode() *Node {
	min := n.Right
	if min != nil {
		for min.Left != nil {
			min = min.Left
		}
	}

	return min
}

/*
leftMaxNode
@Desc 找前驱节点
*/
func (n *Node) leftMaxNode() *Node {
	max := n.Left
	if max != nil {
		for max.Right != nil {
			max = max.Right
		}
	}

	return max
}

//左旋
/*
调整前
	1
	 \
	  3
	 /  \
	2	 4
调整后
	3
   /  \
  1	   4
   \
	2
*/
func leftRotate(node *Node) {
	if node == nil {
		return
	}

	three := node.Right
	two := three.Left
	node.Right = two
	three.Parent = node.Parent
	node.Parent = three
	if two != nil {
		two.Parent = node
	}
}

//右旋
/*
调整前
		 6
	   /
 	  4
	 /	\
	3	 5
调整后
 	  4
	 /	\
	3	 6
        /
	   5
*/
func rightRotate(node *Node) {
	if node == nil {
		return
	}
}

//平衡英子,左右子树的高度差
func getBalanceValue(node *Node) int {
	if node == nil {
		return 0
	}

	return getHeight(node.Left) - getHeight(node.Right)
}

//平衡
func doBalance(node *Node) {
	if node == nil {
		return
	}
	//高度差
	balanceValue := getBalanceValue(node)

	//右旋（左比右高）
	if balanceValue == 2 {
		leftNode := node.Left
		if leftNode.Right != nil {
			leftRotate(leftNode)
		}
		rightRotate(node)
	}

	//左旋（右比左高）
	if balanceValue == -2 {
		rightNode := node.Right
		if rightNode.Left != nil {
			rightRotate(rightNode)
		}
		leftRotate(node)
	}
}
