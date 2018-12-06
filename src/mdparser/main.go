package main

import (
	"fmt"
	"mdparser/libs"
)

func main() {
	fmt.Print("this is main")
	data := libs.ReadMDFile("README.md")
	res := libs.MarkDownParser(data)
	fmt.Println(res)
}
