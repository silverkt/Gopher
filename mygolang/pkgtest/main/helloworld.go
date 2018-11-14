package main
import(
	"fmt"
	aliasName "../test"
	myArray "../array"
	mySlice "../slice"
	myMap "../map"
	tinf "../interface"
)

func main() {
	fmt.Println("Hello World!")
	aliasName.Tt()
	aliasName.TestStruct()
	myArray.ShowArray()
	mySlice.ShowSlice()
	myMap.ShowMap()
	mytestinfterfaceobj := new(tinf.Square)
	mytestinfterfaceobj.TestInf()

	var myname tinf.Namer
	myname = mytestinfterfaceobj
	fmt.Println(myname.TestInterface2())
}