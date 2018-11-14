package main;

import (
	//"fmt";
	"net/http";
	"io";
	"log";
)

/**
公共response设置
*/
type commonHandler struct {

}
func (this *commonHandler) myResponseWriter(w http.ResponseWriter) http.ResponseWriter  {
	w.Header().Set("Content-Type", "text/html");
	return w;
}


/**
首页处理
*/
type indexHandler struct {
	message string;
	commonHandler;
}
func (this *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = this.myResponseWriter(w);
	io.WriteString(w, this.message);
}

/**
test处理
*/
type testHandler struct {
	message string;
	commonHandler;
}
func (this *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = this.myResponseWriter(w);
	io.WriteString(w, this.message);
}







func main() {
	mux := http.NewServeMux();
	mux.Handle("/", &indexHandler{message:"test<div style=background:#eee;>kskksksk</div>"});
	mux.Handle("/test", &testHandler{message:"test<div style=background:red;>testpage</div>"});
	err := http.ListenAndServe(":12345", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}