package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	message := "Hello World from Go in minimal Docker container"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
	})
	fmt.Println("Started, serving at 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
