package binary

type Node struct {
	Data   int32
	Left   *Node
	Right  *Node
	Parent *Node
}
type Tree struct {
	Root *Node
}

func (t *Tree) InsertNode(data int32) {
	//根节点
	if t.Root == nil {
		t.Root = &Node{
			Data:   data,
			Left:   nil,
			Right:  nil,
			Parent: nil,
		}
		return
	}

	tempNode := t.Root
	for {
		if tempNode.Data > data {
			if tempNode.Left != nil {
				tempNode = tempNode.Left
			} else {
				tempNode.Left = &Node{
					Data:   data,
					Left:   nil,
					Right:  nil,
					Parent: tempNode,
				}
				break
			}
		} else {
			if tempNode.Right != nil {
				tempNode = tempNode.Right
			} else {
				tempNode.Right = &Node{
					Data:   data,
					Left:   nil,
					Right:  nil,
					Parent: tempNode,
				}
				break
			}
		}
	}
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

func (t *Tree) Search(data int32) *Node {
	temp := t.Root
	for temp != nil {
		if temp.Data == data {
			return temp
		}

		if temp.Data > data {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}

	return nil
}

func (t *Tree) Delete(data int32) {
	node := t.Search(data)
	if node == nil {
		return
	}

	//是叶子节点
	parentNode := node.Parent
	if node.Left == nil && node.Right == nil {
		if node == t.Root {
			t.Root = nil
		} else {
			if parentNode.Left == node {
				parentNode.Left = nil
			} else {
				parentNode.Right = nil
			}
		}
		return
	}

	//只有一个叶子节点,假设删除4和10
	/*
				   6
				/    \
			  4        10
			    \     /
		         5   9
	*/
	if node.Left == nil && node.Right != nil || node.Left != nil && node.Right == nil {
		replaceNode := node.Left
		if node.Right != nil {
			replaceNode = node.Right
		}

		if node == t.Root {
			replaceNode.Parent = nil
			t.Root = replaceNode
		} else {
			replaceNode.Parent = node.Parent
			if parentNode.Left == node {
				parentNode.Left = replaceNode
			} else {
				parentNode.Right = replaceNode
			}
		}
		return
	}

	//有2个节点
	/*
				   6
				/    \
			  4       10
			/  \     /   \
		   3    5   9     11
					 \
					  10
	*/
	//假设删除4或者10,后继节点刚好是右节点
	replaceNode := RightMin(node)
	if node.Right == replaceNode {
		if parentNode == nil {
			replaceNode.Parent = nil
			t.Root = replaceNode
		} else {
			replaceNode.Parent = parentNode
			if parentNode.Left == node {
				parentNode.Left = replaceNode
			} else {
				parentNode.Right = replaceNode
			}
		}

		if node.Left != nil {
			replaceNode.Left = node.Left
			node.Left.Parent = replaceNode
		}
	} else {
		/*
						   6
						/    \
					  4       12
					/  \     /   \
				   3    5   10     13
						  /	 \		\
			            8	  11	 14
						 \
						  9
		*/
		//假设删除6
		replaceParentNode := replaceNode.Parent //10
		if parentNode == nil {
			replaceNode.Parent = nil
			t.Root = replaceNode

			if replaceNode.Right != nil {
				replaceParentNode.Left = replaceNode.Right
				replaceNode.Right.Parent = replaceParentNode
			}

			if node.Left != nil {
				node.Left.Parent = replaceParentNode
				replaceNode.Left = node.Left
			}
		} else {
			//维护替代节点
			if replaceNode.Right != nil {
				replaceParentNode.Left = replaceNode.Right
				replaceNode.Right.Parent = replaceParentNode
			} else {
				replaceParentNode.Left = nil
			}

		}
	}
}

//后继节点
func RightMin(node *Node) *Node {
	minNode := node.Right
	var parent *Node
	for {
		parent = minNode
		if minNode.Left != nil {
			minNode = minNode.Left
		} else {
			break
		}
	}

	return parent
}
