# algorithms
Implementation of different algorithms in Golang 

These algorithms are based on the class "Algorithms, Part I" by Robert Sedgewick available for free on Coursera (https://www.coursera.org/course/algs4partI)

E.g. the quickfind algorithm to solve the "dynamic connectivity problem"

package main

import (
	"fmt"
	"github.com/enrichico/algorithms/quickfind"
)

func main() {

	//a collection of 10 elements
	c := quickfind.New(10)

	//sample union operations
	c.Union(6, 4)
	c.Union(0, 8)
	c.Union(7, 8)
	c.Union(6, 2)
	c.Union(0, 1)
	c.Union(3, 9)

	fmt.Printf("The current array of ids is: \n")
	c.PrintIds()
}

