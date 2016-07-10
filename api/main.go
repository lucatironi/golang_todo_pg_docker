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

  log.Printf(fmt.Sprintf("Service listening on port %v", os.Getenv("PORT")))

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}
