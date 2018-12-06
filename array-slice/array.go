package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := a
	b[1]++
	fmt.Println(a, b)

	c := [3]int{3, 2, 1}
	d := &c
	d[1]++
	fmt.Println(c, *d)
}