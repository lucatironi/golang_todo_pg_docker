package main

import (
  "fmt"
  "os"

  "gopkg.in/pg.v4"
)

func InitDB() *pg.DB {
  db := pg.Connect(&pg.Options{
    Addr: fmt.Sprintf("%v:%v", os.Getenv("DB_ADDR"), os.Getenv("DB_PORT")),
    User: os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PSWD"),
    Database: os.Getenv("DB_NAME"),
  })

  return db
}
