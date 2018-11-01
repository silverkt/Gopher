package main;

import (
	"fmt";
)

func Add(x, y int, ch chan int) {
	z := x + y;
	fmt.Println(z);
	flag := <- ch;
	if flag != 10 {
		ch <- flag + 1;
	}
}

func main() {
	var ch chan int;
	ch = make(chan int);
	for i := 0; i < 10; i++ {
		go Add(i, i, ch);
	}
	ch <- 0;
	flag := <- ch;
	
	if flag != 10 {
		ch <- flag + 1;
	}

	fmt.Println("the ch is : ", <- ch);
	
}