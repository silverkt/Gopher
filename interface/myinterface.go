package myinterface

import(
	"fmt"
)

func shit() {
	fmt.Print()
}

//define a interface
type Namer interface {
	testInterface() string
	TestInterface2() string
}

//define a struct use interface
type Square struct {
	side string
}

func (this *Square) testInterface() string {
	return "interface fist method"
}

func (this *Square) TestInterface2() string {
	return "interface fist method2"
}

func (this *Square) TestInf() {
	fmt.Println(this.testInterface())
}




