package main;

import (
	//"io";
	"net/http";
	"html/template";
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtmlPath := "./index.html";
	htmtemplate, _ := template.ParseFiles(indexHtmlPath);
	htmtemplate.Execute(w, nil);
	//io.WriteString(w, `<p>hello world</p>`);
}



func main() {
	http.HandleFunc("/", indexHandler);
	http.ListenAndServe(":1234", nil);
}