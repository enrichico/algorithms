// Package unionfind contains approaches(lazy approach) to solve the (union find) dynamic connectivity problem.
// In this case, "weighted quick union + path compression".
// It consists of a set of trees structure (forest). Each entry in the array will contain a reference to its parent in the tree
// E.g. id := [...]int{0, 1, 9, 4, 9, 6, 6, 7, 8, 9}
// Maintains an extra array "size[i]" to count the number of objects in the tree rooted at i
// Adds path compression, making every other node (in the path to the root) pointing to its grandparent
package unionfind

import "fmt"

type quickUnionWPC struct {
	id   []int
	size []int
}

func NewQuickUnionWPC(n int) *quickUnionWPC {

	qu := new(quickUnionWPC)

	qu.id = make([]int, n)
	qu.size = make([]int, n)

	//initialize slice, each id to its own index
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}

	return qu
}

func (qu *quickUnionWPC) root(i int) int {

	//follow parents until root is found
	for i != qu.id[i] {
		qu.id[i] = qu.id[qu.id[i]] //path compression
		i = qu.id[i]
	}

	return i
}

func (qu *quickUnionWPC) Connected(p int, q int) bool {

	return qu.root(p) == qu.root(q)
}

func (qu *quickUnionWPC) Union(p int, q int) {

	i := qu.root(p)
	j := qu.root(q)
	if i == j {
		return
	}

	//link root of smaller tree to root of larger tree
	if qu.size[i] < qu.size[j] {
		qu.id[i] = j
		qu.size[j] += qu.size[i] //update size
	} else {
		qu.id[j] = i
		qu.size[i] += qu.size[j]
	}
}

func (qu *quickUnionWPC) PrintIds() {

	fmt.Printf("%v\n", qu.id)
}
