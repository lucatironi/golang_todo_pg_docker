package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
)

var db = InitDB()

func main() {
  router := NewRouter()

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}
