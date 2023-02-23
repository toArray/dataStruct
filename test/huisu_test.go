package test

import (
	"fmt"
	"testing"
)

var data = []int32{1, 2, 3, 4}
var res = [][]int32{}
var path = []int32{}

func TestAAA(t *testing.T) {

	fmt.Println(res)
	a(4, 2)
	fmt.Println(res)
}

func backtracking(n, k, startIndex int, path []int32) {
	if len(path) == k {
		res = append(res, path)
		return
	}

	for i := startIndex; i < n; i++ {
		path = append(path, data[i])
		startIndex += 1
		backtracking(n, k, startIndex, path)
		path = path[:len(path)-1]
	}
}

func a(n, k int) {
	backtracking(n, k, 0, path)
}
