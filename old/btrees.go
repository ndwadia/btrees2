package main

import (
	"fmt"

	tree "example.com/ndwadia/btrees/code"
)

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// NOTE: The implementation leaks goroutines when trees are different.
// See binarytrees_quit.go for a better solution.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		//fmt.Println("Ch 1:", v1, ok1)
		//fmt.Println("Ch 2:", v2, ok2)
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	//fmt.Println(t1)
	//fmt.Println(t2)
	fmt.Println("tree.New(1) == tree.New(1): Test status:")
	if Same(t1, t2) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}

	t2 = tree.New(2)
	//fmt.Println(t1)
	//fmt.Println(t2)
	fmt.Println("tree.New(1) != tree.New(2): Test Status:")
	if !Same(t1, t2) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
