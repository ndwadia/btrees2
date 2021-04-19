package main

import (
	"fmt"

	compare "example.com/ndwadia/btrees2/code"
	tree "example.com/ndwadia/btrees2/code"
)

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	//fmt.Println(t1)
	//fmt.Println(t2)
	fmt.Println("tree.New(1) == tree.New(1): Test status:")
	if compare.Same(t1, t2) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}

	t2 = tree.New(2)
	//fmt.Println(t1)
	//fmt.Println(t2)
	fmt.Println("tree.New(1) != tree.New(2): Test Status:")
	if !compare.Same(t1, t2) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
