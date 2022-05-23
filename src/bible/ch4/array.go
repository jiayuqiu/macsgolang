package ch4

import "fmt"

func TestArray(){
	// 生成了一个数组a, 其中a的index=2处，数值为0
	// golang中，生成的数组中，若没有设置数值，则默认为0
	var a [3]int = [3]int{1, 2}
	var b [3]int
	fmt.Println(a) // "0"
	fmt.Println(b)

	// 生成了一个100个元素的数组c，其中index=99处的数值为-1
	// index in [0, 98]闭区间的数值，都为0
	c := [...]int{99: -1}
	fmt.Println(c[90: ])
}
