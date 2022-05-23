package ch3

import "fmt"

func TestInt() {
	var x uint8 = 1<<1 | 1<<5  // "00100010"

	for j := uint(0); j < 8; j ++ {
		fmt.Printf("j unit(0) = %08b\n", j)
	}


	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println("i = ", i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
