package ch4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func MapDemo1() {
	var names []string
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"macs":    28,
		"lynn":    18,
	}

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for i, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
		fmt.Println(i)
	}
}

func CharCount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	cc := 0
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		cc++
		if cc > 50 {
			break
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//func Practice48Charcount(){
//	counts := make(map[rune]int)    // counts of Unicode characters
//	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
//	invalid := 0                    // count of invalid UTF-8 characters
//
//	in := bufio.NewReader(os.Stdin)
//}
