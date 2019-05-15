package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()
	fmt.Println(t)
	time.Sleep(3 * time.Second)
	r := time.Now().Unix() - t
	fmt.Println(r)

}