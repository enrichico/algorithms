// Package unionfind contains approaches(lazy approach) to solve the (union find) dynamic connectivity problem.
// In this case, "quick union" (also called lazy apporach).
// It consists of a set of trees structure (forest). Each entry in the array will contain a reference to its parent in the tree
// E.g. id := [...]int{0, 1, 9, 4, 9, 6, 6, 7, 8, 9}
package unionfind

import "fmt"

type quickUnion struct {
	id []int
}

func NewQuickUnion(n int) *quickUnion {

	qu := new(quickUnion)

	qu.id = make([]int, n)
	//initialize slice, each id to its own index
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}

	return qu
}

func (qu *quickUnion) root(i int) int {

	//follow parents until root is found
	for i != qu.id[i] {
		i = qu.id[i]
	}

	return i
}

func (qu *quickUnion) Connected(p int, q int) bool {

	return qu.root(p) == qu.root(q)
}

func (qu *quickUnion) Union(p int, q int) {

	//put root of p under root of q
	i := qu.root(p)
	j := qu.root(q)
	qu.id[i] = j
}

func (qu *quickUnion) PrintIds() {

	fmt.Printf("%v\n", qu.id)
}
