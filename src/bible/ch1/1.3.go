package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func DropDuplicatesLines(){
	var counts map[string]int
	counts = make(map[string]int)
	files := [2]string{"/Users/qiujiayu/GolandProjects/awesomeProject/utils/data1.csv",
		"/Users/qiujiayu/GolandProjects/awesomeProject/utils/data2.csv"}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		CountLines(f, counts)
		f.Close()
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d \t %s \n", n, line)
		}
	}
}

func CountLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
