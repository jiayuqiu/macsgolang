package ch2

import "fmt"

var c int

func init() {
	c = 1
}

func TestAssignmentSlice() []int{
	// 测试赋值slice
	var testSlice []int // 没有定义长度的slice，需要用append进行添加元素
	testSlice = append(testSlice, 1)
	testSlice = append(testSlice, 2)
	fmt.Println("c = ", c)
	return testSlice
 }
