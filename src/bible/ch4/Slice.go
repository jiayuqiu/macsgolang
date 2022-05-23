package ch4

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func Practice43Reverse(s *[5]int){
	// 重写reverse函数，使用数组指针代替slice。
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Practice45Dup(s []string) []string {
	if len(s) == 0 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		fmt.Printf("i = %d, j = %d \n", i, j)
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		} else{
			fmt.Println(">>>>>>>>>>>")
		}
	}
	return s[:i+1]
}


func Practice46(utf8Bytes []byte) []byte{
	//编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
	if len(utf8Bytes) == 0 {
		return utf8Bytes
	}

	// out 用于输出
	var out []byte

	// last, 由于记录上一个字符的码点
	var last rune

	for i := 0; i < len(utf8Bytes); {
		r, s := utf8.DecodeRune(utf8Bytes[i: ])

		if !unicode.IsSpace(r){
			// 若字符不为空，记录
			out = append(out, utf8Bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last){
			// 若当前字符为空，但是last字符非空，则out添加一个空格
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}

func Practice47(s []byte) []byte {
	// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
	fmt.Println("Practice 4.7 ")
	var out []byte

	// https://github.com/ray-g/gopl/tree/master/ch04/ex4.07

	return out
}
