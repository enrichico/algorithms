// Package quickunion contains the quickunion approach (lazy approach) to solve the (union find) dynamic connectivity problem.
// It is a set of trees structure (forest). Each entry in the array will contain a reference to its parent in the tree
// E.g. id := [...]int{0, 1, 9, 4, 9, 6, 6, 7, 8, 9}
package quickunion

import "fmt"

type collection struct {
	id []int
}

func New(n int) *collection {

	c := new(collection)

	c.id = make([]int, n)
	//initialize slice, each id to its own index
	for i := 0; i < n; i++ {
		c.id[i] = i
	}

	return c
}

func (c *collection) root(i int) int {

	//follow parents until root is found
	for i != c.id[i] {
		i = c.id[i]
	}

	return i
}

func (c *collection) Connected(p int, q int) bool {

	return c.root(p) == c.root(q)
}

func (c *collection) Union(p int, q int) {

	//put root of p under root of q
	i := c.root(p)
	j := c.root(q)
	c.id[i] = j
}

func (c *collection) PrintIds() {

	fmt.Printf("%v\n", c.id)
}
