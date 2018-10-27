package main

import (
	"fmt"
)

type FirstMan struct {
	name, address string;
	age int;
	weight int;
	height int;
}

type SecondMan struct {
	FirstMan
	company string
}

func main() {
	liming := new(FirstMan);
	liming.name = "silver";
	liming.address = "china";
	liming.weight = 146;
	liming.height = 182;

	var hanmeimei FirstMan;
	hanmeimei = FirstMan{
		name: "Clare",
		address: "USA", 
		age: 23,
		weight: 164, 
		height: 90};
	res := liming.run("shitly");
	fmt.Println(res,"test");

	_ = hanmeimei.run("fuckly");


	var xiaoming SecondMan;
	xiaoming.name ="xiaoming";
	xiaoming.address = "Japan";
	xiaoming.weight = 200;
	xiaoming.height = 163;
	xiaoming.company = "ABB";
	_ = xiaoming.run("傻啦吧唧");


	xiaoming = FirstMan{
		name: "Clare",
		address: "USA", 
		age: 23,
		weight: 164, 
		height: 90};
}

func (this FirstMan) run(how string) bool {
	fmt.Println(this.name+" from "+this.address+" is "+how+" running now!");
	return true;
}


