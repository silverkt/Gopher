package main

import (
	"fmt"
	"log"
	"net/http"
)


func HelloControl(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler");
	fmt.Fprintf(res, "Hello,"+req.URL.Path[1:]);
	//fmt.Fprintf(res, "<h1>sxq</h1><div>Hello World!!</div>");
}

func RootControl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is  all from root ");
}

func main() {
	fmt.Println("Hello World!");
	http.HandleFunc("/abc/", HelloControl);
	http.HandleFunc("/", RootControl);
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error());
	}
}