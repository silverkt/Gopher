package main 

import (
	"fmt"
	"net/http"
)


func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req.Header.Set("Connection", `keep-alive`)
	req.Header.Set("Accept-Encoding", `gzip, deflate, br`)
	req.Header.Set("Accept-Language", `zh-CN,zh;q=0.9`)
	req.Header.Set("Cache-Control", `max-age=0`)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`)
	req.Header.Set("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`)
	
	fmt.Println("Sending the Revised Request")
	_, _ = client.Do(req)
}