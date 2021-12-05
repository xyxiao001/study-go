package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = "localhost:8080"
	http.HandleFunc("/", handler)
	fmt.Printf("服务已经启动: %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
