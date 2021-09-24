package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/admin/index", Index)
	http.ListenAndServe(":80", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("admin/index,hello")
	w.Write([]byte("hello"))
}
