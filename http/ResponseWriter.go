package main;

import (
	//"fmt";
	"net/http";
	"io";
	"log";
)

/**
首页处理
*/
type indexHandler struct {
	message string;
}
func (this *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = myResponseWriter(w);
	io.WriteString(w, this.message);
}

/**
test处理
*/
type testHandler struct {
	message string;
}
func (this *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = myResponseWriter(w);
	io.WriteString(w, this.message);
}



/**
公共response设置
*/
func myResponseWriter(w http.ResponseWriter) http.ResponseWriter  {
	w.Header().Set("Content-Type", "text/html");
	return w;
}


func main() {
	mux := http.NewServeMux();
	mux.Handle("/", &indexHandler{"test<div style=background:#eee;>kskksksk</div>"});
	mux.Handle("/test", &testHandler{"test<div style=background:red;>other</div>"});
	err := http.ListenAndServe(":12345", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}