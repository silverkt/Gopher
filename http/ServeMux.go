package main;

import (
	//"fmt";
	"net/http";
	"io";
	"log";
)

type indexHandler struct {
	message string;
}
func (this *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, this.message);
}



func main() {
	mux := http.NewServeMux();
	mux.Handle("/", &indexHandler{"test<div style=background:#eee;>kskksksk</div>"});
	err := http.ListenAndServe(":12345", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}