// Package quickunion contains the quickunion approach (lazy approach) to solve the (union find) dynamic connectivity problem.
// It is a set of trees structure (forest). Each entry in the array will contain a reference to its parent in the tree
// E.g. id := [...]int{0, 1, 9, 4, 9, 6, 6, 7, 8, 9}
package unionfind

import "fmt"

type QuickUnion struct {
	id []int
}

func NewQuickUnion(n int) *QuickUnion {

	qu := new(QuickUnion)

	qu.id = make([]int, n)
	//initialize slice, each id to its own index
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}

	return qu
}

func (qu *QuickUnion) root(i int) int {

	//follow parents until root is found
	for i != qu.id[i] {
		i = qu.id[i]
	}

	return i
}

func (qu *QuickUnion) Connected(p int, q int) bool {

	return qu.root(p) == qu.root(q)
}

func (qu *QuickUnion) Union(p int, q int) {

	//put root of p under root of q
	i := qu.root(p)
	j := qu.root(q)
	qu.id[i] = j
}

func (qu *QuickUnion) PrintIds() {

	fmt.Printf("%v\n", qu.id)
}
