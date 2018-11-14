package myarray

import(
	"fmt"
)

var a [3]int = [...]int{2, 5, 7}
var b string = "This is just a shit"
func ShowArray() {
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}
	for i, v := range a {
		fmt.Printf("%d => %d\n", i, v)
	}

	for i, v := range b {
		fmt.Printf("%d => %c\n", i, v)
	}

}