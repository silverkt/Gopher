package main

import (
	"fmt"
	"log"
	"net/http"
)



func HelloServer(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler");
	// fmt.Fprintf(res, "Hello,"+req.URL.Path[1:]);
	fmt.Fprintf(res, "<h1>sxq</h1><div>Hello World!!</div>");
}

func main() {
	fmt.Println("Hello World!");
	http.HandleFunc("/abc", HelloServer);
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error());
	}
}