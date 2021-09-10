package test

import (
	"fmt"
	"testing"
)

func TestLiKou(t *testing.T) {
	s := "()"
	l := len(s)
	if l%2 == 1 {
		fmt.Println(false)
		return
	}

	//所有字符串
	allStr := make(map[byte]byte)
	allStr['('] = ')'
	allStr['['] = ']'
	allStr['{'] = '}'

	stack := []int32{}
	for _, value := range s {
		//左括号入栈
		_, isLeft := allStr[byte(value)]
		if isLeft {
			stack = append(stack, value)
			continue
		}

		//验证括号
		if len(stack) == 0 || byte(value) != allStr[byte(stack[len(stack)-1])] {
			fmt.Println(false)
			return
		}

		//删除
		stack = stack[:len(stack)-1]
	}

	fmt.Println(len(stack) == 0)
}
