package main
import(
	"fmt"
	aliasName "../test"
	myArray "../array"
	mySlice "../slice"
)

func main() {
	fmt.Println("Hello World!")
	aliasName.Tt()
	aliasName.TestStruct()
	myArray.ShowArray()
	mySlice.ShowSlice()
}