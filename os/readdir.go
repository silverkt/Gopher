package main

import (
	"fmt"
	"io/ioutil"
	"os"
	//"time"
)

func main() {

	//利用ioutil读取目录列表
	list, err := ioutil.ReadDir("../")
	if err != nil {
		//
	}
	for _, item := range list {
		year, mon, day := item.ModTime().Date()
		week := item.ModTime().Weekday()
		fmt.Println(item.Name(), year, mon, day, week)
	}

	//利用os读取目录列表
	file,  _ := os.Open("test")
	fmt.Println(file.Name())
	list, err = file.Readdir(0)
	if err != nil {
		fmt.Println("error")
	}
	for _, item := range list {
		fmt.Println(item.Name(), item.ModTime())
	}

	defer func() {
		fmt.Println("defer called")
		file.Close()
	}()
	

}