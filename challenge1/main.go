package main

import (
  "fmt"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/paxos-engineering-challenge/challenge1/handler"
)

func main() {
  hdlr := handler.NewHandler()
  router := gin.Default()
  router.POST("/messages", hdlr.Post)
  router.GET("/messages/:digest", hdlr.Get)
  if err := router.Run(); err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}
