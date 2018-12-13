package main

import (
	"fmt"
	"mdparser/libs"
	"html/template"
)



type Index struct {
	Name string
	Content template.HTML
}

type Inner struct {
	Name string
	Content template.HTML
}

func main() {
	index := Index{}
	inner := Inner{}
	fmt.Print("this is main")
	data := libs.ReadMDFile("README.md")
	res := libs.MarkDownParser(data)
	//fmt.Printf("%s",res)
	libs.SaveHtmlFile("../abc.html", res)
	index.Name = "index"
	index.Content = template.HTML(string(res))

	inner.Name = "inner"
	inner.Content = template.HTML(`<p>this is innerHTML</p>`)
    
	tplnames := []string{"index.html", "header.html", "footer.html", "leftbar.html", "rightbar.html"}
	libs.CombineFile(tplnames, "index.html", index)


	tplnames1 := []string{"innerpage.html", "header.html", "footer.html", "leftbar.html", "rightbar.html"}
	libs.CombineFile(tplnames1, "innerpage.html", inner)


	// list := libs.ScanFiles("./")

	// for i, item := range list {
	// 	fmt.Println(i, item.Modtime)
	// }

	libs.CompareFiles("./")


}
