package main

import (
    "fmt"
    "log"
    "net/http"
)

func greetings(msg string) string {
    return fmt.Sprintf("<b>%s</b>", msg)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, greetings("Code.education Rocks!")+"\n")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}