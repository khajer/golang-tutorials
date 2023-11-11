package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("server: %s /\n", r.Method)
	// fmt.Printf("http: / request\n")
	// fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
	// fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
	io.WriteString(w, "root \n")
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method: %s /hello\n", r.Method)

	paramValue := r.URL.Query().Get("name")
	fmt.Println("params[name]: ", paramValue)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("nil")
	}
	fmt.Println("request body")
	fmt.Println(string(body))

	if r.Method == "GET" {
		fmt.Println("method get")

		if err != nil {
			log.Fatal("error")
		}

	} else if r.Method == "POST" {
		fmt.Println("method post")
	}

	io.WriteString(w, "hello, world\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", HelloHandle)
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
