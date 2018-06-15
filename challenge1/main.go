package main

import (
  "fmt"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/paxos-engineering-challenge/challenge1/handler"
)

func main() {
  // Create the route handler, which creates a connection to the SQLite database.
  hdlr, err := handler.NewHandler()
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  // Define the routes.
  router := gin.Default()
  router.POST("/messages", hdlr.Post)
  router.GET("/messages/:digest", hdlr.Get)

  // Run the gin router.
  if err := router.Run(); err != nil {
    fmt.Println(err.Error())
    os.Exit(2)
  }
}
