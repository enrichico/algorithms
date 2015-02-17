// Package unionfind contains approaches to solve the (union find) dynamic connectivity problem.
package unionfind

import "fmt"

type quickFind struct {
	id []int
}

func NewQuickFind(n int) *quickFind {

	qf := new(quickFind)

	qf.id = make([]int, n)
	//initialize slice, each id to its own index
	for i := 0; i < n; i++ {
		qf.id[i] = i
	}

	return qf
}

func (qf *quickFind) Connected(p int, q int) bool {

	return qf.id[p] == qf.id[q]
}

func (qf *quickFind) Union(p int, q int) {

	pid := qf.id[p]
	qid := qf.id[q]

	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
}

func (qf *quickFind) PrintIds() {

	fmt.Printf("%v\n", qf.id)
}
