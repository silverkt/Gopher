package main

import (
	"fmt"
	"encoding/gob"
	"os"
)




func main() {
	testinfo := make(map[string]string)
	file, _ := os.Create("msg.gob")
	testinfo["abc"] = "abd"
	testinfo["shit"] = "shiddt"
	fmt.Println(testinfo)
	goben := gob.NewEncoder(file)
	goben.Encode(testinfo)
	file.Close()


    showInfo := make(map[string]string)
	file2, _ := os.Open("msg.gob")
	gobde := gob.NewDecoder(file2)
	gobde.Decode(&showInfo)
	file2.Close()
	fmt.Println(showInfo)
	fmt.Println(len(showInfo))

}