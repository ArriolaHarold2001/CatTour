package main

import (
  "net/http"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

func main(){

  router := mux.NewRouter()
  router.HandleFunc("/", getCatImg).Methods("GET", "OPTIONS")
  http.ListenAndServe(":8000", handlers.CORS()(router))
}

