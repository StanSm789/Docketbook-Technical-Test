package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Initialize and load environment variales
func init() {

    err := godotenv.Load(".env")

    if err != nil {
		fmt.Print(err)
    }
}

func main() {
	router := gin.Default()
	router.Run("localhost:8080")
}
