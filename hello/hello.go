package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	message := "Hello World from Go in minimal Docker container"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
	})
	port := 80
	if os.Getenv("PORT") != "" {
		if p, err := strconv.Atoi(os.Getenv("PORT")); err != nil {
			fmt.Printf("invalid port %s\n", os.Getenv("PORT"))
		} else {
			port = p
		}
	}
	fmt.Printf("Started, serving at %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
