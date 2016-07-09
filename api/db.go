package main

import (
  "fmt"
  "os"

  "gopkg.in/pg.v4"
)

func InitDB() *pg.DB {
  db := pg.Connect(&pg.Options{
    Addr: fmt.Sprintf("%v:%v",
      os.Getenv("DOCKER_POSTGRES_1_PORT_5432_TCP_ADDR"),
      os.Getenv("DOCKER_POSTGRES_1_PORT_5432_TCP_PORT"),
    ),
    User: os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PSWD"),
    Database: os.Getenv("DB_NAME"),
  })

  return db
}
