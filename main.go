package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"task/docketbook-technical-test/models"
)

// This function handles GET request to /dockets endpoint and then returns the response as JSON file
func getDockets(c *gin.Context) {
	// Get the response from dockets URL as []uint8 type
	responseData := getResponseData(os.Getenv("DOCKETS_URL"), os.Getenv("SHARED_KEY_ID"), os.Getenv("SHARED_KEY"))

	// Unmarshal data into array
	var dockets []models.Docket
	json.Unmarshal([]byte(responseData), &dockets)
	
	// Return the response as indented JSON
	c.IndentedJSON(http.StatusOK, dockets)
}

// This function handles GET request to /dockets/:id endpoint and then returns the response as JSON file
func getDocketById(c *gin.Context) {
	// Get id 
	id := c.Param("id")

	// assemle single docket URL
	docketsUrl := os.Getenv("DOCKETS_URL") + "/" + id

	// Get the response from single docket URL as []uint8 type
	responseData := getResponseData(docketsUrl, os.Getenv("SHARED_KEY_ID"), os.Getenv("SHARED_KEY"))

	// Unmarshal data into Docket
	var docket models.Docket
	json.Unmarshal([]byte(responseData), &docket)

	// Return the response as indented JSON
	c.IndentedJSON(http.StatusOK, docket)
}

// This function fetches data from an endpoint
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
	// Initialize router and rountes, then run the router
	router := gin.Default()
	router.GET("/dockets", getDockets)
	router.GET("/dockets/:id", getDocketById)
	router.Run("localhost:8080")
}
