/**
所有模块包
*/
package libs

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"html/template"
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

func SaveHtmlFile(path string, content []byte) {
	ioutil.WriteFile(path, content, os.FileMode(0777))
}

func CombineFile(tplnames []string, filename string, data interface{}) {
	tpl := template.New(tplnames[0]) 
	//tpl.Funcs(map[string]interface{}{"tihuan": tihuan})
	for i, item := range tplnames {
		tplnames[i] = "./templates/"+ item
		fmt.Println(tplnames[i])
	}
	_, err := tpl.ParseFiles(tplnames...)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := os.Create("./public/"+filename)
	tpl.Execute(file, data)
}