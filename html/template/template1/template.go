package main;

import (
	"net/http";
	"html/template";
)

type Person struct {
	Name string;
	Age int;
	Title string;	
}


func testHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "./tpl/test.html";
	tpl, _ := template.ParseFiles(templatePath);
	p := Person{ Name: "silver", Age: 18, Title: "front-end-developer" }
	tpl.Execute(w, p);
}





func main() {
	http.HandleFunc("/test", testHandler);
	http.ListenAndServe(":3000", nil);
}