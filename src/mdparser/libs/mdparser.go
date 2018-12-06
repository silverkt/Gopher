/**
所有模块包
*/
package libs

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func MarkDownParser(str []byte) []byte {
	//str := "![avatar](/home/picture/1.png)"
	html := blackfriday.MarkdownCommon([]byte(str))
	//fmt.Printf("%s/n", html)
	return html
}

func ReadMDFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Read MarkDown Files Failure")
	} 
	return data	
}