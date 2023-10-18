package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
    gin.SetMode(gin.ReleaseMode)



    r := gin.Default()



    // Serve static files using Gin's built-in static file server.
    r.Static("/assets", "./public/assets")


    r.GET("/", func(c *gin.Context) {
        c.File("public/index.html")
    })





    if err := r.Run(":8080"); err != nil {
        fmt.Println("Error starting server:", err)
    } else {
        fmt.Println("Server is running on :8080")
    }
}

