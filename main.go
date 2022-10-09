package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getResponseData(endpoint string, sharedKeyId string, sharedKey string) []uint8 {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", endpoint, nil)
	
	// Headers
	req.Header.Add("db-shared-key-id", sharedKeyId)
	req.Header.Add("db-shared-key", sharedKey)
	
	// Fetch Request
	resp, err := client.Do(req)

	// If errorprint it and exit
	if err != nil {
	fmt.Print(err.Error())
	os.Exit(1)
    }

	// Get body of the response
	responseData, err := ioutil.ReadAll(resp.Body)

	// If errorprint it and exit
    if err != nil {
    	fmt.Print(err)
    }

	// Return the result
	return responseData
}

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
