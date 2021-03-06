/**
所有模块包
try to study gin
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
	var ArticleList []ArticleInfo
	file,  _ := os.Open(dirpath)
	list, err := file.Readdir(0)
	if err != nil {
		fmt.Println("Readdir error")
	}
	defer file.Close()

	article := ArticleInfo{}
	for i, item := range list {
		if item.Name() == ".DS_Store" {
			continue
		}
		article.Id = i
		article.Name = item.Name()
		article.Modtime = item.ModTime()
		// article.Modtime = func() time.Time {
		// 	mfile, _ := os.Open(dirpath+item.Name())
		// 	mfileinfo, _ := mfile.Stat()
		// 	defer mfile.Close()
		// 	return mfileinfo.ModTime()
			
		// }()
		ArticleList = append(ArticleList, article)
	}
	return ArticleList
}

func WriteGob(data []ArticleInfo, filepath string) {
	// 不管存不存在 list.gob文件， 不存在创建，存在就覆盖
	file, _ := os.Create(filepath)
	goben := gob.NewEncoder(file)
	goben.Encode(data)
	file.Close()
}


func CompareFiles(dirpath string) []ArticleInfo {
	var ArticleList []ArticleInfo
	list := ScanFiles(dirpath)
	_, err := os.Stat("./list.gob")
	if err == nil {
		//存在 list.gob 处理
		//对比 gob和读取的目录是否一致，一致则返回nil， 不一致则返回不一致文件的数据
		file, _ := os.Open("./list.gob")
		gobde := gob.NewDecoder(file)
		gobde.Decode(&ArticleList)
		file.Close()

		if reflect.DeepEqual(list, ArticleList){    //Equality for slices is not defined. slice can only be compared to nil
			fmt.Println("same list")
			return nil
		} else {
			fmt.Println("different list")
			fmt.Println("stored List:", ArticleList)
			fmt.Println("realdd List:", list)
			return PickupChanges(list, ArticleList) 
		}
		fmt.Println("stored List:", ArticleList)
		fmt.Println("realdd List:", list)
		// fmt.Println(len(ArticleList))
	}
	if os.IsNotExist(err) {
		//不存在 list.gob 处理
		WriteGob(list, "./list.gob")
		return list
	}
	fmt.Println("应该永远泡不到这里")
	return nil  //应该永远跑不到这里
	
}


func PickupChanges(reallist []ArticleInfo, storedlist []ArticleInfo) []ArticleInfo {
	var extlist []ArticleInfo  // 增量更新
	var flag bool
	var maxindex int



	// 实际列表中有删除文件
	for i, storeditem := range storedlist {
		for _, realitem := range reallist {
			if realitem.Name == storeditem.Name {
				flag = true
			}
		}
		if !flag {
			storedlist = append(storedlist[:i], storedlist[i+1:]...)
			// 删除 storedlist里面对这一项
		}
		flag = false
	}
	// 实际列表中有修改文件或者增加文件的情况
	for _, realitem := range reallist {
		for _, storeditem := range storedlist {
			if realitem.Name == storeditem.Name {
				flag = true
				if realitem.Modtime != storeditem.Modtime {
					storeditem.Modtime = realitem.Modtime
					extlist = append(extlist, storeditem)  				 
				}
			}
			//获取存储对最大id
			if storeditem.Id > maxindex {
				maxindex = storeditem.Id
			}
		}
		if !flag {
			realitem.Id = maxindex + 1
			extlist = append(extlist, realitem)	
		} 
		flag = false

	}
	
	//更新stored gob
	storedlist = append(storedlist, extlist...)
	WriteGob(storedlist, "./list.gob")
	fmt.Println("========= ==")
	fmt.Println(extlist)
	fmt.Println("===========")
	return extlist
	
}