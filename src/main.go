package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("server started\n")
	http.HandleFunc("/hello-world", getHello)
	http.HandleFunc("/status", getStatus)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got '/hello-world' request\n")
	io.WriteString(w, "{body: “app-server”}\n")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got '/status' request\n")
	io.WriteString(w, "")
}

