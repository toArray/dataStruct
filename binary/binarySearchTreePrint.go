package binary

import (
	"fmt"
	"strconv"
	"strings"
)

func Show(root *Node) {
	if root == nil {
		fmt.Println("EMPTY!")
		return
	}

	// 得到树的深度
	treeDepth := getHeight(root)

	// 最后一行的宽度为2的（n - 1）次方乘3，再加1
	// 作为整个二维数组的宽度
	arrayHeight := treeDepth*2 - 1
	temp := treeDepth - 2
	if temp < 0 {
		temp = 0
	}
	arrayWidth := (2<<(temp))*3 + 1
	//arrayWidth := (2<<(treeDepth))*3 + 1
	// 用一个字符串数组来存储每个位置应显示的元素
	res := make([][]string, arrayHeight)
	for i := 0; i < arrayHeight; i++ {
		res[i] = make([]string, arrayWidth)
	}

	// 对数组进行初始化，默认为一个空格
	for i := 0; i < arrayHeight; i++ {
		for j := 0; j < arrayWidth; j++ {
			res[i][j] = " "
		}
	}

	// 从根节点开始，递归处理整个树
	writeArray(root, 0, arrayWidth/2, res, treeDepth)

	// 此时，已经将所有需要显示的元素储存到了二维数组中，将其拼接并打印即可
	for _, value := range res {
		var build strings.Builder
		for i := 0; i < len(value); i++ {
			build.WriteString(fmt.Sprintf("%s", value[i]))
			if len(value) > 1 && i < len(value)-1 {
				temp := len(value[i])
				if temp > 4 {
					i += 2
				} else {
					i += temp - 1
				}
			}
		}
		fmt.Println(build.String())
	}
}

func writeArray(currNode *Node, rowIndex int, columnIndex int, res [][]string, treeDepth int) {
	//保证输入的树不为空
	if currNode == nil {
		return
	}

	//先将当前节点保存到二维数组中
	res[rowIndex][columnIndex] = strconv.Itoa(int(currNode.Data))

	//计算当前位于树的第几层
	currLevel := (rowIndex + 1) / 2

	//若到了最后一层，则返回
	if currLevel == treeDepth {
		return
	}

	//计算当前行到下一行，每个元素之间的间隔（下一行的列索引与当前元素的列索引之间的间隔）
	gap := treeDepth - currLevel - 1

	//对左儿子进行判断，若有左儿子，则记录相应的"/"与左儿子的值
	if currNode.Left != nil {
		res[rowIndex+1][columnIndex-gap] = "/"
		writeArray(currNode.Left, rowIndex+2, columnIndex-gap*2, res, treeDepth)
	}

	//对右儿子进行判断，若有右儿子，则记录相应的"\"与右儿子的值
	if currNode.Right != nil {
		res[rowIndex+1][columnIndex+gap] = "\\"
		writeArray(currNode.Right, rowIndex+2, columnIndex+gap*2, res, treeDepth)
	}
}
