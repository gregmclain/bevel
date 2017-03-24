package main

import (
    "fmt"
    "net/http"
    "time"
    "io"
    "log"
)

func ServerStart() {
	fmt.Println("Starting server")
    http.HandleFunc("/foo", ServeFoo)
    http.HandleFunc("/", Index)
    s := &http.Server{
    	Addr:           ":8080",
    	Handler:        nil,
    	ReadTimeout:    10 * time.Second,
    	WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}

func ServeFoo(response http.ResponseWriter, request *http.Request) {
    io.WriteString(response, "Hello, Foo!\n")
}

func Index(response http.ResponseWriter, request *http.Request) {
    io.WriteString(response, "Hello, World!\n")
}
