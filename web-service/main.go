package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// user attributes
type user struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
}

var ben = user{
		ID: "1",
		Name: "Ben",
		Email: "ben.klosky@gmail.com",
		Address: "2426 ontario rd NW, Washington, DC 20009",	
}

func getBen(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ben)
}

func main() {
	router := gin.Default()
	router.GET("/ben", getBen)

	//START GIN SERVER
	fmt.Println("Starting server on localhost:8080")
	router.Run("localhost:8080")
}