package main 

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sending the go origin Request")
	_, _ = http.Get("http://localhost:8080/test")
}