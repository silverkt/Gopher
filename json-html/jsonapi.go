package main;

import (
	"net/http";
	//"encoding/json";
	//"io";
)


func testHandler(w http.ResponseWriter, r *http.Request) {
	str := `
	{
		"name": "silver",
		"title": "developer",
		"age": "18"
	}
	`;

	// // Stop here if its Preflighted OPTIONS request
    // if origin := r.Header.Get("Origin"); origin != "" {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers",
	// "Action, Module")   //有使用自定义头 需要这个,Action, Module是例子
	// }
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type","application/json");
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "OPTIONS" {
		return
	};
	if r.Method == "POST" {
		w.Write([]byte(str));
	};
	if r.Method == "GET" {
		w.WriteHeader(405);
		w.Write([]byte("Method Not Allowed"));
	}
	
}






func main() {
	http.HandleFunc("/test", testHandler);
	http.ListenAndServe(":8080", nil);
}