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
	"time"
	"encoding/gob"
)


type ArticleInfo struct {
	Id int
	Name string
	Modtime time.Time
}

var ArticleList []ArticleInfo

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



func ScanFiles(dirpath string) []ArticleInfo {
	file,  _ := os.Open(dirpath)
	list, err := file.Readdir(0)
	if err != nil {
		fmt.Println("Readdir error")
	}
	defer file.Close()

	article := ArticleInfo{}
	for i, item := range list {
		article.Id = i
		article.Name = item.Name()
		article.Modtime = item.ModTime()
		ArticleList = append(ArticleList, article)
	}
	return ArticleList
}

func CompareFiles(dirpath string) {
	list := ScanFiles(dirpath)
	_, err := os.Stat("list.gob")
	if err == nil {
		//存在 msg.gob 处理
		file, _ := os.Open("list.gob")
		gobde := gob.NewDecoder(file)
		gobde.Decode(&ArticleList)
		file.Close()
		fmt.Println(ArticleList)
		fmt.Println(len(ArticleList))
	}
	if os.IsNotExist(err) {
		//不存在 msg.gob 处理
		file, _ := os.Create("list.gob")
		goben := gob.NewEncoder(file)
		goben.Encode(list)
		file.Close()
	}
}