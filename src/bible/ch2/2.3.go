package ch2

import (
	"flag"
	"fmt"
)


var f = flag.Bool("f", false, "input -f then force is true, else false.")
//var sep = flag.String("sep", " ", "separator")
var name = flag.String("name", "macs", "nameStr, default: macs")

func TestFlag() {
	flag.Parse()
	fmt.Println("  inputs:", flag.Args())
	fmt.Println("force is ", *f)
	if *f {
		//fmt.Println("sep = ", *sep)
		fmt.Println("name = ", *name)
	}
}
