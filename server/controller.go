package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

type CatImg struct {
  Url string `json:"url"`
}

const url = "https://api.thecatapi.com/v1/images/search?0709efa2-9cd7-42ea-a70f-a04237af5797"

func getCatImg(w http.ResponseWriter, r *http.Request) {

  var catImg []CatImg

  err := getJson(url, &catImg)
  if err != nil {
    log.Fatal(err.Error())
  }

  jBody,err := json.Marshal(catImg)
  if err != nil {
    log.Fatal(err.Error())
  }

  w.Write(jBody)
}

func getJson(url string, target interface{}) error{
  c := http.Client{Timeout: time.Duration(1) * time.Second}

  resp, err := c.Get(url)
  if err != nil {
    log.Fatal(err.Error())
  }
  
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err.Error())
  }

  return json.Unmarshal(body, target)
}
