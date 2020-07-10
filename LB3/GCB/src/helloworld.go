package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := "8080"
    if fromEnv := os.Getenv("port"); fromEnv != "" {
        port = fromEnv
    }

    server := http.NewServeMux()
    server.HandleFunc("/", hello)
    log.Fatal(http.ListenAndServe("0.0.0.0:"+port, server))
    log.Printf("Listening on 0.0.0.0:%s", port)
}

func hello(w http.ResponseWriter, r *http.Request) {
    log.Printf("Serving request: %s", r.URL.Path)
    host, _ := os.Hostname()
    fmt.Fprintf(w, "Hello world!\n")
    fmt.Fprintf(w, "Hostname: %s\n", host)
}