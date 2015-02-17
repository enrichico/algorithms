// Package quickfind contains the quickfind approach to solve the (union-find) dynamic connectivity problem
package quickfind

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

func (c *collection) Connected(p int, q int) bool {

	return c.id[p] == c.id[q]
}

func (c *collection) Union(p int, q int) {

	pid := c.id[p]
	qid := c.id[q]

	for i := 0; i < len(c.id); i++ {
		if c.id[i] == pid {
			c.id[i] = qid
		}
	}
}

func (c *collection) PrintIds() {

	fmt.Printf("%v\n", c.id)
}
