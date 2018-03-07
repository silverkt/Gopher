package myslice

import(
	 "fmt"
)

func ShowSlice() {
	mouths := [...]string{
		1: "1月",
		2: "2月",
		3: "3月",
		4: "4月",
		5: "5月"}
	for i, v := range mouths {
		fmt.Println(i,v)
	}	
	println("------------------")
	first3 := mouths[2:4]
	
	for i, v := range first3 {
		println(i,v)
	}	
	println("------------------")

	s2 := make([]int, 10, 20)

	for i, v := range s2 {
		println(i,v)
	}	
	println("------------------")

	var runes []rune
	for _, r := range "Hello 世界" {
		runes = append(runes, r)
	}

	for i, v := range runes {
		fmt.Printf("%d => %q => %d\n", i, v, v)
	}	
	println("------------------")

}