package main

import (
	"macsgolang/src/bible/ch1"
	"macsgolang/src/bible/ch5"
)

func main() {
	// 调用自定义函数
	ch1.PrintHello()

	// web
	// ch1.FetchUrlByChan()

	//flag
	//ch2.TestFlag()
	//s := ch2.TestAssignmentSlice()
	//fmt.Println(">>>>>>>>", s)

	// int
	//ch3.TestInt()

	// float surface
	//ch3.Surface()

	// array
	//ch4.TestArray()

	// slice
	//var aPtr *[5]int
	//var a = [5]int{2,4,6,8,10}
	//aPtr = &a  // 生成数组指针的方法
	//ch4.Practice43Reverse(aPtr)
	//fmt.Println(a)
	//
	//var s = []string{"zero", "one", "one", "two", "three"}
	//var s2 []string
	//s2 = ch4.Practice45Dup(s)
	//fmt.Println(s2)
	//
	//s46 := "hello      世界 46"
	//b46 := []byte(s46)
	//fmt.Println(s46)
	//b46Res := ch4.Practice46(b46)
	//fmt.Println(string(b46Res))

	// map
	//ch4.Practice48CharCount()
	//ch4.MapDemo1()

	// json
	//str := `{"page": 1, "fruits": ["apple", "peach"]}`
	//res := ch4.TestJsonStruct{}
	//json.Unmarshal([]byte(str), &res)
	//fmt.Println(res)
	//fmt.Println(res.Page)
	// ch4.Practice412()
	ch5.Findlinks1()

	// fmt.Printf("%v %T", time.Second, time.Second)
}
