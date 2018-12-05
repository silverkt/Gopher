package main

import (
	"fmt"
	"net/http"
)


func testHeader(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello world")
}



func main () {
	http.HandleFunc("/test", testHeader)
	http.ListenAndServe(":8080", nil)
}