// Package quickfind contains different approaches to solve the (union find) dynamic connectivity problem
package quickfind

import "fmt"

// QuickFind is an implementation of the QuickFind algorithm (eager approch) to resolve the dynamic connectivity problem (union find).
type QuickFind struct {
	id []int
}

func New(n int) *QuickFind {

	qf := new(QuickFind)

	qf.id = make([]int, n)
	//initialize slice
	for i := 0; i < n; i++ {
		qf.id[i] = i
	}

	return qf
}

func (qf *QuickFind) Connected(p int, q int) bool {

	return qf.id[p] == qf.id[q]
}

func (qf *QuickFind) Union(p int, q int) {

	pid := qf.id[p]
	qid := qf.id[q]

	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
}

func (qf *QuickFind) PrintIds() {

	fmt.Printf("%v\n", qf.id)
}
