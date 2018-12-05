package main

import (
	"fmt"
	"net/http"
)


func testHeader(res http.ResponseWriter, req *http.Request) {
	for i, item := range req.Header {
		fmt.Println(i, item)
	}
}



func main () {
	http.HandleFunc("/test", testHeader)
	http.ListenAndServe(":8080", nil)
}