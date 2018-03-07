package main
import(
	"fmt"
	aliasName "../test"
	myArray "../array"
	mySlice "../slice"
	myMap "../map"
)

func main() {
	fmt.Println("Hello World!")
	aliasName.Tt()
	aliasName.TestStruct()
	myArray.ShowArray()
	mySlice.ShowSlice()
	myMap.ShowMap()
}