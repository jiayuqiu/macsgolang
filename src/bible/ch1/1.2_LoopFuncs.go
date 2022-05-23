package ch1

import (
	"fmt"
	"strings"
)

func PrintArgs(agrs []string) {
	fmt.Println(strings.Join(agrs, "@@"))
}

func TestPointer(p *int) int {
	*p++
	return *p
}


