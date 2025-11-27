package main

import (
    "net/http"
    "log"
)

func main() {
    // Serve all files in Frontend folder under /static/
    fs := http.FileServer(http.Dir("../Frontend"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Serve index.html at root URL
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../Frontend/index.html")
    })

    log.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}