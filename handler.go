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
  res, err := json.Marshal(Articles{[]Article{{ Title: "title" }}})
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Unknown error occurred")
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)

}

type RequestBody struct {
  Title string `json:"title"`
}

type CreateDocumentResponse struct {
  Title string `json:"title"`
}

func postDocuments(w http.ResponseWriter, r *http.Request) {
  reqBody := RequestBody{}
  dec := json.NewDecoder(r.Body)
  err := dec.Decode(&reqBody)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "interface { title: string }")
    return
  }

  resBody := CreateDocumentResponse{ Title: reqBody.Title }
  res, err := json.Marshal(resBody)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Unknown error occurred")
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}
