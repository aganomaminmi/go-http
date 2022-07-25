package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      getDocuments(w, r)
    case "POST":
      postDocuments(w, r)
    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
  }
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "get document")

}

func postDocuments(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)

  fmt.Fprintf(w, "create document")
}
