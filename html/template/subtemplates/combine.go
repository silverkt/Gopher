package main 

import (
	"fmt"
	"html/template"
	//"io/ioutil"
	"os"
)

type Person struct {
	Id int
	Name string
	Country string
}

func main() {
	silver := Person{Id: 100, Name: "silver.sun", Country: "china"}
	sun := &Person{}
	sun.Id = 23

	fmt.Println(silver, sun)
	tpl := template.New("layout.html") 
	tpl.Funcs(map[string]interface{}{"tihuan": tihuan})
	tpl.ParseFiles("./layout.html", "./sub.html", "header.html", "footer.html")
	
	file, _ := os.Create("./outputFiles/newFile.html")
	tpl.Execute(file, silver)
	tpl.Execute(os.Stdout, silver)
}

//注入模板的函数
func tihuan(str string) string {
	return str + "-------" + "成功！"
}