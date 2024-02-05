package main

import (
  "github.com/gin-gonic/gin"
  "api_pdv/src/controller/routes"
  "api_pdv/src/database"

  "fmt"
  "os"
)

func main() {
  if err := database.SetupDB(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  r := gin.Default()

  routes.InitRoutes(&r.RouterGroup)

  r.Run(":3000")
}
