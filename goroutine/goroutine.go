package main;

import (
	"fmt";
)

func Add(x, y int, ch chan int) {
	z := x + y;
	fmt.Println(z);// need to change to io
	ch <- x;
}

func main() {
	var ch chan int;
	ch = make(chan int);
	for i := 0; i < 10; i++ {
		go Add(i, i, ch);
	}
	 
	for j :=0; j < 10; j++ {
		fmt.Println("the ch is : ", <- ch);
	}
	
	
}