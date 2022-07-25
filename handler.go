package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
  Title string `json:"title"`
}

type Articles struct {
  Articles []Article `json:"articles"`
}

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
  res, err := json.Marshal(Articles{[]Article{{"title"}}})
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "An error occurred")
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)

}

func postDocuments(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)

  fmt.Fprintf(w, "create document")
}
