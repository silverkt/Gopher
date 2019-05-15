package main

import (
	"fmt"
	"reflect"
)


/**
利用反射包，获取类型
*/
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	a := [3]int{1, 2, 3}
	b := a
	b[1]++
	fmt.Println(a, b)

	c := [3]int{3, 2, 1}
	d := &c
	d[1]++
	fmt.Println(c, *d)
	fmt.Println(typeof(d))
	fmt.Println(d[1], (*d)[1])
}