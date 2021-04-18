//Construct and print a random binary tree
package tree

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

/*func main() {
	intMultiplier, _ := strconv.Atoi(os.Args[1])
	nodeCount, _ := strconv.Atoi(os.Args[2])
	t := New(intMultiplier, nodeCount)
	fmt.Println(t)
	t = New(intMultiplier, nodeCount)
	fmt.Println(t)
}*/
