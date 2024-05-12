package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("newWeb/CialloWeb.txt")
	_, _ = fmt.Fprintln(w, string(b))
}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http failed", err)
		return
	}
}
