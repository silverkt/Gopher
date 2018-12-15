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
	"reflect"
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
		//article.Modtime = item.ModTime()
		article.Modtime = func() time.Time {
			mfile, _ := os.Open(dirpath+item.Name())
			mfileinfo, _ := mfile.Stat()
			defer mfile.Close()
			return mfileinfo.ModTime()
			
		}()
		ArticleList = append(ArticleList, article)
	}
	return ArticleList
}

func CompareFiles(dirpath string) {
	list := ScanFiles(dirpath)
	_, err := os.Stat("list.gob")
	if err == nil {
		//存在 list.gob 处理
		//对比 gob和读取的目录是否一致，一致则返回nil， 不一致则返回不一致文件的数据
		file, _ := os.Open("list.gob")
		gobde := gob.NewDecoder(file)
		gobde.Decode(&ArticleList)
		file.Close()

		if reflect.DeepEqual(list, ArticleList){    //Equality for slices is not defined. slice can only be compared to nil
			fmt.Println("same list")
		} else {
			fmt.Println("different list")
		}
		fmt.Println(ArticleList)
		fmt.Println(list)
		// fmt.Println(len(ArticleList))
	}
	if os.IsNotExist(err) {
		//不存在 list.gob 处理
		// 返回文件数据
		file, _ := os.Create("list.gob")
		goben := gob.NewEncoder(file)
		goben.Encode(list)
		file.Close()
	}
}