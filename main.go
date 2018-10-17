package main

import (
	_ "WebPractice/router"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome You~")
}

func main() {
	//	http.Server = {
	//		Addr: ":8000"
	//	}
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
	//fmt.Println("Hello Web")
}
