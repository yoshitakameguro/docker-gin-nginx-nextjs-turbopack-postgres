package main

import (
    "net/http"
    "app/config"

    "github.com/gin-gonic/gin"
)

func main() {
    db.Init()

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, "pong")
    })

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
