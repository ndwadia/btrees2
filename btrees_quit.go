package main

import (
	"fmt"

	tree "example.com/ndwadia/btrees/code"
)

func walkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// Value successfully sent.
	case <-quit:
		return
	}
	walkImpl(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	walkImpl(t, ch, quit)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, w1, quit)
	go Walk(t2, w2, quit)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		fmt.Println("Ch1 ", v1, ok1)
		fmt.Println("Ch2 ", v2, ok2)
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
