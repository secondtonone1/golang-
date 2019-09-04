package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("err is ", err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to zack web server")
}
