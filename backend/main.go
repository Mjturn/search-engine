package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/Mjturn/search-engine/routes"
)

func main() {
    router := gin.Default()
    router.Static("/static", "../frontend/static")
    router.LoadHTMLGlob("../frontend/templates/*")
    routes.HandleRoutes(router)

    err := router.Run(":8080")
    if err != nil {
        log.Fatal(err)
    }
}
