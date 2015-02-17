# algorithms
Implementation of different algorithms in Golang 

These algorithms are based on the class "Algorithms, Part I" by Robert Sedgewick available for free on Coursera (https://www.coursera.org/course/algs4partI)

E.g. the quickfind algorithm to solve the "dynamic connectivity problem"

package main

import (
	"fmt"
	"github.com/enrichico/algorithms/unionfind"
)

func main() {

	//a collection of 10 elements
	qf := unionfind.NewQuickFind(10)

	//sample union operations
	qf.Union(6, 4)
	qf.Union(0, 8)
	qf.Union(7, 8)
	qf.Union(6, 2)
	qf.Union(0, 1)
	qf.Union(3, 9)

	fmt.Printf("The current array of ids is: \n")
	qf.PrintIds()
}

