package mymap

import(
	"fmt"
)

func ShowMap() {
	ages := make(map[string]int)
	ages["shit"] = 12

	ages2 := map[string]string{
		"one": "this is one",
		"two": "this is two",
		"three": "this is three"}

	fmt.Println(ages["shit"])
	
	println("=========")

	for i, v := range ages2 {
		fmt.Println(i, v)
	}
}