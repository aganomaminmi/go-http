package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/documents", handler)

  if err := http.ListenAndServe(":8080", nil); err !=nil {
    log.Fatal(err)
  }
}
